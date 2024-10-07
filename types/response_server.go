package types

const TypeServerResponseCreated ServerEventType = "response.created"

type ServerResponseCreated struct {
	EventID  string          `json:"event_id"`
	Type     ServerEventType `json:"type"`
	Response *ServerResponse `json:"response"`
	Usage    map[string]any  `json:"usage"`
}

const TypeServerResponse ServerEventType = "realtime.response"

type ServerResponse struct {
	ID            string           `json:"id"`
	Object        ServerEventType  `json:"object"`
	Status        *ResponseStatus  `json:"status"`
	StatusDetails map[string]any   `json:"status_details"`
	Output        []map[string]any `json:"output"` // NOTE: rendering for the outputs broken for me
}

const TypeServerResponseDone ServerEventType = "response.done"

type ServerResponseDone struct {
	EventID  string          `json:"event_id"`
	Type     ServerEventType `json:"type"`
	Response *ServerResponse `json:"response"`
	Usage    map[string]any  `json:"usage"`
}

const TypeServerResponseOutputItemAdded ServerEventType = "response.output_item.added"

type ServerResponseOutputItemAdded struct {
	EventID     string          `json:"event_id"`
	Type        ServerEventType `json:"type"`
	ResponseID  string          `json:"response_id"`
	OutputIndex uint32          `json:"output_index"`
	Item        *ServerResponseItem
}

const TypeServerResponseItem ObjectType = "realtime.item"

type ServerResponseItem struct {
	ID      string                 `json:"id"`
	Object  ObjectType             `json:"object"`
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

const TypeServerResponseOutputItemDone ServerEventType = "response.output_item.done"

type ServerResponseOutputItemDone struct {
	EventID     string              `json:"event_id"`
	Type        ServerEventType     `json:"type"`
	ResponseID  string              `json:"response_id"`
	OutputIndex uint32              `json:"output_index"`
	Item        *ServerResponseItem `json:"item"`
}

const TypeServerResponseContentPartAdded ServerEventType = "response.content_part.added"

type ServerResponseContentPartAdded struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Part         *Content        `json:"part"`
}

const TypeServerResponseContentPartDone ServerEventType = "response.content_part.done"

type ServerResponseContentPartDone struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Part         *Content        `json:"part"`
}

const TypeServerResponseTextDelta ServerEventType = "response.text.delta"

type ServerResponseTextDelta struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Delta        string          `json:"delta"`
}

const TypeServerResponseTextDone ServerEventType = "response.text.done"

type ServerResponseTextDone struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Text         string          `json:"text"`
}

const TypeServerResponseAudioTranscriptDelta ServerEventType = "response.audio_transcript.delta"

type ServerResponseAudioTranscriptDelta struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Delta        string          `json:"delta"`
}

const TypeServerResponseAudioTranscriptDone ServerEventType = "response.audio_transcript.done"

type ServerResponseAudioTranscriptDone struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Transcript   string          `json:"transcript"`
}

const TypeServerResponseAudioDelta ServerEventType = "response.audio.delta"

type ServerResponseAudioDelta struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	Delta        string          `json:"delta"`
}

const TypeServerResponseAudioDone ServerEventType = "response.audio.done"

type ServerResponseAudioDone struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
}

const TypeServerResponseFunctionCallArgumentsDelta ServerEventType = "response.function_call_arguments.delta"

type ServerResponseFunctionCallArgumentsDelta struct {
	EventID      string          `json:"event_id"`
	Type         ServerEventType `json:"type"`
	ResponseID   string          `json:"response_id"`
	ItemID       string          `json:"item_id"`
	OutputIndex  uint32          `json:"output_index"`
	ContentIndex uint32          `json:"content_index"`
	CallID       string          `json:"call_id"`
	Delta        string          `json:"delta"`
}

const TypeServerResponseFunctionCallArgumentsDone ServerEventType = "response.function_call_arguments.done"

type ServerResponseFunctionCallArgumentsDone struct {
	EventID     string          `json:"event_id"`
	Type        ServerEventType `json:"type"`
	ResponseID  string          `json:"response_id"`
	ItemID      string          `json:"item_id"`
	OutputIndex uint32          `json:"output_index"`
	CallID      string          `json:"call_id"`
	Arguments   string          `json:"arguments"`
}
