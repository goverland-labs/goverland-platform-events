package aggregator

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectDelegateERC20 = "aggregator.erc20indexer.updates"
)

type ERC20DelegatePayload struct {
	AddressFrom    string `json:"address_from"`
	AddressTo      string `json:"address_to"`
	Token          string `json:"token"`
	Network        string `json:"network"`
	BlockNumber    int    `json:"block_number"`
	BlockTimestamp int    `json:"block_timestamp"`
	VotingPower    int    `json:"voting_power"`
}

type ERC20DelegateHandler = events.Handler[ERC20DelegatePayload]
