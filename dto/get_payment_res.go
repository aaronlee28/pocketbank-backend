package dto

type PaymentRes struct {
	SenderAccount   int     `json:"sender_account"`
	ReceiverAccount int     `json:"receiver_account"`
	Amount          float32 `json:"amount"`
	Status          string  `json:"status"`
	Description     string  `json:"Description"`
}
