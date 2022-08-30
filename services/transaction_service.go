package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
)

type TransactionService interface {
	TopupSavings(req *dto.TopupSavingsReq, id int) (*dto.TopupSavingsRes, error)
	Payment(req *dto.PaymentReq, id int) (*dto.PaymentRes, error)
	RunCronJobs()
	TopupDeposit(req *dto.TopupDepositReq, id int) (*dto.DepositRes, error)
	TopUpQr(req *dto.TopUpQr, id int) (*dto.TopUpQr, error)
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

func (a *transactionService) TopupSavings(req *dto.TopupSavingsReq, id int) (*dto.TopupSavingsRes, error) {

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
		Type:               "Top Up",
	}

	transaction, err1, err2 := a.transactionRepository.TopupSavings(t, id)
	if err1 != nil || err2 != nil {
		return nil, error(httperror.BadRequestError("Bad Request", ""))
	}
	ret := &dto.TopupSavingsRes{
		Amount:      transaction.Amount,
		Description: transaction.Description,
	}
	return ret, nil
}

func (a *transactionService) Payment(req *dto.PaymentReq, id int) (*dto.PaymentRes, error) {

	if req.Amount < 1000 {
		return nil, error(httperror.BadRequestError("Minimum Amount is Rp.1000", "402"))
	}
	if req.Amount > 50000000 {
		return nil, error(httperror.BadRequestError("Maximum Amount is Rp.50000000", "402"))
	}
	t := &models.Transaction{
		SenderWalletNumber:   id,
		ReceiverWalletNumber: req.ReceiverAccount,
		Amount:               req.Amount,
		Description:          req.Description,
		Type:                 "Transfer",
	}
	payment, err := a.transactionRepository.Payment(t, id)

	ret := &dto.PaymentRes{
		SenderAccount:   payment.SenderWalletNumber,
		ReceiverAccount: payment.ReceiverWalletNumber,
		SenderName:      payment.SenderName,
		ReceiverName:    payment.ReceiverName,
		Amount:          payment.Amount,
		Status:          payment.Status,
		Description:     payment.Description,
	}
	return ret, err
}

func (a *transactionService) RunCronJobs() {

	a.transactionRepository.RunCronJobs()

}

func (a *transactionService) TopupDeposit(req *dto.TopupDepositReq, id int) (*dto.DepositRes, error) {

	if req.Amount < 1000000 {
		return nil, error(httperror.BadRequestError("Minimum Amount is Rp.1000000", "400"))
	}

	deposit, err1 := a.transactionRepository.TopupDeposit(req, id)
	if err1 != nil || deposit == nil {
		return nil, error(httperror.BadRequestError("Insufficient Balance", "401"))
	}

	ret := &dto.DepositRes{
		Amount:      deposit.Balance,
		Duration:    deposit.Duration,
		AutoDeposit: deposit.AutoDeposit,
	}
	return ret, nil
}

func (a *transactionService) TopUpQr(req *dto.TopUpQr, id int) (*dto.TopUpQr, error) {
	ret, err1 := a.transactionRepository.TopUpQr(req, id)
	if err1 != nil || ret == nil {
		return nil, error(httperror.BadRequestError("User Not Found", "401"))
	}

	return ret, nil
}
