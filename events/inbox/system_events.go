package inbox

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectSystem = "inbox.system"
)

type SystemEventType string

const (
	SystemEventTypeSessionCreated SystemEventType = "session_created"
)

type SessionCreatedEvent struct {
	UserID      uuid.UUID `json:"user_id"`
	AppVersion  string    `json:"app_version"`
	AppPlatform string    `json:"app_platform"`
}

type SystemEventPayload struct {
	Type     SystemEventType `json:"type"`
	Snapshot json.RawMessage `json:"snapshot"`
}

type SystemEventsHandler = events.Handler[SystemEventPayload]
