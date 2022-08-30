package dto

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"

type UsersListRes struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ReferralCode int    `json:"referralCode"`
	IsActive     bool   `json:"IsActive"`
}

func (_ *UsersListRes) FromUser(t *models.User) *UsersListRes {
	return &UsersListRes{
		Id:           t.Id,
		Name:         t.Name,
		ReferralCode: *t.ReferralNumber,
		IsActive:     t.IsActive,
	}
}
