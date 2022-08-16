package models

type Wallet struct {
	Id            int `json:"id" gorm:"primarykey"`
	WalletNumber  int `json:"wallet_number"`
	SavingsNumber int `json:"savings_number"`
	DepositNumber int `json:"deposit_number"`
	UserID        int `json:"user_id"`
	Balance       int `json:"balance"`
}
