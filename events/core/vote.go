package core

import (
	"github.com/google/uuid"
	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectVoteCreated = "core.vote.created"
)

type VotePayload struct {
	ID           string    `json:"id"`
	Ipfs         string    `json:"ipfs"`
	Voter        string    `json:"voter"`
	Created      int       `json:"created"`
	DaoID        uuid.UUID `json:"dao_id"`
	ProposalID   string    `json:"proposal_id"`
	Choice       int       `json:"choice"`
	Reason       string    `json:"reason"`
	App          string    `json:"app"`
	Vp           float64   `json:"vp"`
	VpByStrategy []float64 `json:"vp_by_strategy"`
	VpState      string    `json:"vp_state"`
}

type VotesPayload []VotePayload

type VotesHandler = events.Handler[VotesPayload]
