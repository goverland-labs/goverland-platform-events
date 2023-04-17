package events

type (
	FilterPayload struct {
		MinScore    float32 `json:"min_score"`
		OnlyMembers bool    `json:"only_members"`
	}

	ValidationPayload struct {
		Name   string      `json:"name"`
		Params interface{} `json:"params"`
	}

	TreasuryPayload struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Network string `json:"network"`
	}

	StrategyPayload struct {
		Name    string `json:"name"`
		Network string `json:"network"`
	}

	RawMessageHandler func([]byte) error
)
