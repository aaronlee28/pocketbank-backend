package models

type Merchandise struct {
	Id            int     `json:"id" gorm:"primarykey"`
	UserID        int     `json:"userID"`
	TotalTransfer float32 `json:"totalTransfer"`
	Pen           bool    `json:"pen"`
	Umbrella      bool    `json:"umbrella"`
	CardHolder    bool    `json:"cardHolder"`
}
