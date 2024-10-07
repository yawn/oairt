package oairt

import "github.com/yawn/oairt/types"

type Event interface {
	*types.ServerError |
		*types.ServerSessionCreated | *types.ServerSessionUpdated |
		*types.ServerConversationCreated |
		*types.ServerInputAudioBufferCommitted | *types.ServerInputAudioBufferCleared | *types.ServerInputAudioBufferSpeechStarted | *types.ServerInputAudioBufferSpeechStopped |
		*types.ServerConversationItemCreated | *types.ServerConversationInputAudioTranscriptionCompleted | *types.ServerConversationInputAudioTranscriptionFailed | *types.ServerConversationItemTruncated | *types.ServerConversationItemDeleted |
		*types.ServerResponseCreated | *types.ServerResponseDone |
		*types.ServerResponseOutputItemAdded | *types.ServerResponseOutputItemDone |
		*types.ServerResponseContentPartAdded | *types.ServerResponseContentPartDone |
		*types.ServerResponseTextDelta | *types.ServerResponseTextDone |
		*types.ServerResponseAudioTranscriptDelta | *types.ServerResponseAudioTranscriptDone |
		*types.ServerResponseAudioDelta | *types.ServerResponseAudioDone |
		*types.ServerResponseFunctionCallArgumentsDelta | *types.ServerResponseFunctionCallArgumentsDone |
		*types.ServerRateLimitsUpdated
}

type Handler[T Event] struct {
	Handle func(T) (bool, error)
	ID     string
	Type   types.ServerEventType
}

func (c *Handler[T]) handle(event any) (bool, error) {
	return c.Handle(event.(T))
}

func (c *Handler[T]) id() string {
	return c.ID
}

func (c *Handler[T]) isApplicable(_type string) bool {
	return c.Type == _type
}

type handler interface {
	handle(event any) (cont bool, err error)
	id() (id string)
	isApplicable(_type string) (ok bool)
}
