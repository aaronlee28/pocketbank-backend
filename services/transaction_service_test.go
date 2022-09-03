package services_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//t.Run("Should return response body", func(t *testing.T) {
//	repoReq :=
//	repoRes :=
//	serviceReq :=
//
//	mockRepo := new(mocks.TransactionRepository)
//	mockRepo.On("Payment", &repoReq, 0).Return(&repoRes, nil, nil)
//	transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})
//
//	_, err := transactionService.Payment(&serviceReq, 0)
//	assert.Nil(t, err)
//})
func TestTransactionService_TopupSavings(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		repoReq := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               60000,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "Top Up",
			Status:               "",
		}
		repoRes := models.Transaction{}
		serviceReq := dto.TopupSavingsReq{
			Amount:             60000,
			SenderWalletNumber: 0,
			Description:        "",
		}
		mockRepo := new(mocks.TransactionRepository)
		mockRepo.On("TopupSavings", &repoReq, 0).Return(&repoRes, nil, nil)
		transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})
		_, err := transactionService.TopupSavings(&serviceReq, 0)
		assert.Nil(t, err)
	})

	t.Run("Should return error when amount is less than 50000", func(t *testing.T) {
		repoReq := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               40000,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "Top Up",
			Status:               "",
		}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		serviceReq := dto.TopupSavingsReq{
			Amount:             40000,
			SenderWalletNumber: 0,
			Description:        "",
		}
		mockRepo := new(mocks.TransactionRepository)
		mockRepo.On("TopupSavings", &repoReq, 0).Return(nil, errorRepoRes, nil)
		transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})
		_, err := transactionService.TopupSavings(&serviceReq, 0)
		assert.NotNil(t, err)
	})
	t.Run("Should return error when amount is more than 50000000", func(t *testing.T) {
		repoReq := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               60000000,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "Top Up",
			Status:               "",
		}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		serviceReq := dto.TopupSavingsReq{
			Amount:             60000000,
			SenderWalletNumber: 0,
			Description:        "",
		}
		mockRepo := new(mocks.TransactionRepository)
		mockRepo.On("TopupSavings", &repoReq, 0).Return(nil, nil, errorRepoRes)
		transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})
		_, err := transactionService.TopupSavings(&serviceReq, 0)
		assert.NotNil(t, err)
	})
}

func TestTransactionService_Payment(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		repoReq := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               2000,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "Transfer",
			Status:               "",
		}
		repoRes := models.Transaction{}
		serviceReq := dto.PaymentReq{
			ReceiverAccount: 0,
			Amount:          2000,
			Description:     "",
		}

		mockRepo := new(mocks.TransactionRepository)
		mockRepo.On("Payment", &repoReq, 0).Return(&repoRes, nil, nil)
		transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})

		_, err := transactionService.Payment(&serviceReq, 0)
		assert.Nil(t, err)
	})

	t.Run("Should return error if amount is less than 1000", func(t *testing.T) {
		repoReq := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               999,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "Transfer",
			Status:               "",
		}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		serviceReq := dto.PaymentReq{
			ReceiverAccount: 0,
			Amount:          999,
			Description:     "",
		}

		mockRepo := new(mocks.TransactionRepository)
		mockRepo.On("Payment", &repoReq, 0).Return(nil, errorRepoRes, nil)
		transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})

		_, err := transactionService.Payment(&serviceReq, 0)
		assert.NotNil(t, err)
	})

	t.Run("Should return error if amount is more than 50000000", func(t *testing.T) {
		repoReq := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               51000000,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "Transfer",
			Status:               "",
		}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		serviceReq := dto.PaymentReq{
			ReceiverAccount: 0,
			Amount:          51000000,
			Description:     "",
		}

		mockRepo := new(mocks.TransactionRepository)
		mockRepo.On("Payment", &repoReq, 0).Return(nil, nil, errorRepoRes)
		transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: mockRepo})

		_, err := transactionService.Payment(&serviceReq, 0)
		assert.NotNil(t, err)
	})
}
