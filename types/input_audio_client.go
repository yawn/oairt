package types

const TypeClientInputAudioBufferClear ClientEventType = "input_audio_buffer.clear"

type ClientInputAudioBufferAppend struct {
	EventID string          `json:"event_id"`
	Type    ClientEventType `json:"type"`
	Audio   string          `json:"audio"`
}

func (c *ClientInputAudioBufferAppend) isClientEvent() {}

const TypeClientInputAudioBufferCommit ClientEventType = "input_audio_buffer.commit"

type ClientInputAudioBufferCommit struct {
	EventID string          `json:"event_id"`
	Type    ClientEventType `json:"type"`
}

func (c *ClientInputAudioBufferCommit) isClientEvent() {}

const TypeClientInputAudioBufferAppend ClientEventType = "input_audio_buffer.append"

type ClientInputAudioBufferClear struct {
	EventID string          `json:"event_id"`
	Type    ClientEventType `json:"type"`
}

func (c *ClientInputAudioBufferClear) isClientEvent() {}
