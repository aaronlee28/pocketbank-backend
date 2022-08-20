package models

type Favoritecontact struct {
	Id             int `json:"id" gorm:"primarykey"`
	UserID         int `json:"user_id"`
	FavoriteUserID int `json:"favorite_user_id"`
}
