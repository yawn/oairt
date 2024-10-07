package types

const TypeServerRateLimitsUpdated ServerEventType = "rate_limits.updated"

type ServerRateLimitsUpdated struct {
	EventID    string             `json:"event_id"`
	Type       ServerEventType    `json:"type"`
	RateLimits []*ServerRateLimit `json:"rate_limits"`
}

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
