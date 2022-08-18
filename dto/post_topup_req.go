package dto

type TopupReq struct {
	Amount             float32 `json:"amount" binding:"required"`
	SenderWalletNumber int     `json:"sender_wallet_number"`
	Description        string  `json:"description"`
}
