package events

type (
	RawMessageHandler func([]byte) error
)
