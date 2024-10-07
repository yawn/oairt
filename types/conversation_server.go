package types

const TypeServerConversationCreated ServerEventType = "conversation.created"

type ServerConversationCreated struct {
	EventID      string              `json:"event_id"`
	Type         ServerEventType     `json:"type"`
	Conversation *ServerConversation `json:"conversation"`
}

const TypeServerConversation ObjectType = "realtime.conversation"

type ServerConversation struct {
	ID     string     `json:"id"`
	Object ObjectType `json:"object"`
}

const TypeServerConversationItemCreated ServerEventType = "conversation.item.created"

type ServerConversationItemCreated struct {
	EventID        string                  `json:"event_id"`
	Type           ServerEventType         `json:"type"`
	PreviousItemID string                  `json:"previous_item_id"`
	Item           *ServerConversationItem `json:"item"`
}

const TypeServerConversationItem ObjectType = "realtime.item"

type ServerConversationItem struct {
	ID        string     `json:"id"`
	Object    ObjectType `json:"object"`
	Type      string     `json:"type"`
	Status    ItemStatus `json:"status"`
	Role      Role       `json:"role"`
	Content   []Content  `json:"content"`
	CallID    string     `json:"call_id"`
	Name      string     `json:"name"`
	Arguments string     `json:"arguments"`
	Output    string     `json:"output"`
}

const TypeServerConversationInputAudioTranscriptionCompleted ServerEventType = "conversation.input_audio_transcription.completed"

type ServerConversationInputAudioTranscriptionCompleted struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ItemID       string          `json:"item_id"`
	ContentIndex uint32          `json:"content_index"`
	Transcript   string          `json:"transcript"`
}

const TypeServerConversationInputAudioTranscriptionFailed ServerEventType = "conversation.input_audio_transcription.failed"

type ServerConversationInputAudioTranscriptionFailed struct {
	EventID      string              `json:"event_id"`
	Type         ServerEventType     `json:"type"`
	ItemID       string              `json:"item_id"`
	ContentIndex uint32              `json:"content_index"`
	Error        *ServerErrorDetails `json:"error"` // NOTE: EventID is not part of the error details in this struct, used it anyways
}

const TypeServerConversationItemTruncated ServerEventType = "conversation.item.truncated"

type ServerConversationItemTruncated struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ItemID       string          `json:"item_id"`
	ContentIndex uint32          `json:"content_index"`
	AudioEndMS   uint32          `json:"audio_end_ms"`
}

const TypeServerConversationItemDeleted ServerEventType = "conversation.item.deleted"

type ServerConversationItemDeleted struct {
	EventID string          `json:"event_id"`
	Type    ServerEventType `json:"type"`
	ItemID  string          `json:"item_id"`
}
