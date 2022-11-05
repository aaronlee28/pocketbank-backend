package models

type Wallet struct {
	Id            int     `json:"id" gorm:"primarykey"`
	WalletNumber  int     `json:"walletNumber"`
	DepositNumber int     `json:"depositNumber"`
	UserID        int     `json:"userID"`
	Balance       float32 `json:"balance"`
}
