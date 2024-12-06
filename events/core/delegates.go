package core

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectDelegateCreated        = "core.delegates.created"
	SubjectDelegateCreateProposal = "core.delegates.create_proposal"
	SubjectDelegateVotingVoted    = "core.delegates.voting.voted"
	SubjectDelegateVotingSkipVote = "core.delegates.voting.skip_vote"
)

type DelegatePayload struct {
	// Initiator is delegate address who initiate the event
	Initiator string `json:"initiator"`
	// Delegator is delegator address (delegator delegate voting power to the initiator)
	Delegator string `json:"delegator"`
	// DaoID is internal dao identifier
	DaoID uuid.UUID `json:"dao_id"`
	// ProposalID is internal proposal identifier
	ProposalID string `json:"proposal_id"`
}

type DelegatesHandler = events.Handler[DelegatePayload]
