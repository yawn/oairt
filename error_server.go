package oairt

const TypeServerError Type = "error"

type ServerError struct {
	EventID string              `json:"event_id"`
	Type    Type                `json:"type"`
	Details *ServerErrorDetails `json:"error"`
}

func (s *ServerError) isServer() {}

type ServerErrorDetails struct {
	Type    string  `json:"type"`
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
	Param   *string `json:"param,omitempty"`
	EventID *string `json:"event_id,omitempty"`
}
