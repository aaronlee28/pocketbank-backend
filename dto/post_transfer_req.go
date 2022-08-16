package dto

type TransferReq struct {
	ReceiverWalletNumber int    `json:"receiver_wallet_number" binding:"required"`
	Amount               int    `json:"amount" binding:"required"`
	Description          string `json:"description" `
}
