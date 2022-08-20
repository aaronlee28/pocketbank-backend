package models

import "time"

type Deposit struct {
	Id            int       `json:"id" gorm:"primarykey"`
	UserID        int       `json:"user_id"`
	DepositNumber int       `json:"deposit_number"`
	Balance       float32   `json:"balance"`
	InterestRate  float32   `json:"interest_rate"`
	Tax           float32   `json:"tax"`
	Interest      float32   `json:"interest"`
	AutoDeposit   bool      `json:"auto_deposit"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
