package aggregator

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectERC20Transfer = "aggregator.erc20indexer.transfers"
)

type ERC20TransferPayload struct {
	AddressFrom    string `json:"address_from"`
	AddressTo      string `json:"address_to"`
	Token          string `json:"token"`
	Network        string `json:"network"`
	BlockNumber    int64  `json:"block_number"`
	BlockTimestamp int64  `json:"block_timestamp"`
	LogIndex       int64  `json:"log_index"`
	Amount         string `json:"amount"`
}

type ERC20TransfersHandler = events.Handler[ERC20TransferPayload]
