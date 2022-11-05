package dto

type PaymentReq struct {
	ReceiverAccount int     `json:"receiverAccount" binding:"required"`
	Amount          float32 `json:"amount" binding:"required"`
	Description     string  `json:"description"`
}
