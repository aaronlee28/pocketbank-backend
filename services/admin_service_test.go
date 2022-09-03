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
)

func TestAdminService_AdminUsersList(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		var repoRes []models.User
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUsersList").Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})
		_, err := adminService.AdminUsersList()
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "Internal Server Error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUsersList").Return(nil, errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUsersList()
		assert.NotNil(t, err)

	})
}

func TestAdminService_AdminUserTransaction(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		serviceReq := repositories.Query{}
		repoReq := repositories.Query{
			SortBy:     "created_at",
			Sort:       "desc",
			Limit:      "",
			Page:       "",
			Search:     "",
			FilterTime: "74000",
			MinAmount:  "0",
			MaxAmount:  "999999999",
			Type:       "",
		}
		var repoRes []models.Transaction

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUserTransaction", &repoReq, 0).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUserTransaction(&serviceReq, 0)
		assert.Nil(t, err)
	})

	t.Run("should return error if repo returns error", func(t *testing.T) {
		serviceReq := repositories.Query{}
		repoReq := repositories.Query{
			SortBy:     "created_at",
			Sort:       "desc",
			Limit:      "",
			Page:       "",
			Search:     "",
			FilterTime: "74000",
			MinAmount:  "0",
			MaxAmount:  "999999999",
			Type:       "",
		}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUserTransaction", &repoReq, 0).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUserTransaction(&serviceReq, 0)
		assert.NotNil(t, err)

	})
}

func TestAdminService_AdminUserDetails(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		repoRes := dto.UserDetailsRes{}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUserDetails", 0).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUserDetails(0)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "User not found",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUserDetails", 0).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUserDetails(0)
		assert.NotNil(t, err)
	})
}

func TestAdminService_AdminUserReferralDetails(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		var repoRes []models.Referral
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUserReferralDetails", 0).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUserReferralDetails(0)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "User not found",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("AdminUserReferralDetails", 0).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.AdminUserReferralDetails(0)
		assert.NotNil(t, err)
	})
}

func TestAdminService_ChangeUserStatus(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("ChangeUserStatus", 0).Return(nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.ChangeUserStatus(0)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "User not found",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("ChangeUserStatus", 0).Return(errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.ChangeUserStatus(0)
		assert.NotNil(t, err)
	})

}

func TestAdminService_Merchandise(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		repoRes := models.Merchandise{}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("Merchandise", 0).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.Merchandise(0)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "Merchandise not found",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("Merchandise", 0).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.Merchandise(0)
		assert.NotNil(t, err)
	})
}

func TestAdminService_UserDepositInfo(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		var repoRes []models.Deposit

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UserDepositInfo", 0).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.UserDepositInfo(0)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "User not found",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UserDepositInfo", 0).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.UserDepositInfo(0)
		assert.NotNil(t, err)
	})
}
