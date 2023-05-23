package core

type TreasuryPayload struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
}

type StrategyPayload struct {
	Name    string `json:"name"`
	Network string `json:"network"`
}
