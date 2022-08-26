package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type AdminService interface {
	AdminUsersList() (*[]dto.UsersListRes, error)
	AdminUserTransaction(q *repositories.Query, id int) (*[]dto.TransRes, error)
	AdminUserDetails(id int) (*dto.UserDetailsRes, error)
	AdminUserReferralDetails(id int) (*dto.UserReferralDetailsRes, error)
	ChangeUserStatus(id int) error
	Merchandise(id int) (*models.Merchandise, error)
	UserDepositInfo(id int) (*dto.UserDepositInfo, error)
	UserRate(id int, data *dto.ChangeInterestRateReq) error
	UsersRate(data *dto.ChangeInterestRateReq) error
	CreatePromotion(data *dto.PromotionReq) (*dto.PromotionReq, error)
	GetPromotion() (*[]models.Promotion, error)
	UpdatePromotion(id int, data *dto.PatchPromotionReq) (*dto.PatchPromotionReq, error)
	DeletePromotion(id int) (*models.Promotion, error)
	EligibleMerchandiseList() (*[]models.Merchandise, error)
	MerchandiseStatus(data *dto.MerchandiseStatus) error
	UpdateMerchStocks(data *dto.UpdateMerchStocksReq) (*models.Merchstock, error)
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

func (a *adminService) Merchandise(id int) (*models.Merchandise, error) {
	m, err := a.adminRepository.Merchandise(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}
	return m, err
}

func (a *adminService) UserDepositInfo(id int) (*dto.UserDepositInfo, error) {
	deposits, err := a.adminRepository.UserDepositInfo(id)
	var totalDepositBalance float32
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}
	for _, d := range *deposits {
		totalDepositBalance = totalDepositBalance + d.Balance
	}
	ret := &dto.UserDepositInfo{
		UserID:      id,
		Balance:     totalDepositBalance,
		AllDeposits: deposits,
	}
	return ret, err
}

func (a *adminService) UserRate(id int, data *dto.ChangeInterestRateReq) error {

	err := a.adminRepository.UserRate(id, data)

	if err != nil {
		return error(httperror.BadRequestError("User not found", "400"))
	}

	return err
}

func (a *adminService) UsersRate(data *dto.ChangeInterestRateReq) error {

	err := a.adminRepository.UsersRate(data)

	if err != nil {
		return error(httperror.BadRequestError("User not found", "400"))
	}

	return err
}

func (a *adminService) CreatePromotion(data *dto.PromotionReq) (*dto.PromotionReq, error) {

	p, err := a.adminRepository.CreatePromotion(data)

	if err != nil {
		return nil, error(httperror.BadRequestError("Failed to Create Promotion", "400"))
	}
	ret := &dto.PromotionReq{
		Title: p.Title,
		Photo: p.Photo,
	}
	return ret, err
}

func (a *adminService) GetPromotion() (*[]models.Promotion, error) {

	p, err := a.adminRepository.GetPromotion()
	if err != nil {
		return nil, error(httperror.BadRequestError("Failed to Create Promotion", "400"))
	}

	return p, err
}

func (a *adminService) UpdatePromotion(id int, data *dto.PatchPromotionReq) (*dto.PatchPromotionReq, error) {

	p, err := a.adminRepository.UpdatePromotion(id, data)
	if err != nil {
		return nil, error(httperror.BadRequestError("Failed to Create Promotion", "400"))
	}

	return p, err
}

func (a *adminService) DeletePromotion(id int) (*models.Promotion, error) {

	p, err := a.adminRepository.DeletePromotion(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("Failed to Create Promotion", "400"))
	}

	return p, err
}

func (a *adminService) EligibleMerchandiseList() (*[]models.Merchandise, error) {

	p, err := a.adminRepository.EligibleMerchandiseList()
	if err != nil {
		return nil, error(httperror.BadRequestError("Failed to Create Promotion", "400"))
	}

	return p, err
}

func (a *adminService) MerchandiseStatus(data *dto.MerchandiseStatus) error {

	err1, err2, code := a.adminRepository.MerchandiseStatus(data)
	if code == 1 {
		return error(httperror.BadRequestError("Merchandise stock insufficient", "400"))
	}
	if err1 != nil {
		return error(httperror.BadRequestError("User Id Not Found", "400"))
	}
	if err2 != nil {
		return error(httperror.BadRequestError("Failed To Update Delivery", "400"))
	}

	return nil
}

func (a *adminService) UpdateMerchStocks(data *dto.UpdateMerchStocksReq) (*models.Merchstock, error) {

	p, err := a.adminRepository.UpdateMerchStocks(data)
	if err != nil {
		return nil, error(httperror.BadRequestError("Merchandise could not be found", "400"))
	}

	return p, err
}
