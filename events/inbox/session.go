package inbox

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectSessionCreated = "inbox.session.created"
)

type SessionCreatedEvent struct {
	UserID      uuid.UUID `json:"user_id"`
	AppVersion  string    `json:"app_version"`
	AppPlatform string    `json:"app_platform"`
}

type SessionCreatedHandler = events.Handler[SessionCreatedEvent]
