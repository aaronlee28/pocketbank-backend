package dto

type TopupDepositReq struct {
	Amount float32 `json:"amount" binding:"required"`
}
