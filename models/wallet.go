package models

type Wallet struct {
	Id           int `json:"id" gorm:"primarykey"`
	WalletNumber int `json:"wallet_number"`
	UserID       int `json:"user_id"`
	Balance      int `json:"balance"`
}
