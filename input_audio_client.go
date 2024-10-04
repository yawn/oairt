package oairt

const TypeClientInputAudioBufferClear Type = "input_audio_buffer.clear"

type ClientInputAudioBufferAppend struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
	Audio   string `json:"audio"`
}

func (c *ClientInputAudioBufferAppend) isClient() {}

const TypeClientInputAudioBufferCommit Type = "input_audio_buffer.commit"

type ClientInputAudioBufferCommit struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
}

func (c *ClientInputAudioBufferCommit) isClient() {}

const TypeClientInputAudioBufferAppend Type = "input_audio_buffer.append"

type ClientInputAudioBufferClear struct {
	EventID string `json:"event_id"`
	Type    Type   `json:"type"`
}

func (c *ClientInputAudioBufferClear) isClient() {}
