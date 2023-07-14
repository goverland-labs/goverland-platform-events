package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Action string

const (
	ActionCreated             Action = "created"
	ActionUpdated             Action = "updated"
	ActionUpdatedState        Action = "updated.state"
	ActionVotingStartsSoon    Action = "voting.starts_soon"
	ActionVotingStarted       Action = "voting.started"
	ActionVotingQuorumReached Action = "voting.quorum_reached"
	ActionVotingEnded         Action = "voting.ended"

	DaoCreated                  TimelineAction = "dao.created"
	DaoUpdated                  TimelineAction = "dao.updated"
	ProposalCreated             TimelineAction = "proposal.created"
	ProposalUpdated             TimelineAction = "proposal.updated"
	ProposalVotingStartsSoon    TimelineAction = "proposal.voting.starts_soon"
	ProposalVotingStarted       TimelineAction = "proposal.voting.started"
	ProposalVotingQuorumReached TimelineAction = "proposal.voting.quorum_reached"
	ProposalVotingEnded         TimelineAction = "proposal.voting.ended"
)

type Type string

const (
	TypeDao      Type = "dao"
	TypeProposal Type = "proposal"
)

type FeedItem struct {
	ID           uuid.UUID `json:"id"`
	DaoID        uuid.UUID `json:"dao_id"`
	ProposalID   string    `json:"proposal_id"`
	DiscussionID string    `json:"discussion_id"`
	Type         Type      `json:"type"`
	Action       Action    `json:"action"`

	Snapshot json.RawMessage `json:"snapshot"`
	Timeline []TimelineItem  `json:"timeline"`
}

type TimelineItem struct {
	CreatedAt time.Time      `json:"created_at"`
	Action    TimelineAction `json:"action"`
}

type TimelineAction string

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
