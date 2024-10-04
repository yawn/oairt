package oairt

type (
	AudioFormat = string
)

const (
	AudioFormatAlaw AudioFormat = "g711_alaw"
	AudioFormatPCM  AudioFormat = "pcm16"
	AudioFormatUlaw AudioFormat = "g711_ulaw"
)

type InputAudioTranscription struct {
	Enabled bool   `json:"enabled"`
	Model   string `json:"model"`
}

const TypeTurnDetectionServerVAD Type = "server_vad"

type TurnDetection struct {
	Type              Type    `json:"type"`
	Threshold         float64 `json:"threshold"`
	PrefixPaddingMS   uint32  `json:"prefix_padding_ms"`
	SilenceDurationMS uint32  `json:"silence_duration_ms"`
}

type Tool struct {
	Type        string         `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}
