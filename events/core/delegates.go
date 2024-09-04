package core

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectDelegateUpsert = "core.delegate.upsert"
)

const (
	DelegateActionSet    = "set"
	DelegateActionClear  = "clear"
	DelegateActionExpire = "expire"
)

type DelegationDetails struct {
	Address string `json:"address"`
	Weight  int    `json:"weight"`
}

type Delegations struct {
	Details    []DelegationDetails `json:"details"`
	Expiration int                 `json:"expiration"`
}

type DelegatePayload struct {
	Action          string `json:"action"`
	AddressFrom     string `json:"address_from"`
	OriginalSpaceID string `json:"original_space_id"`
	ChainID         string `json:"chain_id"`
	BlockNumber     int    `json:"block_number"`
	BlockTimestamp  int    `json:"block_timestamp"`
	// Available for "set" action only
	Delegations Delegations `json:"delegations"`
}

type DelegateHandler = events.Handler[DelegatePayload]
