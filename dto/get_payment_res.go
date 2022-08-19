package dto

type PaymentRes struct {
	ReceiverWalletNumber int     `json:"receiver_wallet_number"`
	Amount               float32 `json:"amount"`
	Description          string  `json:"description"`
}
