package inbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
	"github.com/goverland-labs/goverland-platform-events/events/core"
)

const (
	SubjectFeedUpdated         = "inbox.feed.updated"
	SubjectFeedSettingsUpdated = "inbox.feed.settings.updated"
)

const (
	DaoCreated                  TimelineAction = "dao.created"
	DaoUpdated                  TimelineAction = "dao.updated"
	ProposalCreated             TimelineAction = "proposal.created"
	ProposalUpdated             TimelineAction = "proposal.updated"
	ProposalVotingStartsSoon    TimelineAction = "proposal.voting.starts_soon"
	ProposalVotingEndsSoon      TimelineAction = "proposal.voting.ends_soon"
	ProposalVotingStarted       TimelineAction = "proposal.voting.started"
	ProposalVotingQuorumReached TimelineAction = "proposal.voting.quorum_reached"
	ProposalVotingEnded         TimelineAction = "proposal.voting.ended"
	DelegateCreateProposal      TimelineAction = "delegate.proposal.created"
	DelegateVotingVoted         TimelineAction = "delegate.voting.voted"
	DelegateVotingSkipVote      TimelineAction = "delegate.voting.skip_vote"

	TypeDao      Type = "dao"
	TypeProposal Type = "proposal"
	TypeDelegate Type = "delegate"
)

var (
	ErrUnsupportedPayload = errors.New("unsupported payload")
	ErrWrongPayload       = errors.New("wrong payload")
)

type Type string

type TimelineAction string

type TimelineItem struct {
	CreatedAt time.Time      `json:"created_at"`
	Action    TimelineAction `json:"action"`
}

type FeedPayload struct {
	ID           uuid.UUID      `json:"id"`
	DaoID        uuid.UUID      `json:"dao_id"`
	ProposalID   string         `json:"proposal_id,omitempty"`
	DiscussionID string         `json:"discussion_id,omitempty"`
	Type         Type           `json:"type"`
	Action       TimelineAction `json:"action"`

	Snapshot json.RawMessage `json:"snapshot"`
	Timeline []TimelineItem  `json:"timeline"`
}

type FeedSettingsPayload struct {
	SubscriberID         uuid.UUID `json:"subscriber_id"`
	AutoarchiveAfterDays int       `json:"autoarchive_after_days"`
}

// todo: refactor it

func (f *FeedPayload) GetDAO() (*core.DaoPayload, error) {
	if f.Type != TypeDao {
		return nil, ErrUnsupportedPayload
	}

	var snapshot *core.DaoPayload
	err := json.Unmarshal(f.Snapshot, snapshot)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrWrongPayload, err)
	}

	return snapshot, nil
}

// todo: refactor it

func (f *FeedPayload) GetProposal() (*core.ProposalPayload, error) {
	if f.Type != TypeProposal {
		return nil, ErrUnsupportedPayload
	}

	var snapshot *core.ProposalPayload
	err := json.Unmarshal(f.Snapshot, snapshot)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrWrongPayload, err)
	}

	return snapshot, nil
}

type FeedHandler = events.Handler[FeedPayload]
type FeedSettingsHandler = events.Handler[FeedSettingsPayload]
