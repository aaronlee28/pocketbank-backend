package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type WalletService interface {
	Topup(req *dto.TopupReq, id int) (*dto.TopupRes, error)
	Transaction(q *repositories.Query, id int) (*[]dto.TransRes, error)
	Transfer(req *dto.TransferReq, id int) (*dto.TransferRes, error)
	UserDetails(id int) (*dto.UserDetailsRes, error)
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

func (a *walletService) Topup(req *dto.TopupReq, id int) (*dto.TopupRes, error) {
	var desc string
	if req.SourceOfFundsID == 1 {
		desc = "Top Up From Bank Transfer"
	}
	if req.SourceOfFundsID == 2 {
		desc = "Top Up From Credit Card"
	}
	if req.SourceOfFundsID == 3 {
		desc = "Top Up From Cash"
	}

	if req.Amount < 50000 {
		return nil, error(httperror.BadRequestError("Minimum Amount is Rp.50000", "400"))
	}
	if req.Amount > 50000000 {
		return nil, error(httperror.BadRequestError("Maximum Amount is Rp.50000000", "401"))
	}
	t := &models.Transaction{
		Amount:         req.Amount,
		SourceOfFundID: req.SourceOfFundsID,
		Description:    desc,
	}

	transaction, err1, err2 := a.walletRepository.Topup(t, id)
	if err1 != nil || err2 != nil {
		return nil, error(httperror.BadRequestError("Bad Request", ""))
	}
	ret := &dto.TopupRes{
		Amount:      transaction.Amount,
		Description: transaction.Description,
	}
	return ret, nil
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
	t, err := a.walletRepository.Transaction(q, id)
	if err != nil {
		return nil, error(httperror.BadRequestError("Bad Request", "400"))
	}
	for _, transaction := range *t {
		tr := new(dto.TransRes).FromTransaction(&transaction)
		if transaction.SourceOfFundID == 1 {
			tr.SourceOfFund = "Bank Transfer"
		}
		if transaction.SourceOfFundID == 2 {
			tr.SourceOfFund = "Credit Card"
		}
		if transaction.SourceOfFundID == 3 {
			tr.SourceOfFund = "Cash"
		}

		result = append(result, *tr)
	}
	return &result, err
}

func (a *walletService) Transfer(req *dto.TransferReq, id int) (*dto.TransferRes, error) {

	if req.Amount < 1000 {
		return nil, error(httperror.BadRequestError("Minimum Amount is Rp.1000", ""))
	}
	if req.Amount > 50000000 {
		return nil, error(httperror.BadRequestError("Maximum Amount is Rp.50000000", ""))
	}
	if len(req.Description) > 35 {
		return nil, error(httperror.BadRequestError("Description Maximum Length is 35 Characters", "402"))
	}

	t := &models.Transaction{
		ReceiverWalletNumber: req.ReceiverWalletNumber,
		Amount:               req.Amount,
		Description:          req.Description,
		SourceOfFundID:       1,
	}

	transaction, err1, err2, err3 := a.walletRepository.Transfer(t, id)
	if err1 != nil {
		return nil, error(httperror.BadRequestError("Insufficient Balance", ""))
	}
	if err2 != nil {
		return nil, error(httperror.BadRequestError("Target Wallet Not Found", ""))
	}
	if err3 != nil {
		return nil, error(httperror.BadRequestError("Should Not Be Able To Transfer with Someone Else's Wallet", ""))
	}

	ret := &dto.TransferRes{
		ReceiverWalletNumber: transaction.ReceiverWalletNumber,
		Amount:               transaction.Amount,
		Description:          transaction.Description,
	}
	return ret, nil
}

func (a *walletService) UserDetails(id int) (*dto.UserDetailsRes, error) {

	ret, err := a.walletRepository.UserDetails(id)
	if err != nil {
		return nil, error(httperror.BadRequestError("User not found", "400"))
	}

	return ret, err
}
