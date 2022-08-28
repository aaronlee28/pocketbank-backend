package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type WalletService interface {
	TransactionHistory(q *repositories.Query, id int) (*[]dto.TransRes, error)
	UserDetails(id int) (*dto.UserDetailsRes, error)
	DepositInfo(id int) (*[]dto.DepositInfoRes, error)
	SavingsInfo(id int) (*dto.SavingsRes, error)
	//PaymentHistory(id int) (*[]dto.PaymentHistoryRes, error)
	FavoriteContact(param *dto.FavoriteContactReq, favoriteid int) (*dto.FavoriteContactRes, error)
	FavoriteContactList(id int) (*[]dto.FavoriteContactRes, error)
	ChangeUserDetails(data *dto.ChangeUserDetailsReqRes, id int) (*dto.ChangeUserDetailsReqRes, error)
}

type walletService struct {
	walletRepository repositories.WalletRepository
}

type WSConfig struct {
	WalletRepository repositories.WalletRepository
}

func NewWalletServices(c *WSConfig) *walletService {
	return &walletService{
		walletRepository: c.WalletRepository,
	}
}

func (a *walletService) TransactionHistory(q *repositories.Query, id int) (*[]dto.TransRes, error) {

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
	t, err := a.walletRepository.TransactionHistory(q, id)
	if err != nil {
		return nil, error(httperror.BadRequestError("Bad Request", "400"))
	}
	for _, transaction := range *t {
		tr := new(dto.TransRes).FromTransaction(&transaction)
		result = append(result, *tr)
	}
	return &result, err
}

func (a *walletService) UserDetails(id int) (*dto.UserDetailsRes, error) {

	ret, err := a.walletRepository.UserDetails(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}

	return ret, err
}

func (a *walletService) DepositInfo(id int) (*[]dto.DepositInfoRes, error) {
	var res []dto.DepositInfoRes
	ret, err := a.walletRepository.DepositInfo(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}

	for _, d := range *ret {
		tr := new(dto.DepositInfoRes).FromDepositInfo(&d)

		res = append(res, *tr)
	}
	return &res, err
}

func (a *walletService) SavingsInfo(id int) (*dto.SavingsRes, error) {
	ret, err := a.walletRepository.SavingsInfo(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}
	res := &dto.SavingsRes{
		UserID:   ret.UserID,
		Balance:  ret.Balance,
		Interest: ret.Interest,
		Tax:      ret.Tax,
	}
	return res, err
}

//func (a *walletService) PaymentHistory(id int) (*[]dto.PaymentHistoryRes, error) {
//	var res []dto.PaymentHistoryRes
//	ret, err := a.walletRepository.PaymentHistory(id)
//	if err != nil {
//		return nil, error(httperror.BadRequestError("INTERNAL SERVER ERROR", "400"))
//	}
//
//	for _, t := range *ret {
//		tr := new(dto.PaymentHistoryRes).FromTransaction(&t)
//
//		res = append(res, *tr)
//	}
//	return &res, err
//}

func (a *walletService) FavoriteContact(favoriteid *dto.FavoriteContactReq, selfid int) (*dto.FavoriteContactRes, error) {
	fid := favoriteid.FavoriteUserID
	ret, err := a.walletRepository.FavoriteContact(fid, selfid)

	if err != nil || fid == selfid {
		return nil, error(httperror.BadRequestError("INTERNAL SERVER ERROR", "400"))
	}
	res := new(dto.FavoriteContactRes).FromFavoritecontact(ret)

	return res, err
}
func (a *walletService) FavoriteContactList(id int) (*[]dto.FavoriteContactRes, error) {
	var res []dto.FavoriteContactRes
	ret, err := a.walletRepository.FavoriteContactList(id)

	if err != nil {
		return nil, error(httperror.BadRequestError("INTERNAL SERVER ERROR", "400"))
	}
	for _, t := range *ret {
		tmp := new(dto.FavoriteContactRes).FromFavoritecontact(&t)

		res = append(res, *tmp)
	}

	return &res, err
}

func (a *walletService) ChangeUserDetails(data *dto.ChangeUserDetailsReqRes, id int) (*dto.ChangeUserDetailsReqRes, error) {

	ret, err := a.walletRepository.ChangeUserDetails(data, id)
	if err != nil {
		return nil, error(httperror.BadRequestError("INTERNAL SERVER ERROR", "400"))
	}
	return ret, err
}
