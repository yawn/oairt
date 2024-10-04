package oairt

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Client interface {
	isClient()
}

func MarshalJSON(c Client) ([]byte, Type, error) {

	var tag Type

	switch c.(type) {

	case *ClientSessionUpdate:
		tag = TypeClientSessionUpdate

	case *ClientInputAudioBufferAppend:
		tag = TypeClientInputAudioBufferAppend
	case *ClientInputAudioBufferCommit:
		tag = TypeClientInputAudioBufferCommit
	case *ClientInputAudioBufferClear:
		tag = TypeClientInputAudioBufferClear

	case *ClientConversationItemCreate:
		tag = TypeClientConversationItemCreate
	case *ClientConversationItemTruncate:
		tag = TypeClientConversationItemTruncate
	case *ClientConversationItemDelete:
		tag = TypeClientConversationItemDelete

	case *ClientResponseCreate:
		tag = TypeClientResponseCreate
	case *ClientResponseCancel:
		tag = TypeClientResponseCancel

	default:
		return nil, tag, fmt.Errorf("unknown client message type %T", c)

	}

	out, err := json.Marshal(c)

	if err != nil {
		return nil, tag, errors.Wrapf(err, "failed to marshal json")
	}

	return out, tag, nil
}
