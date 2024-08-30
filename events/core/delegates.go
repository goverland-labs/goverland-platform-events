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
	DelegateActionOpt    = "opt"
)

type DelegatePayload struct {
	Action          string `json:"action"`
	AddressFrom     string `json:"address_from"`
	AddressTo       string `json:"address_to"`
	ChainID         string `json:"chain_id"`
	OriginalSpaceID string `json:"original_space_id"`
	ExpiredAt       int64  `json:"expired_at"`
	Weight          int    `json:"weight"`
	BlockNumber     int    `json:"block_number"`
}

type DelegateHandler = events.Handler[DelegatePayload]
