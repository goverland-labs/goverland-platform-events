package core

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectProposalCreated      = "core.proposal.created"
	SubjectProposalUpdated      = "core.proposal.updated"
	SubjectProposalUpdatedState = "core.proposal.updated.state"

	SubjectProposalVotingStarted       = "core.proposal.voting.started"
	SubjectProposalVotingEnded         = "core.proposal.voting.ended"
	SubjectProposalVotingQuorumReached = "core.proposal.voting.quorum_reached"
	SubjectProposalVotingStartsSoon    = "core.proposal.voting.starts_soon"
	SubjectProposalVotingEndsSoon      = "core.proposal.voting.ends_soon"
)

type ProposalPayload struct {
	ID            string            `json:"id"`
	Ipfs          string            `json:"ipfs"`
	Author        string            `json:"author"`
	Created       int               `json:"created"`
	DaoID         uuid.UUID         `json:"dao_id"`
	Network       string            `json:"network"`
	Symbol        string            `json:"symbol"`
	Type          string            `json:"type"`
	Strategies    []StrategyPayload `json:"strategies"`
	Title         string            `json:"title"`
	Body          string            `json:"body"`
	Discussion    string            `json:"discussion"`
	Choices       []string          `json:"choices"`
	Start         int               `json:"start"`
	End           int               `json:"end"`
	Quorum        float64           `json:"quorum"`
	Privacy       string            `json:"privacy"`
	Snapshot      string            `json:"snapshot"`
	State         string            `json:"state"`
	Link          string            `json:"link"`
	App           string            `json:"app"`
	Scores        []float32         `json:"scores"`
	ScoresState   string            `json:"scores_state"`
	ScoresTotal   float32           `json:"scores_total"`
	ScoresUpdated int               `json:"scores_updated"`
	Votes         int               `json:"votes"`
	EnsName       string            `json:"ens_name"`
	Spam          bool              `json:"spam"`
}

type ProposalHandler = events.Handler[ProposalPayload]
