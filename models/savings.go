package models

type Savings struct {
	Id       int `json:"id" gorm:"primarykey"`
	UserID   int `json:"user_id"`
	Balance  int `json:"balance"`
	Interest int `json:"interest"`
	Tax      int `json:"tax"`
}
