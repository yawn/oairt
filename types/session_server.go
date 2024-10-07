package types

const TypeServerSessionCreated ServerEventType = "session.created"

type ServerSessionCreated struct {
	EventID string          `json:"event_id"`
	Type    ServerEventType `json:"type"`
	Session *ServerSession  `json:"session"`
}

const TypeServerSession ObjectType = "realtime.session"

type ServerSession struct {
	ID                      string                   `json:"id"`
	Object                  ObjectType               `json:"object"`
	Model                   string                   `json:"model"`
	Modalities              []string                 `json:"modalities"`
	Instructions            string                   `json:"instructions"`
	Voice                   string                   `json:"voice,omitempty"`
	InputAudioFormat        AudioFormat              `json:"input_audio_format"`
	OutputAudioFormat       AudioFormat              `json:"output_audio_format"`
	InputAudioTranscription *InputAudioTranscription `json:"input_audio_transcription"`
	TurnDetection           *TurnDetection           `json:"turn_detection"`
	Tools                   []*Tool                  `json:"tools,omitempty"`
	ToolChoice              string                   `json:"tool_choice"`
	Temperature             float64                  `json:"temperature"`
	MaxOutputTokens         uint32                   `json:"max_output_tokens"`
}

const TypeServerSessionUpdated ServerEventType = "session.updated"

type ServerSessionUpdated struct {
	EventID string          `json:"event_id"`
	Type    ServerEventType `json:"type"`
	Session *ServerSession  `json:"session"`
}
