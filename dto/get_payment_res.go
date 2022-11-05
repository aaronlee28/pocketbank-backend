package dto

type PaymentRes struct {
	SenderAccount   int     `json:"senderAccount"`
	ReceiverAccount int     `json:"receiverAccount"`
	SenderName      string  `json:"senderName"`
	ReceiverName    string  `json:"receiverName"`
	Amount          float32 `json:"amount"`
	Status          string  `json:"status"`
	Description     string  `json:"description"`
}
