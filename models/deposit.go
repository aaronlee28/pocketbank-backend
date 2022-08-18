package models

type Deposit struct {
	Id            int     `json:"id" gorm:"primarykey"`
	UserID        int     `json:"user_id"`
	DepositNumber int     `json:"deposit_number"`
	Balance       float32 `json:"balance"`
	InterestRate  float32 `json:"interest_rate"`
	Tax           float32 `json:"tax"`
	Interest      float32 `json:"interest"`
}
