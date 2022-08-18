package models

type Deposit struct {
	Id            int     `json:"id" gorm:"primarykey"`
	UserID        int     `json:"user_id"`
	DepositNumber int     `json:"deposit_number"`
	Balance       float32 `json:"balance"`
	Interest      float32 `json:"interest"`
	Tax           float32 `json:"tax"`
}
