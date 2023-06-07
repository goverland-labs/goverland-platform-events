package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Action string

const (
	ActionCreated             Action = "created"
	ActionUpdated                    = "updated"
	ActionUpdatedState               = "updated.state"
	ActionVotingStarted              = "voting.started"
	ActionVotingEnded                = "voting.ended"
	ActionVotingQuorumReached        = "voting.quorum_reached"
	ActionVotingStartsSoon           = "voting.starts_soon"
)

func ConvertActionToExternal(action string) Action {
	switch action {
	case SubjectDaoCreated, SubjectProposalCreated:
		return ActionCreated
	case SubjectDaoUpdated, SubjectProposalUpdated:
		return ActionUpdated
	case SubjectProposalUpdatedState:
		return ActionUpdatedState
	case SubjectProposalVotingStarted:
		return ActionVotingStarted
	case SubjectProposalVotingStartsSoon:
		return ActionVotingStartsSoon
	case SubjectProposalVotingQuorumReached:
		return ActionVotingQuorumReached
	case SubjectProposalVotingEnded:
		return ActionVotingEnded
	default:
		return ""
	}
}

type Type string

const (
	TypeDao      Type = "dao"
	TypeProposal Type = "proposal"
)

type FeedItem struct {
	DaoID        string `json:"dao_id"`
	ProposalID   string `json:"proposal_id"`
	DiscussionID string `json:"discussion_id"`
	Type         Type   `json:"type"`
	Action       Action `json:"action"`

	Snapshot json.RawMessage `json:"snapshot"`
}

// todo: refactor it
func (f *FeedItem) GetDAO() (*DaoPayload, error) {
	if f.Type != TypeDao {
		return nil, errors.New("unsupported payload")
	}

	var dao *DaoPayload
	err := json.Unmarshal(f.Snapshot, dao)
	if err != nil {
		return nil, fmt.Errorf("unmarshal snapshot: %w", err)
	}

	return dao, nil
}

// todo: refactor it
func (f *FeedItem) GetProposal() (*ProposalPayload, error) {
	if f.Type != TypeProposal {
		return nil, errors.New("unsupported payload")
	}

	var pp *ProposalPayload
	err := json.Unmarshal(f.Snapshot, pp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal snapshot: %w", err)
	}

	return pp, nil
}
