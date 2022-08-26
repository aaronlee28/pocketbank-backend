package models

import "time"

type Transaction struct {
	Id                   int       `json:"id" gorm:"primarykey"`
	SenderWalletNumber   int       `json:"senderWalletNumber"`
	ReceiverWalletNumber int       `json:"receiverWalletNumber"`
	ReceiverName         string    `json:"name"`
	Amount               float32   `json:"amount"`
	Description          string    `json:"description"`
	CreatedAt            time.Time `json:"createdAt"`
	Type                 string    `json:"type"`
	Status               string    `json:"status"`
}
