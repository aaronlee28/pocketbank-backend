package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type TransactionService interface {
	TopupSavings(req *dto.TopupReq, id int) (*dto.TopupRes, error)
	Transfer(req *dto.TransferReq, id int) (*dto.TransferRes, error)
	RunCronJobs()
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
}

type TSConfig struct {
	TransactionRepository repositories.TransactionRepository
}

func NewTransactionServices(c *TSConfig) *transactionService {
	return &transactionService{
		transactionRepository: c.TransactionRepository,
	}
}

func (a *transactionService) TopupSavings(req *dto.TopupReq, id int) (*dto.TopupRes, error) {

	if req.Amount < 50000 {
		return nil, error(httperror.BadRequestError("Minimum Amount is Rp.50000", "400"))
	}
	if req.Amount > 50000000 {
		return nil, error(httperror.BadRequestError("Maximum Amount is Rp.50000000", "401"))
	}
	t := &models.Transaction{
		Amount:             req.Amount,
		SenderWalletNumber: req.SenderWalletNumber,
		Description:        req.Description,
	}

	transaction, err1, err2 := a.transactionRepository.TopupSavings(t, id)
	if err1 != nil || err2 != nil {
		return nil, error(httperror.BadRequestError("Bad Request", ""))
	}
	ret := &dto.TopupRes{
		Amount:      transaction.Amount,
		Description: transaction.Description,
	}
	return ret, nil
}

func (a *transactionService) Transfer(req *dto.TransferReq, id int) (*dto.TransferRes, error) {

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
	}

	transaction, err1, err2, err3 := a.transactionRepository.Transfer(t, id)
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

func (a *transactionService) RunCronJobs() {

	a.transactionRepository.RunCronJobs()

}
