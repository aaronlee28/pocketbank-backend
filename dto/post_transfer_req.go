package dto

type TransferReq struct {
	ReceiverWalletNumber int     `json:"receiver_wallet_number" binding:"required"`
	Amount               float32 `json:"amount" binding:"required"`
	Description          string  `json:"description" `
}
