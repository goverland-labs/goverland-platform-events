package core

import (
	"time"

	"github.com/google/uuid"

	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectTimelineUpdate = "core.timeline.updated"
)

type TimelineAction string

type TimelineItem struct {
	CreatedAt time.Time      `json:"created_at"`
	Action    TimelineAction `json:"action"`
}

type TimelinePayload struct {
	DaoID        uuid.UUID      `json:"dao_id"`
	ProposalID   string         `json:"proposal_id,omitempty"`
	DiscussionID string         `json:"discussion_id,omitempty"`
	Timeline     []TimelineItem `json:"timeline"`
}

type TimelineHandler = events.Handler[TimelinePayload]
