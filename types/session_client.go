package types

const TypeClientSessionUpdate ClientEventType = "session.update"

type ClientSessionUpdate struct {
	EventID string          `json:"event_id"`
	Type    ClientEventType `json:"type"`
	Session ClientSession   `json:"session"`
}

func (c *ClientSessionUpdate) isClientEvent() {}

type ClientSession struct {
	Modalities              []string                 `json:"modalities,omitempty"`
	Instructions            *string                  `json:"instructions,omitempty"`
	Voice                   *string                  `json:"voice,omitempty"` // TODO: makes this an enum (alloy, echo, shimmer)
	InputAudioFormat        *AudioFormat             `json:"input_audio_format,omitempty"`
	OutputAudioFormat       *AudioFormat             `json:"output_audio_format,omitempty"`
	InputAudioTranscription *InputAudioTranscription `json:"input_audio_transcription,omitempty"`
	TurnDetection           *TurnDetection           `json:"turn_detection,omitempty"`
	Tools                   []*Tool                  `json:"tools,omitempty"`
	ToolChoice              *string                  `json:"tool_choice,omitempty"`
	Temperature             *float64                 `json:"temperature,omitempty"`
	MaxOutputTokens         *uint32                  `json:"max_output_tokens,omitempty"`
}
