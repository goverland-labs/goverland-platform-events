package inbox

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectPushCreated = "inbox.push.created"
)

type PushPayload struct {
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	ImageURL string    `json:"image_url"`
	UserID   uuid.UUID `json:"user_id"`
}

type PushHandler = events.Handler[PushPayload]
