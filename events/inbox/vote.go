package inbox

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectVoteCreated = "inbox.vote.created"
)

type VotePayload struct {
	UserID     uuid.UUID `json:"user_id"`
	DaoID      string    `json:"dao_id"`
	ProposalID string    `json:"proposal_id"`
}

type VoteHandler = events.Handler[VotePayload]
