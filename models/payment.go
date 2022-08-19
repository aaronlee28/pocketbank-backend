package models

import "time"

type Payment struct {
	Id              int       `json:"id" gorm:"primarykey"`
	SenderAccount   int       `json:"sender_account"`
	ReceiverAccount int       `json:"receiver_account"`
	Amount          float32   `json:"amount"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
