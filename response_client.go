package oairt

const TypeClientResponseCreate Type = "response.create"

type ClientResponseCreate struct {
	EventID  string    `json:"event_id"`
	Type     Type      `json:"type"`
	Response *Response `json:"response"`
}

func (c *ClientResponseCreate) isClient() {}

type Response struct { // NOTE: breaking with the naming schema a bit
	ID                string       `json:"id"`
	Modalities        []string     `json:"modalities"`
	Instructions      string       `json:"instructions"`
	Voice             *string      `json:"voice,omitempty"`
	OutputAudioFormat *AudioFormat `json:"output_audio_format,omitempty"`
	Tools             []*Tool      `json:"tools,omitempty"`
	ToolChoice        *string      `json:"tool_choice,omitempty"`
	Temperature       *float64     `json:"temperature,omitempty"`
	MaxOutputTokens   *uint32      `json:"max_output_tokens,omitempty"`
}

const TypeClientResponseCancel Type = "response.cancel"

type ClientResponseCancel struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
}

func (c *ClientResponseCancel) isClient() {}
