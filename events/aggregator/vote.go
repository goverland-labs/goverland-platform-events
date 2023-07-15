package aggregator

import (
	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectVoteCreated = "aggregator.vote.created"
)

type VotePayload struct {
	ID         string `json:"id"`
	Ipfs       string `json:"ipfs"`
	ProposalID string `json:"proposal_id"`
	Voter      string `json:"voter"`
	Created    int    `json:"created"`
	Reason     string `json:"reason"`
}

type VotesPayload []VotePayload

type VotesHandler = events.Handler[VotesPayload]
