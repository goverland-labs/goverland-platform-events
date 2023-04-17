package events

import "encoding/json"

const (
	SubjectProposalCreated = "proposal.created"
	SubjectProposalUpdated = "proposal.updated"
)

type (
	ProposalPayload struct {
		ID               string            `json:"id"`
		Ipfs             string            `json:"ipfs"`
		Author           string            `json:"author"`
		Created          int               `json:"created"`
		Dao              DaoPayload        `json:"dao"`
		Network          string            `json:"network"`
		Symbol           string            `json:"symbol"`
		Type             string            `json:"type"`
		Strategies       []StrategyPayload `json:"strategies"`
		Validation       ValidationPayload `json:"validation"`
		Plugins          interface{}       `json:"plugins"`
		Title            string            `json:"title"`
		Body             string            `json:"body"`
		Discussion       string            `json:"discussion"`
		Choices          []string          `json:"choices"`
		Start            int               `json:"start"`
		End              int               `json:"end"`
		Quorum           float64           `json:"quorum"`
		Privacy          string            `json:"privacy"`
		Snapshot         string            `json:"snapshot"`
		State            string            `json:"state"`
		Link             string            `json:"link"`
		App              string            `json:"app"`
		Scores           []float32         `json:"scores"`
		ScoresByStrategy interface{}       `json:"scores_by_strategy"`
		ScoresState      string            `json:"scores_state"`
		ScoresTotal      float32           `json:"scores_total"`
		ScoresUpdated    int               `json:"scores_updated"`
		Votes            int               `json:"votes"`
	}

	ProposalHandler func(ProposalPayload) error
)

func (h ProposalHandler) RawHandler() RawMessageHandler {
	return func(raw []byte) error {
		var d ProposalPayload
		if err := json.Unmarshal(raw, &d); err != nil {
			return err
		}

		return h(d)
	}
}
