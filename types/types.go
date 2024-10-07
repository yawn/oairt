package types

type (
	ClientEventType = string
	ObjectType      = string
	ServerEventType = string
)

type ClientEvent interface {
	isClientEvent()
}

func Float64(v float64) *float64 {
	return &v
}

func String(v string) *string {
	return &v
}

func Uint32(v uint32) *uint32 {
	return &v
}
