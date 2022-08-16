package transfer

type TransferRes struct {
	ReceiverWalletNumber int    `json:"receiver_wallet_number"`
	Amount               int    `json:"amount"`
	Description          string `json:"description"`
}
