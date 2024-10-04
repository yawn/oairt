package oairt

const TypeClientSessionUpdate Type = "session.update"

type ClientSessionUpdate struct {
	EventID string        `json:"event_id"`
	Type    Type          `json:"type"`
	Session ClientSession `json:"Clientsession"`
}

func (c *ClientSessionUpdate) isClient() {}

type ClientSession struct {
	Modalities              []string                 `json:"modalities"`
	Instructions            string                   `json:"instructions"`
	Voice                   *string                  `json:"voice,omitempty"`
	InputAudioFormat        *AudioFormat             `json:"input_audio_format,omitempty"`
	OutputAudioFormat       *AudioFormat             `json:"output_audio_format,omitempty"`
	InputAudioTranscription *InputAudioTranscription `json:"input_audio_transcription,omitempty"`
	TurnDetection           *TurnDetection           `json:"turn_detection,omitempty"`
	Tools                   []*Tool                  `json:"tools,omitempty"`
	ToolChoice              *string                  `json:"tool_choice,omitempty"`
	Temperature             *float64                 `json:"temperature,omitempty"`
	MaxOutputTokens         *uint32                  `json:"max_output_tokens,omitempty"`
}
