package dto

type TopupDepositReq struct {
	Amount      float32 `json:"amount" binding:"required"`
	Duration    int     `json:"duration" binding:"required"`
	AutoDeposit bool    `json:"autoDeposit"`
}
