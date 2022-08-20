package dto

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"

type FavoriteContactRes struct {
	UserID         int `json:"user_id" binding:"required"`
	FavoriteUserID int `json:"favorite_user_id" binding:"required"`
}

func (_ *FavoriteContactRes) FromFavoritecontact(t *models.Favoritecontact) *FavoriteContactRes {
	return &FavoriteContactRes{
		UserID:         t.UserID,
		FavoriteUserID: t.FavoriteUserID,
	}
}
