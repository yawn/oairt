package oairt

type ItemStatus = string

const (
	ItemStatusCompleted  ItemStatus = "completed"
	ItemStatusInProgress ItemStatus = "in_progress"
	ItemStatusIncomplete ItemStatus = "incomplete"
)

type Role = string

const (
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
)

type Content struct {
	Type       ContentType `json:"type"`
	Text       *string     `json:"text,omitempty"`
	Audio      *string     `json:"audio,omitempty"`
	Transcript *string     `json:"transcript,omitempty"`
}

type ContentType = string

const (
	ContentTypeAudio      ContentType = "audio"
	ContentTypeInputAudio ContentType = "input_audio"
	ContentTypeInputText  ContentType = "input_text"
	ContentTypeText       ContentType = "text"
)
