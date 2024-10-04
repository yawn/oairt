// go:build example
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/coder/websocket"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/yawn/oairt"
	"golang.org/x/sync/errgroup"
)

func _main(ctx context.Context) error {

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY is not set")
	}

	conn, _, err := websocket.Dial(ctx, "wss://api.openai.com/v1/realtime?model=gpt-4o-realtime-preview-2024-10-01", &websocket.DialOptions{
		HTTPHeader: http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", apiKey)},
			"OpenAI-Beta":   []string{"realtime=v1"},
		},
	})

	if err != nil {
		return errors.Wrapf(err, "failed to dial websocket config")
	}

	defer conn.CloseNow()

	groups, ctx := errgroup.WithContext(ctx)
	groups.SetLimit(-1)

	var (
		in  = make(chan oairt.Client, 1024)
		out = make(chan oairt.Server, 1024)
	)

	groups.Go(func() error {

		for {

			kind, payload, err := conn.Read(ctx)

			if err != nil {
				return errors.Wrapf(err, "failed to read message")
			}

			if kind == websocket.MessageBinary {
				return errors.Wrapf(err, "received binary message")
			}

			msg, tag, err := oairt.UnmarshalJSON(payload)

			if err != nil {
				return errors.Wrapf(err, "failed to unmarshal message (%s)", string(payload))
			}

			log.Printf("received message %q: %s (%s)", tag, spew.Sdump(msg), string(payload))

			out <- msg

		}

	})

	groups.Go(func() error {

		for msg := range in {

			out, err := json.Marshal(msg)

			if err != nil {
				return errors.Wrapf(err, "failed to marshal message %v", msg)
			}

			log.Printf("send message: %s (%s)", spew.Sdump(msg), string(out))

			if err := conn.Write(ctx, websocket.MessageText, out); err != nil {
				return errors.Wrapf(err, "failed to write message %v", msg)
			}

		}

		return nil

	})

	groups.Go(func() error {

		// https://platform.openai.com/docs/guides/realtime

		sessionCreated := (<-out).(*oairt.ServerSessionCreated)
		fmt.Printf("session created: %s\n", sessionCreated.Session.ID)

		instructions := "Please tell me a joke. Just use text, no JSON formatting or any other formatting. Please also don't send images."

		fmt.Printf("> %s\n", instructions)
		fmt.Printf("< ")

		in <- &oairt.ClientResponseCreate{
			Type: oairt.TypeClientResponseCreate,
			Response: &oairt.ClientResponse{
				Modalities: []string{
					"text",
				},
				Instructions: instructions,
			},
		}

		for {

			switch msg := (<-out).(type) {
			case *oairt.ServerResponseTextDelta:
				fmt.Printf("%s", msg.Delta)
			case *oairt.ServerResponseTextDone:
				fmt.Println("\nClosing")
				close(in)
				return fmt.Errorf("done")
			}

		}

	})

	return groups.Wait()

}

func main() {

	if os.Getenv("DEBUG") != "true" {
		log.SetOutput(io.Discard)
	}

	if err := _main(context.Background()); err != nil {

		if err.Error() != "done" {
			panic(err)
		}

	}

}
