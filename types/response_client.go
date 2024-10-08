package types

const TypeClientResponseCreate ClientEventType = "response.create"

type ClientResponseCreate struct {
	EventID  string          `json:"event_id"`
	Type     ClientEventType `json:"type"`
	Response *ClientResponse `json:"response"`
}

func (c *ClientResponseCreate) isClientEvent() {}

type ClientResponse struct {
	Modalities        []string     `json:"modalities"`
	Instructions      string       `json:"instructions"`
	Voice             *string      `json:"voice,omitempty"`
	OutputAudioFormat *AudioFormat `json:"output_audio_format,omitempty"`
	Tools             []*Tool      `json:"tools,omitempty"`
	ToolChoice        *string      `json:"tool_choice,omitempty"`
	Temperature       *float64     `json:"temperature,omitempty"`
	MaxOutputTokens   *uint32      `json:"max_output_tokens,omitempty"`
}

const TypeClientResponseCancel ClientEventType = "response.cancel"

type ClientResponseCancel struct {
	EventID string          `json:"event_id"`
	Type    ClientEventType `json:"type"`
}

func (c *ClientResponseCancel) isClientEvent() {}
