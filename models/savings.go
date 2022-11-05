package models

type Savings struct {
	Id            int     `json:"id" gorm:"primarykey"`
	UserID        int     `json:"user_id"`
	SavingsNumber int     `json:"savingsNumber"`
	Balance       float32 `json:"balance"`
	Interest      float32 `json:"interest"`
	Tax           float32 `json:"tax"`
}
