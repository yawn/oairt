package oairt

const TypeServerResponseCreated Type = "response.created"

type ServerResponseCreated struct {
	EventID  string          `json:"event_id"`
	Type     Type            `json:"type"`
	Response *ServerResponse `json:"response"`
}

func (s *ServerResponseCreated) isServer() {}

const TypeServerResponse Type = "response.done"

type ServerResponse struct {
	ID            string           `json:"id"`
	Object        Type             `json:"object"`
	Status        *ResponseStatus  `json:"status"`
	StatusDetails map[string]any   `json:"status_details"`
	Output        []map[string]any `json:"output"` // NOTE: rendering for the outputs broken for me
	Usage         map[string]any   `json:"usage"`
}

const TypeServerResponseOutputItemAdded Type = "response.output_item.added"

type ServerResponseOutputItemAdded struct {
	EventID     string `json:"event_id"`
	Type        Type   `json:"type"`
	ResponseID  string `json:"response_id"`
	OutputIndex uint32 `json:"output_index"`
	Item        *ServerResponseItem
}

func (s *ServerResponseOutputItemAdded) isServer() {}

const TypeServerResponseItem Type = "realtime.item"

type ServerResponseItem struct {
	ID      string                 `json:"id"`
	Object  Type                   `json:"object"`
	Type    ServerResponseItemType `json:"type"`
	Status  ResponseStatus         `json:"status"`
	Role    Role                   `json:"role"`
	Content []*Content             `json:"content"`
}

type ServerResponseItemType = string

const (
	ServerResponseItemTypeMessage            ServerResponseItemType = "message"
	ServerResponseItemTypeFunctionCall       ServerResponseItemType = "function_call"
	ServerResponseItemTypeFunctionCallOutput ServerResponseItemType = "function_call_output"
)

const TypeServerResponseOutputItemDone Type = "response.output_item.done"

type ServerResponseOutputItemDone struct {
	EventID     string              `json:"event_id"`
	Type        Type                `json:"type"`
	ResponseID  string              `json:"response_id"`
	OutputIndex uint32              `json:"output_index"`
	Item        *ServerResponseItem `json:"item"`
}

func (s *ServerResponseOutputItemDone) isServer() {}

const TypeServerResponseContentPartAdded Type = "response.content_part.added"

type ServerResponseContentPartAdded struct {
	EventID      string   `json:"event_id"`
	Type         Type     `json:"type"`
	ResponseID   string   `json:"response_id"`
	ItemID       string   `json:"item_id"`
	OutputIndex  uint32   `json:"output_index"`
	ContentIndex uint32   `json:"content_index"`
	Part         *Content `json:"part"`
}

func (s *ServerResponseContentPartAdded) isServer() {}

const TypeServerResponseContentPartDone Type = "response.content_part.done"

type ServerResponseContentPartDone struct {
	EventID      string   `json:"event_id"`
	Type         Type     `json:"type"`
	ResponseID   string   `json:"response_id"`
	ItemID       string   `json:"item_id"`
	OutputIndex  uint32   `json:"output_index"`
	ContentIndex uint32   `json:"content_index"`
	Part         *Content `json:"part"`
}

func (s *ServerResponseContentPartDone) isServer() {}

const TypeServerResponseTextDelta Type = "response.text.delta"

type ServerResponseTextDelta struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
	Delta        string `json:"delta"`
}

func (s *ServerResponseTextDelta) isServer() {}

const TypeServerResponseTextDone Type = "response.text.done"

type ServerResponseTextDone struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
	Text         string `json:"text"`
}

func (s *ServerResponseTextDone) isServer() {}

const TypeServerResponseAudioTranscriptDelta Type = "response.audio_transcript.delta"

type ServerResponseAudioTranscriptDelta struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
	Delta        string `json:"delta"`
}

func (s *ServerResponseAudioTranscriptDelta) isServer() {}

const TypeServerResponseAudioTranscriptDone Type = "response.audio_transcript.done"

type ServerResponseAudioTranscriptDone struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
	Transcript   string `json:"transcript"`
}

func (s *ServerResponseAudioTranscriptDone) isServer() {}

const TypeServerResponseAudioDelta Type = "response.audio.delta"

type ServerResponseAudioDelta struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
	Delta        string `json:"delta"`
}

func (s *ServerResponseAudioDelta) isServer() {}

const TypeServerResponseAudioDone Type = "response.audio.done"

type ServerResponseAudioDone struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
}

func (s *ServerResponseAudioDone) isServer() {}

const TypeServerResponseFunctionCallArgumentsDelta Type = "response.function_call_arguments.delta"

type ServerResponseFunctionCallArgumentsDelta struct {
	EventID      string `json:"event_id"`
	Type         Type   `json:"type"`
	ResponseID   string `json:"response_id"`
	ItemID       string `json:"item_id"`
	OutputIndex  uint32 `json:"output_index"`
	ContentIndex uint32 `json:"content_index"`
	CallID       string `json:"call_id"`
	Delta        string `json:"delta"`
}

func (s *ServerResponseFunctionCallArgumentsDelta) isServer() {}

const TypeServerResponseFunctionCallArgumentsDone Type = "response.function_call_arguments.done"

type ServerResponseFunctionCallArgumentsDone struct {
	EventID     string `json:"event_id"`
	Type        Type   `json:"type"`
	ResponseID  string `json:"response_id"`
	ItemID      string `json:"item_id"`
	OutputIndex uint32 `json:"output_index"`
	CallID      string `json:"call_id"`
	Arguments   string `json:"arguments"`
}

func (s *ServerResponseFunctionCallArgumentsDone) isServer() {}
