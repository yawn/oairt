package types

const TypeServerInputAudioBufferCommitted ServerEventType = "input_audio_buffer.committed"

type ServerInputAudioBufferCommitted struct {
	EventID        string          `json:"event_id"`
	Type           ServerEventType `json:"type"`
	PreviousItemID string          `json:"previous_item_id"`
	ItemID         string          `json:"item_id"`
}

const TypeServerInputAudioBufferCleared ServerEventType = "input_audio_buffer.cleared"

type ServerInputAudioBufferCleared struct {
	EventID string          `json:"event_id"`
	Type    ServerEventType `json:"type"`
}

const TypeServerInputAudioBufferSpeechStarted ServerEventType = "input_audio_buffer.speech_started"

type ServerInputAudioBufferSpeechStarted struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	AudioStartMS uint32          `json:"audio_start_ms"`
	ItemID       string          `json:"item_id"`
}

const TypeServerInputAudioBufferSpeechStopped ServerEventType = "input_audio_buffer.speech_stopped"

type ServerInputAudioBufferSpeechStopped struct {
	EventID    string          `json:"event_id"`
	Type       ServerEventType `json:"type"`
	AudioEndMS uint32          `json:"audio_end_ms"`
	ItemID     string          `json:"item_id"`
}
