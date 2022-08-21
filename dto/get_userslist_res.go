package dto

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"

type UsersListRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (_ *UsersListRes) FromUser(t *models.User) *UsersListRes {
	return &UsersListRes{
		Id:   t.Id,
		Name: t.Name,
	}
}
