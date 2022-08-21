package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type AdminService interface {
	UsersList() (*[]dto.UsersListRes, error)
}

type adminService struct {
	adminRepository repositories.AdminRepository
}

type ADSConfig struct {
	AdminRepository repositories.AdminRepository
}

func NewAdminServices(c *ADSConfig) *adminService {
	return &adminService{
		adminRepository: c.AdminRepository,
	}
}

func (a *adminService) UsersList() (*[]dto.UsersListRes, error) {

	var result []dto.UsersListRes
	usersList, err := a.adminRepository.UsersList()
	if err != nil {
		return nil, error(httperror.BadRequestError("Internal Server Error", "400"))
	}
	for _, u := range *usersList {
		usr := new(dto.UsersListRes).FromUser(&u)

		result = append(result, *usr)
	}

	return &result, nil
}
