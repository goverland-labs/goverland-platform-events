package aggregator

const (
	SourceSnapshot  = "snapshot"
	SourceSnapshotX = "snapshot-x"
)

type FilterPayload struct {
	MinScore    float32 `json:"min_score"`
	OnlyMembers bool    `json:"only_members"`
}

type ValidationPayload struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"`
}

type TreasuryPayload struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
}

type StrategyPayload struct {
	Name    string                 `json:"name"`
	Network string                 `json:"network"`
	Params  map[string]interface{} `json:"params"`
}
