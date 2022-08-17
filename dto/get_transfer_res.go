package dto

type TransferRes struct {
	ReceiverWalletNumber int     `json:"receiver_wallet_number"`
	Amount               float32 `json:"amount"`
	Description          string  `json:"description"`
}
