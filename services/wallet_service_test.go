package services_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWalletService_TransactionHistory(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		request := &repositories.Query{
			SortBy:     "created_at",
			Sort:       "desc",
			Limit:      "",
			Page:       "",
			Search:     "",
			FilterTime: "74000",
			MinAmount:  "0",
			MaxAmount:  "999999999999",
			Type:       "",
		}
		var repoResponse []models.Transaction
		resp := models.Transaction{
			Id:                   0,
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               0,
			Description:          "",
			CreatedAt:            time.Time{},
			Type:                 "",
			Status:               "",
		}
		tr := dto.TransRes{
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "",
			ReceiverName:         "",
			Amount:               0,
			Type:                 "",
			Status:               "",
			Description:          "",
			CreatedAt:            time.Time{},
		}
		var tra []dto.TransRes
		tra = append(tra, tr)
		response := dto.TransHistoryRes{
			TotalLength:  1,
			Transactions: tra,
		}
		repoResponse = append(repoResponse, resp)
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("TransactionHistory", request, 0).Return(1, &repoResponse, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		res, err := walletService.TransactionHistory(request, 0)
		assert.Nil(t, err)
		assert.Equal(t, res, &response)

	})

	t.Run("Should return error when error != nil", func(t *testing.T) {
		request := &repositories.Query{
			SortBy:     "created_at",
			Sort:       "desc",
			Limit:      "",
			Page:       "",
			Search:     "",
			FilterTime: "74000",
			MinAmount:  "0",
			MaxAmount:  "999999999999",
			Type:       "",
		}
		response := httperror.AppError{
			Message: "users_contact_uindex",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("TransactionHistory", request, 0).Return(0, nil, response)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.TransactionHistory(request, 0)
		assert.NotNil(t, err)
	})
}

func TestWalletService_UserDetails(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		repoResponse := dto.UserDetailsRes{}
		expectedServiceResponse := dto.UserDetailsRes{}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("UserDetails", 0).Return(&repoResponse, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		res, err := walletService.UserDetails(0)
		assert.Nil(t, err)
		assert.Equal(t, res, &expectedServiceResponse)
	})
	t.Run("Should return error when repoo returns error", func(t *testing.T) {
		response := httperror.AppError{
			Message: "users_contact_uindex",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("UserDetails", 0).Return(nil, response)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.UserDetails(0)
		assert.NotNil(t, err)
	})
}
