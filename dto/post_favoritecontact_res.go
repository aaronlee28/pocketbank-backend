package dto

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"

type FavoriteContactRes struct {
	UserID                int `json:"UserID" binding:"required"`
	FavoriteAccountNumber int `json:"favoriteAccountNumber" binding:"required"`
}

func (_ *FavoriteContactRes) FromFavoritecontact(t *models.Favoritecontact) *FavoriteContactRes {
	return &FavoriteContactRes{
		UserID:                t.UserID,
		FavoriteAccountNumber: t.FavoriteAccountNumber,
	}
}
