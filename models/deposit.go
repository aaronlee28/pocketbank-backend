package models

import "time"

type Deposit struct {
	Id            int       `json:"id" gorm:"primarykey"`
	UserID        int       `json:"user_id"`
	DepositNumber int       `json:"depositNumber"`
	Balance       float32   `json:"balance"`
	InterestRate  float32   `json:"interestRate"`
	Tax           float32   `json:"tax"`
	Interest      float32   `json:"interest"`
	AutoDeposit   bool      `json:"autoDeposit"`
	Duration      int       `json:"duration"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
}
