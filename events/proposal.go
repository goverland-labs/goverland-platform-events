package events

import "encoding/json"

const (
	SubjectProposalCreated = "proposal.created"
	SubjectProposalUpdated = "proposal.updated"
)

type (
	SpacePayload struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	ProposalPayload struct {
		ID       string         `json:"id"`
		Title    string         `json:"title"`
		Body     string         `json:"body"`
		Choices  []string       `json:"choices"`
		Start    int64          `json:"start"`
		End      int64          `json:"end"`
		Snapshot string         `json:"snapshot"`
		State    string         `json:"state"`
		Author   string         `json:"author"`
		Spaces   []SpacePayload `json:"spaces"`
	}

	ProposalCreateHandler func(ProposalPayload) error
)

func (h ProposalCreateHandler) RawHandler() RawMessageHandler {
	return func(raw []byte) error {
		var d ProposalPayload
		if err := json.Unmarshal(raw, &d); err != nil {
			return err
		}

		return h(d)
	}
}
