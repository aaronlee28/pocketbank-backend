package dto

type TopUpQr struct {
	SenderWalletNumber int     `json:"senderWalletNumber"`
	Amount             float32 `json:"amount"`
	Description        string  `json:"description"`
}
