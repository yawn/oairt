package oairt

import (
	"encoding/json"
	"fmt"
)

type Server interface {
	isServer()
}

func UnmarshalJSON(data []byte) (Server, Type, error) {

	var t tag

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, t.Type, err
	}

	var out Server

	switch t.Type {

	case TypeServerError:
		out = new(ServerError)

	case TypeServerSessionCreated:
		out = new(ServerSessionCreated)
	case TypeServerSessionUpdated:
		out = new(ServerSessionUpdated)

	case TypeServerConversationCreated:
		out = new(ServerConversationCreated)

	case TypeServerInputAudioBufferCommitted:
		out = new(ServerInputAudioBufferCommitted)
	case TypeServerInputAudioBufferCleared:
		out = new(ServerInputAudioBufferCleared)
	case TypeServerInputAudioBufferSpeechStarted:
		out = new(ServerInputAudioBufferSpeechStarted)
	case TypeServerInputAudioBufferSpeechStopped:
		out = new(ServerInputAudioBufferSpeechStopped)

	case TypeServerConversationItemCreated:
		out = new(ServerConversationItemCreated)
	case TypeServerConversationInputAudioTranscriptionCompleted:
		out = new(ServerConversationInputAudioTranscriptionCompleted)
	case TypeServerConversationInputAudioTranscriptionFailed:
		out = new(ServerConversationInputAudioTranscriptionFailed)
	case TypeServerConversationItemTruncated:
		out = new(ServerConversationItemTruncated)
	case TypeServerConversationItemDeleted:
		out = new(ServerConversationItemDeleted)

	case TypeServerResponseCreated:
		out = new(ServerResponseCreated)
	case TypeServerResponseDone:
		out = new(ServerResponseDone)

	case TypeServerResponseOutputItemAdded:
		out = new(ServerResponseOutputItemAdded)
	case TypeServerResponseOutputItemDone:
		out = new(ServerResponseOutputItemDone)

	case TypeServerResponseContentPartAdded:
		out = new(ServerResponseContentPartAdded)
	case TypeServerResponseContentPartDone:
		out = new(ServerResponseContentPartDone)

	case TypeServerResponseTextDelta:
		out = new(ServerResponseTextDelta)
	case TypeServerResponseTextDone:
		out = new(ServerResponseTextDone)

	case TypeServerResponseAudioTranscriptDelta:
		out = new(ServerResponseAudioTranscriptDelta)
	case TypeServerResponseAudioTranscriptDone:
		out = new(ServerResponseAudioTranscriptDone)

	case TypeServerResponseAudioDelta:
		out = new(ServerResponseAudioDelta)
	case TypeServerResponseAudioDone:
		out = new(ServerResponseAudioDone)

	case TypeServerResponseFunctionCallArgumentsDelta:
		out = new(ServerResponseFunctionCallArgumentsDelta)
	case TypeServerResponseFunctionCallArgumentsDone:
		out = new(ServerResponseFunctionCallArgumentsDone)

	case TypeServerRateLimitsUpdated:
		out = new(ServerRateLimitsUpdated)

	default:
		return nil, t.Type, fmt.Errorf("unknown server message ype %q", t.Type)

	}

	if err := json.Unmarshal(data, out); err != nil {
		return nil, t.Type, err
	}

	return out, t.Type, nil

}
