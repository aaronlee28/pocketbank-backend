package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type AdminService interface {
	AdminUsersList() (*[]dto.UsersListRes, error)
	AdminUserTransaction(q *repositories.Query, id int) (*[]dto.TransRes, error)
	AdminUserDetails(id int) (*dto.UserDetailsRes, error)
	AdminUserReferralDetails(id int) (*dto.UserReferralDetailsRes, error)
	ChangeUserStatus(id int) error
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

func (a *adminService) AdminUsersList() (*[]dto.UsersListRes, error) {

	var result []dto.UsersListRes
	usersList, err := a.adminRepository.AdminUsersList()
	if err != nil {
		return nil, error(httperror.BadRequestError("Internal Server Error", "400"))
	}
	for _, u := range *usersList {
		usr := new(dto.UsersListRes).FromUser(&u)

		result = append(result, *usr)
	}

	return &result, nil
}

func (a *adminService) AdminUserTransaction(q *repositories.Query, id int) (*[]dto.TransRes, error) {

	var result []dto.TransRes
	if q.SortBy == "" {
		q.SortBy = "created_at"
	}
	if q.Sort == "" {
		q.Sort = "desc"
	}
	if q.Limit == "" {
		q.Limit = "10"
	}
	if q.FilterTime == "" {
		q.FilterTime = "74000"
	}
	if q.MinAmount == "" {
		q.MinAmount = "0"
	}
	if q.MaxAmount == "" {
		q.MaxAmount = "999999999"
	}
	t, err := a.adminRepository.AdminUserTransaction(q, id)
	if err != nil {
		return nil, error(httperror.BadRequestError("Bad Request", "400"))
	}
	for _, transaction := range *t {
		tr := new(dto.TransRes).FromTransaction(&transaction)

		result = append(result, *tr)
	}
	return &result, err
}

func (a *adminService) AdminUserDetails(id int) (*dto.UserDetailsRes, error) {

	ret, err := a.adminRepository.AdminUserDetails(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}

	return ret, err
}

func (a *adminService) AdminUserReferralDetails(id int) (*dto.UserReferralDetailsRes, error) {
	var list []int
	ret, err := a.adminRepository.AdminUserReferralDetails(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}
	for _, ids := range *ret {
		list = append(list, ids.UsedByUserID)
	}

	refdetails := &dto.UserReferralDetailsRes{
		Count:       len(list),
		ListOfUsers: list,
	}
	return refdetails, err
}

func (a *adminService) ChangeUserStatus(id int) error {
	err := a.adminRepository.ChangeUserStatus(id)
	if err != nil {
		return error(httperror.BadRequestError("User not found", "400"))
	}
	return err
}
