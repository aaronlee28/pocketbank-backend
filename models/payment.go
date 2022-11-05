package models

import "time"

type Payment struct {
	Id              int       `json:"id" gorm:"primarykey"`
	SenderAccount   int       `json:"senderAccount"`
	ReceiverAccount int       `json:"receiverAccount"`
	Amount          float32   `json:"amount"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
