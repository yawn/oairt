package oairt

import (
	"encoding/json"
	"fmt"
)

type Server interface {
	isServer()
}

func UnmarshalJSON(data []byte) (Server, error) {

	type Tag struct {
		Type Type `json:"type"`
	}

	var tag Tag

	if err := json.Unmarshal(data, &tag); err != nil {
		return nil, err
	}

	var out Server

	switch tag.Type {

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
		return nil, fmt.Errorf("unknown server message ype %q", tag.Type)

	}

	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}

	return out, nil

}
