package oairt

const TypeServerConversationCreated Type = "conversation.created"

type ServerConversationCreated struct {
	EventID      string              `json:"event_id"`
	Type         Type                `json:"type"`
	Conversation *ServerConversation `json:"conversation"`
}

func (s *ServerConversationCreated) isServer() {}

const TypeServerConversation Type = "realtime.conversation"

type ServerConversation struct {
	ID     string `json:"id"`
	Object Type   `json:"object"`
}

const TypeServerConversationItemCreated Type = "conversation.item.created"

type ServerConversationItemCreated struct {
	EventID        string                  `json:"event_id"`
	Type           Type                    `json:"type"`
	PreviousItemID string                  `json:"previous_item_id"`
	Item           *ServerConversationItem `json:"item"`
}

func (s *ServerConversationItemCreated) isServer() {}

const TypeServerConversationItem Type = "realtime.item"

type ServerConversationItem struct {
	ID        string     `json:"id"`
	Object    Type       `json:"object"`
	Type      Type       `json:"type"`
	Status    ItemStatus `json:"status"`
	Role      Role       `json:"role"`
	Content   []Content  `json:"content"`
	CallID    string     `json:"call_id"`
	Name      string     `json:"name"`
	Arguments string     `json:"arguments"`
	Output    string     `json:"output"`
}

const TypeServerConversationInputAudioTranscriptionCompleted Type = "conversation.input_audio_transcription.completed"

type ServerConversationInputAudioTranscriptionCompleted struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ItemID       string `json:"item_id"`
	ContentIndex uint32 `json:"content_index"`
	Transcript   string `json:"transcript"`
}

func (s *ServerConversationInputAudioTranscriptionCompleted) isServer() {}

const TypeServerConversationInputAudioTranscriptionFailed Type = "conversation.input_audio_transcription.failed"

type ServerConversationInputAudioTranscriptionFailed struct {
	EventID      string              `json:"event_id"`
	Type         Type                `json:"type"`
	ItemID       string              `json:"item_id"`
	ContentIndex uint32              `json:"content_index"`
	Error        *ServerErrorDetails `json:"error"` // NOTE: EventID is not part of the error details in this struct, used it anyways
}

func (s *ServerConversationInputAudioTranscriptionFailed) isServer() {}

const TypeServerConversationItemTruncated Type = "conversation.item.truncated"

type ServerConversationItemTruncated struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ItemID       string `json:"item_id"`
	ContentIndex uint32 `json:"content_index"`
	AudioEndMS   uint32 `json:"audio_end_ms"`
}

func (s *ServerConversationItemTruncated) isServer() {}

const TypeServerConversationItemDeleted Type = "conversation.item.deleted"

type ServerConversationItemDeleted struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
	ItemID  string `json:"item_id"`
}

func (s *ServerConversationItemDeleted) isServer() {}
