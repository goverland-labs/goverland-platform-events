package core

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectDaoCreated = "core.dao.created"
	SubjectDaoUpdated = "core.dao.updated"

	SubjectCheckActivitySince     = "core.dao.check.activity_since"
	SubjectPopularityIndexUpdated = "core.dao.popularity_index.updated"
)

type VotingPayload struct {
	Delay       int     `json:"delay"`
	Period      int     `json:"period"`
	Type        string  `json:"type"`
	Quorum      float32 `json:"quorum"`
	Blind       bool    `json:"blind"`
	HideAbstain bool    `json:"hide_abstain"`
	Privacy     string  `json:"privacy"`
	Aliased     bool    `json:"aliased"`
}

type DaoPayload struct {
	ID              uuid.UUID         `json:"id"`
	Alias           string            `json:"alias"`
	Name            string            `json:"name"`
	Private         bool              `json:"private"`
	About           string            `json:"about"`
	Avatar          string            `json:"avatar"`
	Terms           string            `json:"terms"`
	Location        string            `json:"location"`
	Website         string            `json:"website"`
	Twitter         string            `json:"twitter"`
	Github          string            `json:"github"`
	Coingecko       string            `json:"coingecko"`
	Email           string            `json:"email"`
	Network         string            `json:"network"`
	Symbol          string            `json:"symbol"`
	Skin            string            `json:"skin"`
	Domain          string            `json:"domain"`
	Strategies      []StrategyPayload `json:"strategies"`
	Admins          []string          `json:"admins"`
	Members         []string          `json:"members"`
	Moderators      []string          `json:"moderators"`
	Voting          VotingPayload     `json:"voting"`
	Categories      []string          `json:"categories"`
	Treasures       []TreasuryPayload `json:"treasures"`
	FollowersCount  int               `json:"followers_count"`
	ProposalsCount  int               `json:"proposals_count"`
	Guidelines      string            `json:"guidelines"`
	Template        string            `json:"template"`
	ParentID        *uuid.UUID        `json:"parent_id,omitempty"`
	ActiveSince     *int              `json:"active_since,omitempty"`
	PopularityIndex *float64          `json:"popularity_index,omitempty"`
}

type DaoHandler = events.Handler[DaoPayload]
