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

func (a *walletService) Transaction(q *repositories.Query, id int) (*[]dto.TransRes, error) {

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
