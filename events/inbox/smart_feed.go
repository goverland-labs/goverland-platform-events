package inbox

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectSmartFeedUpdated = "inbox.smart_feed.updated"
)

type SmartFeedPayload struct {
	UserID     uuid.UUID       `json:"user_id"`
	DaoID      uuid.UUID       `json:"dao_id"`
	ProposalID *string         `json:"proposal_id,omitempty"`
	Action     TimelineAction  `json:"action"`
	Data       json.RawMessage `json:"data"`
}

type SmartFeedHandler = events.Handler[SmartFeedPayload]
