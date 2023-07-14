package inbox

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectProposalCreated      = "inbox.proposal.created"
	SubjectProposalUpdated      = "inbox.proposal.updated"
	SubjectProposalUpdatedState = "inbox.proposal.updated.state"

	SubjectProposalVotingStarted       = "inbox.proposal.voting.started"
	SubjectProposalVotingEnded         = "inbox.proposal.voting.ended"
	SubjectProposalVotingQuorumReached = "inbox.proposal.voting.quorum_reached"
	SubjectProposalVotingStartsSoon    = "inbox.proposal.voting.starts_soon"
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
}

type ProposalHandler func(ProposalPayload) error

func (h ProposalHandler) RawHandler() events.RawMessageHandler {
	return func(raw []byte) error {
		var d ProposalPayload
		if err := json.Unmarshal(raw, &d); err != nil {
			return err
		}

		return h(d)
	}
}
