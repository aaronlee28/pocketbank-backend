package models

type Favoritecontact struct {
	Id                    int  `json:"id" gorm:"primarykey"`
	UserID                int  `json:"user_id"`
	FavoriteAccountNumber int  `json:"favoriteAccountNumber"`
	Favorite              bool `json:"favorite"`
}
