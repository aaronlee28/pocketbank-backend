package dto

type TopupSavingsReq struct {
	Amount             float32 `json:"amount" binding:"required"`
	SenderWalletNumber int     `json:"senderWalletNumber"`
	Description        string  `json:"description"`
}
