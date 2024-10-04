package oairt

const TypeServerInputAudioBufferCommitted Type = "input_audio_buffer.committed"

type ServerInputAudioBufferCommitted struct {
	EventID        string `json:"event_id"`
	Type           Type   `json:"type"`
	PreviousItemID string `json:"previous_item_id"`
	ItemID         string `json:"item_id"`
}

func (s *ServerInputAudioBufferCommitted) isServer() {}

const TypeServerInputAudioBufferCleared Type = "input_audio_buffer.cleared"

type ServerInputAudioBufferCleared struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
}

func (s *ServerInputAudioBufferCleared) isServer() {}

const TypeServerInputAudioBufferSpeechStarted Type = "input_audio_buffer.speech_started"

type ServerInputAudioBufferSpeechStarted struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	AudioStartMS uint32 `json:"audio_start_ms"`
	ItemID       string `json:"item_id"`
}

func (s *ServerInputAudioBufferSpeechStarted) isServer() {}

const TypeServerInputAudioBufferSpeechStopped Type = "input_audio_buffer.speech_stopped"

type ServerInputAudioBufferSpeechStopped struct {
	EventID    string `json:"event_id"`
	Type       Type   `json:"type"`
	AudioEndMS uint32 `json:"audio_end_ms"`
	ItemID     string `json:"item_id"`
}

func (s *ServerInputAudioBufferSpeechStopped) isServer() {}
