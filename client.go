package oairt

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/coder/websocket"
	"github.com/pkg/errors"
	"github.com/yawn/oairt/types"
)

type tag struct {
	Type string `json:"type"`
}

type Client struct {
	apiKey   string
	conn     *websocket.Conn
	handlers map[string]handler
	lock     sync.RWMutex
}

func New(apiKey string) *Client {
	return &Client{
		apiKey:   apiKey,
		handlers: make(map[string]handler),
	}
}

func (c *Client) AddHandler(handler handler) string {

	id := handler.id()

	if id == "" {
		buf := make([]byte, 16)
		rand.Read(buf)
		id = hex.EncodeToString(buf)
	}

	// make adding handlers inside of handlers more idiomatic
	go func() {

		c.lock.Lock()
		defer c.lock.Unlock()

		c.handlers[id] = handler

	}()

	return id

}

func (c *Client) RemoveHander(id string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.handlers, id)
}

func (c *Client) Send(ctx context.Context, event types.ClientEvent) error {

	var tag string

	set := func(t types.ClientEventType) string {
		tag = t
		return t
	}

	switch event := event.(type) {
	case *types.ClientSessionUpdate:
		event.Type = set(types.TypeClientSessionUpdate)

	case *types.ClientInputAudioBufferAppend:
		event.Type = set(types.TypeClientInputAudioBufferAppend)
	case *types.ClientInputAudioBufferCommit:
		event.Type = set(types.TypeClientInputAudioBufferCommit)
	case *types.ClientInputAudioBufferClear:
		event.Type = set(types.TypeClientInputAudioBufferClear)

	case *types.ClientConversationItemCreate:
		event.Type = set(types.TypeClientConversationItemCreate)
	case *types.ClientConversationItemTruncate:
		event.Type = set(types.TypeClientConversationItemTruncate)
	case *types.ClientConversationItemDelete:
		event.Type = set(types.TypeClientConversationItemDelete)

	case *types.ClientResponseCreate:
		event.Type = set(types.TypeClientResponseCreate)
	case *types.ClientResponseCancel:
		event.Type = set(types.TypeClientResponseCancel)

	default:
		return fmt.Errorf("unknown event type: %#v", event)
	}

	payload, err := json.Marshal(event)

	if err != nil {
		return errors.Wrapf(err, "failed to marshal event")
	}

	slog.Debug("send",
		slog.String("payload", string(payload)),
		slog.String("type", tag),
	)

	if err := c.conn.Write(ctx, websocket.MessageText, payload); err != nil {
		return errors.Wrapf(err, "failed to write message")
	}

	return nil

}

func (c *Client) Start(ctx context.Context) error {

	conn, _, err := websocket.Dial(ctx, "wss://api.openai.com/v1/realtime?model=gpt-4o-realtime-preview-2024-10-01", &websocket.DialOptions{
		HTTPHeader: http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", c.apiKey)},
			"OpenAI-Beta":   []string{"realtime=v1"},
		},
	})

	if err != nil {
		return errors.Wrapf(err, "failed to dial")
	}

	conn.SetReadLimit(-1)

	c.conn = conn

	for {

		kind, payload, err := c.conn.Read(ctx)

		if err != nil {
			return errors.Wrapf(err, "failed to read message")
		}

		if kind == websocket.MessageBinary {
			return errors.Wrapf(err, "received binary message")
		}

		var tag tag

		if err := json.Unmarshal(payload, &tag); err != nil {
			return errors.Wrapf(err, "failed to unmarshal tag")
		}

		slog.Debug("recv",
			slog.String("type", tag.Type),
			slog.String("payload", string(payload)),
		)

		var target any

		switch tag.Type {

		case types.TypeServerError:
			target = new(types.ServerError)

		case types.TypeServerSessionCreated:
			target = new(types.ServerSessionCreated)
		case types.TypeServerSessionUpdated:
			target = new(types.ServerSessionUpdated)

		case types.TypeServerConversationCreated:
			target = new(types.ServerConversationCreated)

		case types.TypeServerInputAudioBufferCommitted:
			target = new(types.ServerInputAudioBufferCommitted)
		case types.TypeServerInputAudioBufferCleared:
			target = new(types.ServerInputAudioBufferCleared)
		case types.TypeServerInputAudioBufferSpeechStarted:
			target = new(types.ServerInputAudioBufferSpeechStarted)
		case types.TypeServerInputAudioBufferSpeechStopped:
			target = new(types.ServerInputAudioBufferSpeechStopped)

		case types.TypeServerConversationItemCreated:
			target = new(types.ServerConversationItemCreated)
		case types.TypeServerConversationInputAudioTranscriptionCompleted:
			target = new(types.ServerConversationInputAudioTranscriptionCompleted)
		case types.TypeServerConversationInputAudioTranscriptionFailed:
			target = new(types.ServerConversationInputAudioTranscriptionFailed)
		case types.TypeServerConversationItemTruncated:
			target = new(types.ServerConversationItemTruncated)
		case types.TypeServerConversationItemDeleted:
			target = new(types.ServerConversationItemDeleted)

		case types.TypeServerResponseCreated:
			target = new(types.ServerResponseCreated)
		case types.TypeServerResponseDone:
			target = new(types.ServerResponseDone)

		case types.TypeServerResponseOutputItemAdded:
			target = new(types.ServerResponseOutputItemAdded)
		case types.TypeServerResponseOutputItemDone:
			target = new(types.ServerResponseOutputItemDone)

		case types.TypeServerResponseContentPartAdded:
			target = new(types.ServerResponseContentPartAdded)
		case types.TypeServerResponseContentPartDone:
			target = new(types.ServerResponseContentPartDone)

		case types.TypeServerResponseTextDelta:
			target = new(types.ServerResponseTextDelta)
		case types.TypeServerResponseTextDone:
			target = new(types.ServerResponseTextDone)

		case types.TypeServerResponseAudioTranscriptDelta:
			target = new(types.ServerResponseAudioTranscriptDelta)
		case types.TypeServerResponseAudioTranscriptDone:
			target = new(types.ServerResponseAudioTranscriptDone)

		case types.TypeServerResponseAudioDelta:
			target = new(types.ServerResponseAudioDelta)
		case types.TypeServerResponseAudioDone:
			target = new(types.ServerResponseAudioDone)

		case types.TypeServerResponseFunctionCallArgumentsDelta:
			target = new(types.ServerResponseFunctionCallArgumentsDelta)
		case types.TypeServerResponseFunctionCallArgumentsDone:
			target = new(types.ServerResponseFunctionCallArgumentsDone)

		case types.TypeServerRateLimitsUpdated:
			target = new(types.ServerRateLimitsUpdated)

		default:
			return fmt.Errorf("unknown event type: %s", tag.Type)
		}

		if err := json.Unmarshal(payload, &target); err != nil {
			return errors.Wrapf(err, "failed to unmarshal message")
		}

		if err := c.handle(&tag, target); err != nil {
			return errors.Wrapf(err, "failed to handle event")
		}

	}

}

func (c *Client) handle(tag *tag, target any) error {

	var called bool

	c.lock.RLock()
	defer c.lock.RUnlock()

	for id, h := range c.handlers {

		if h.isApplicable(tag.Type) {

			called = true

			slog.Debug("calling handler",
				slog.String("type", tag.Type),
				slog.String("id", id),
			)

			cont, err := h.handle(target)

			if err != nil {
				return errors.Wrapf(err, "message handler %q failed", id)
			}

			if !cont {
				break
			}

		}

	}

	if !called {

		slog.Warn("no handlers called for event",
			slog.String("type", tag.Type),
		)

	}

	return nil

}
