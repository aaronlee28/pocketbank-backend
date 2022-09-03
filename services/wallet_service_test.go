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

func TestWalletService_DepositInfo(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		var repoResponse []models.Deposit
		dep := models.Deposit{}
		repoResponse = append(repoResponse, dep)
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("DepositInfo", 0).Return(&repoResponse, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.DepositInfo(0)
		assert.Nil(t, err)
	})

	t.Run("Should return error when repo returns error", func(t *testing.T) {
		response := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("DepositInfo", 0).Return(nil, response)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.DepositInfo(0)
		assert.NotNil(t, err)
	})
}

func TestWalletService_SavingsInfo(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {

		repoResponse := models.Savings{}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("SavingsInfo", 0).Return(&repoResponse, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.SavingsInfo(0)
		assert.Nil(t, err)
	})

	t.Run("Should return error when repo returns error", func(t *testing.T) {
		response := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("SavingsInfo", 0).Return(nil, response)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.SavingsInfo(0)
		assert.NotNil(t, err)
	})
}

func TestWalletService_FavoriteContact(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		serviceRequest := dto.FavoriteContactReq{}
		repoResponse := models.Favoritecontact{}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("FavoriteContact", 0, 1).Return(&repoResponse, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.FavoriteContact(&serviceRequest, 1)
		assert.Nil(t, err)
	})

	t.Run("Should return error when repo returns error", func(t *testing.T) {
		serviceRequest := dto.FavoriteContactReq{}

		errorRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("FavoriteContact", 0, 0).Return(nil, errorRes)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.FavoriteContact(&serviceRequest, 0)
		assert.NotNil(t, err)
	})
}

func TestWalletService_FavoriteContactList(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		resp := models.Favoritecontact{}
		var repoResponse []models.Favoritecontact
		repoResponse = append(repoResponse, resp)
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("FavoriteContactList", 0).Return(&repoResponse, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.FavoriteContactList(0)
		assert.Nil(t, err)
	})
	t.Run("Should return error when repo returns error", func(t *testing.T) {
		errorRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("FavoriteContactList", 0).Return(nil, errorRes)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.FavoriteContactList(0)
		assert.NotNil(t, err)
	})
}

func TestWalletService_ChangeUserDetails(t *testing.T) {
	t.Run("Should return response body", func(t *testing.T) {
		serviceRequest := dto.ChangeUserDetailsReqRes{}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("ChangeUserDetails", &serviceRequest, 1).Return(&serviceRequest, nil)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.ChangeUserDetails(&serviceRequest, 1)
		assert.Nil(t, err)
	})
	t.Run("Should return error when repo returns error", func(t *testing.T) {
		serviceRequest := dto.ChangeUserDetailsReqRes{}

		errorRes := httperror.AppError{
			Message: "users_contact_uindex",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("ChangeUserDetails", &serviceRequest, 1).Return(nil, errorRes)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.ChangeUserDetails(&serviceRequest, 1)
		assert.NotNil(t, err)
	})
	t.Run("Should return error when repo returns error", func(t *testing.T) {
		serviceRequest := dto.ChangeUserDetailsReqRes{}

		errorRes := httperror.AppError{
			Message: "users_email_uindex",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("ChangeUserDetails", &serviceRequest, 1).Return(nil, errorRes)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.ChangeUserDetails(&serviceRequest, 1)
		assert.NotNil(t, err)
	})
	t.Run("Should return error when repo returns error", func(t *testing.T) {
		serviceRequest := dto.ChangeUserDetailsReqRes{}

		errorRes := httperror.AppError{
			Message: "record not found",
		}
		mockRepo := new(mocks.WalletRepository)
		mockRepo.On("ChangeUserDetails", &serviceRequest, 1).Return(nil, errorRes)
		walletService := services.NewWalletServices(&services.WSConfig{
			WalletRepository: mockRepo,
		})
		_, err := walletService.ChangeUserDetails(&serviceRequest, 1)
		assert.NotNil(t, err)
	})
}
