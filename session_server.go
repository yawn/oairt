package oairt

const TypeServerSessionCreated Type = "session.created"

type ServerSessionCreated struct {
	EventID string         `json:"event_id"`
	Type    Type           `json:"type"`
	Session *ServerSession `json:"session"`
}

func (s *ServerSessionCreated) isServer() {}

const TypeServerSession Type = "realtime.session"

type ServerSession struct {
	ID                      string                   `json:"id"`
	Object                  Type                     `json:"object"`
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

const TypeServerSessionUpdated Type = "session.updated"

type ServerSessionUpdated struct {
	EventID string         `json:"event_id"`
	Type    Type           `json:"type"`
	Session *ServerSession `json:"session"`
}

func (s *ServerSessionUpdated) isServer() {}
