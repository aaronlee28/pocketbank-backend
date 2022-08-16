package models

import "time"

type Transaction struct {
	Id                   int       `json:"id" gorm:"primarykey"`
	SenderWalletNumber   int       `json:"sender_wallet_number"`
	ReceiverWalletNumber int       `json:"receiver_wallet_number"`
	Amount               int       `json:"amount"`
	Description          string    `json:"description"`
	SourceOfFundID       int       `json:"source_of_fund_id"`
	CreatedAt            time.Time `json:"created_at"`
}
