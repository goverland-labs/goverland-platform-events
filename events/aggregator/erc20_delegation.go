package aggregator

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectERC20Delegations = "aggregator.erc20indexer.delegations"
)

type ERC20DelegationPayload struct {
	Delegator      string `json:"delegator"`
	AddressFrom    string `json:"address_from"`
	AddressTo      string `json:"address_to"`
	Token          string `json:"token"`
	Network        string `json:"network"`
	BlockNumber    int64  `json:"block_number"`
	BlockTimestamp int64  `json:"block_timestamp"`
	LogIndex       int64  `json:"log_index"`
}

type ERC20DelegationHandler = events.Handler[ERC20DelegationPayload]
