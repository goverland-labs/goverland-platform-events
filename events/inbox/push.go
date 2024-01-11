package inbox

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectPushCreated = "inbox.push.created"

	PushVersionV1 = "v1"
	PushVersionV2 = "v2"
)

type PushVersion string

type PushPayload struct {
	Title         string          `json:"title"`
	Body          string          `json:"body"`
	ImageURL      string          `json:"image_url"`
	UserID        uuid.UUID       `json:"user_id"`
	Version       PushVersion     `json:"version,omitempty"`
	CustomPayload json.RawMessage `json:"custom_payload,omitempty"`
}

type PushHandler = events.Handler[PushPayload]
