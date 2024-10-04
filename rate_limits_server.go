package oairt

const TypeServerRateLimitsUpdated Type = "rate_limits.updated"

type ServerRateLimitsUpdated struct {
	EventID    string             `json:"event_id"`
	Type       Type               `json:"type"`
	RateLimits []*ServerRateLimit `json:"rate_limits"`
}

func (s *ServerRateLimitsUpdated) isServer() {}

type ServerRateLimit struct {
	Name         ServerRateLimitName `json:"name"`
	Limit        uint32              `json:"limit"`
	Remaining    uint32              `json:"remaining"`
	ResetSeconds float64             `json:"reset_seconds"` // NOTE: deviation from the docs
}

type ServerRateLimitName = string

const (
	ServerRateLimitNameRequests     ServerRateLimitName = "requests"
	ServerRateLimitNameTokens       ServerRateLimitName = "tokens"
	ServerRateLimitNameInputTokens  ServerRateLimitName = "input_tokens"
	ServerRateLimitNameOutputTokens ServerRateLimitName = "output_tokens"
)
