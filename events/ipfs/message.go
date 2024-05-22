package ipfs

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	// when we received a new message that contains ipfs identifier
	SubjectMessageCreated = "ipfs.message.created"
	// ipfs checker will be publish events to this subject after resolving data by ipfs identifier
	SubjectMessageCollected = "ipfs.message.collected"
)

type MessagePayload struct {
	IpfsID string `json:"ipfs_id"`
	Type   string `json:"type"`
}

type MessageHandler = events.Handler[MessagePayload]
