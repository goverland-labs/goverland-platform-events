package ipfs

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectMessageCreated   = "ipfs.message.created"
	SubjectMessageCollected = "ipfs.message.collected"
)

type MessagePayload struct {
	IpfsID string `json:"ipfs_id"`
	Type   string `json:"type"`
}

type MessageHandler = events.Handler[MessagePayload]
