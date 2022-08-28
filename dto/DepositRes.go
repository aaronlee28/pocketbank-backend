package dto

type DepositRes struct {
	Amount      float32 `json:"amount"`
	Duration    int     `json:"duration"`
	AutoDeposit bool    `json:"autoDeposit"`
}
