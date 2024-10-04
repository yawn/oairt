package oairt

const TypeClientConversationItemCreate Type = "conversation.item.create"

type ClientConversationItemCreate struct {
	EventID        string                  `json:"event_id"`
	Type           Type                    `json:"type"`
	PreviousItemID string                  `json:"previous_item_id"`
	Item           *ClientConversationItem `json:"item"`
}

type ClientConversationItem struct {
	ID        string     `json:"id"`
	Type      Type       `json:"type"`
	Status    ItemStatus `json:"status"`
	Role      Role       `json:"role"`
	Content   []*Content `json:"content"`
	CallID    *string    `json:"call_id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Arguments *string    `json:"arguments,omitempty"`
	Output    *string    `json:"output,omitempty"`
}

const TypeClientConversationItemTruncate Type = "conversation.item.truncate"

type ClientConversationItemTruncate struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ItemID       string `json:"item_id"`
	ContentIndex uint32 `json:"content_index"`
	AudioEndMS   uint32 `json:"audio_end_ms"`
}

func (c *ClientConversationItemTruncate) isClient() {}

const TypeClientConversationItemDelete Type = "conversation.item.delete"

type ClientConversationItemDelete struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
	ItemID  string `json:"item_id"`
}

func (c *ClientConversationItemDelete) isClient() {}
