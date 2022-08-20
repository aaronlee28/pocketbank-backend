package dto

type FavoriteContactReq struct {
	FavoriteUserID int `json:"favorite_user_id" binding:"required"`
}
