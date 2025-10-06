package aggregator

import (
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectERC20VPChanges = "aggregator.erc20indexer.vp_changes"
)

type ERC20VPChangesPayload struct {
	Address        string `json:"address"`
	Token          string `json:"token"`
	Network        string `json:"network"`
	BlockNumber    int64  `json:"block_number"`
	BlockTimestamp int64  `json:"block_timestamp"`
	LogIndex       int64  `json:"log_index"`
	VotingPower    string `json:"voting_power"`
	DeltaPower     string `json:"delta_power"`
}

type ERC20VPChangesHandler = events.Handler[ERC20VPChangesPayload]
