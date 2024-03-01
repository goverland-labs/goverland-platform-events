package aggregator

import (
	"encoding/json"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectVoteCreated = "aggregator.vote.created"
)

type VotePayload struct {
	ID            string          `json:"id"`
	Ipfs          string          `json:"ipfs"`
	Voter         string          `json:"voter"`
	Created       int             `json:"created"`
	OriginalDaoID string          `json:"original_dao_id"`
	ProposalID    string          `json:"proposal_id"`
	Choice        json.RawMessage `json:"choice"`
	Reason        string          `json:"reason"`
	App           string          `json:"app"`
	Vp            float64         `json:"vp"`
	VpByStrategy  []float64       `json:"vp_by_strategy"`
	VpState       string          `json:"vp_state"`
}

type VotesPayload []VotePayload

type VotesHandler = events.Handler[VotesPayload]
