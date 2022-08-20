package dto

type FavoriteContactReq struct {
	UserID         int `json:"user_id" binding:"required"`
	FavoriteUserID int `json:"favorite_user_id" binding:"required"`
}
