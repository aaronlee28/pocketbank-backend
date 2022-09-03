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

func TestAdminService_UserRate(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		serviceReq := dto.ChangeInterestRateReq{}
		repoReq := dto.ChangeInterestRateReq{}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UserRate", 0, &repoReq).Return(nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.UserRate(0, &serviceReq)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		serviceReq := dto.ChangeInterestRateReq{}
		repoReq := dto.ChangeInterestRateReq{}
		errorRepoRes := httperror.AppError{
			Message: "User not found",
		}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UserRate", 0, &repoReq).Return(errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.UserRate(0, &serviceReq)
		assert.NotNil(t, err)
	})
}

func TestAdminService_UsersRate(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		serviceReq := dto.ChangeInterestRateReq{}
		repoReq := dto.ChangeInterestRateReq{}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UsersRate", &repoReq).Return(nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.UsersRate(&serviceReq)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		serviceReq := dto.ChangeInterestRateReq{}
		repoReq := dto.ChangeInterestRateReq{}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UsersRate", &repoReq).Return(errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.UsersRate(&serviceReq)
		assert.NotNil(t, err)
	})
}

func TestAdminService_CreatePromotion(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		serviceReq := dto.PromotionReq{}
		repoReq := dto.PromotionReq{}
		repoRes := models.Promotion{}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("CreatePromotion", &repoReq).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.CreatePromotion(&serviceReq)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		serviceReq := dto.PromotionReq{}
		repoReq := dto.PromotionReq{}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("CreatePromotion", &repoReq).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.CreatePromotion(&serviceReq)
		assert.NotNil(t, err)
	})
}

func TestAdminService_GetPromotion(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		var repoRes []models.Promotion
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("GetPromotion").Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.GetPromotion()
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "error",
		}

		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("GetPromotion").Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.GetPromotion()
		assert.NotNil(t, err)
	})
}

func TestAdminService_UpdatePromotion(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		serviceReq := dto.PatchPromotionReq{}
		repoReq := dto.PatchPromotionReq{}
		repoRes := dto.PatchPromotionReq{}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UpdatePromotion", 0, &repoReq).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})
		_, err := adminService.UpdatePromotion(0, &serviceReq)
		assert.Nil(t, err)
	})
	t.Run("should return error when repo returns error", func(t *testing.T) {
		serviceReq := dto.PatchPromotionReq{}
		repoReq := dto.PatchPromotionReq{}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UpdatePromotion", 0, &repoReq).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})
		_, err := adminService.UpdatePromotion(0, &serviceReq)
		assert.NotNil(t, err)
	})

}

func TestAdminService_DeletePromotion(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		var repoRes models.Promotion
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("DeletePromotion", 0).Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.DeletePromotion(0)
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("DeletePromotion", 0).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})
		_, err := adminService.DeletePromotion(0)
		assert.NotNil(t, err)
	})
}

func TestAdminService_EligibleMerchandiseList(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		var repoRes []models.Merchandise
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("EligibleMerchandiseList").Return(&repoRes, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.EligibleMerchandiseList()
		assert.Nil(t, err)
	})

	t.Run("should return error when repo returns error", func(t *testing.T) {
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("EligibleMerchandiseList").Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})
		_, err := adminService.EligibleMerchandiseList()
		assert.NotNil(t, err)
	})
}

func TestAdminService_MerchandiseStatus(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		serviceReq := dto.MerchandiseStatus{}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("MerchandiseStatus", &serviceReq).Return(nil, nil, 0)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.MerchandiseStatus(&serviceReq)
		assert.Nil(t, err)
	})

	t.Run("should return error if code == 1", func(t *testing.T) {
		serviceReq := dto.MerchandiseStatus{}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("MerchandiseStatus", &serviceReq).Return(nil, nil, 1)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.MerchandiseStatus(&serviceReq)
		assert.NotNil(t, err)
	})

	t.Run("should return error if err1 != nil", func(t *testing.T) {
		serviceReq := dto.MerchandiseStatus{}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("MerchandiseStatus", &serviceReq).Return(&errorRepoRes, nil, 0)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.MerchandiseStatus(&serviceReq)
		assert.NotNil(t, err)
	})
	t.Run("should return error if err2 != nil", func(t *testing.T) {
		serviceReq := dto.MerchandiseStatus{}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("MerchandiseStatus", &serviceReq).Return(nil, &errorRepoRes, 0)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		err := adminService.MerchandiseStatus(&serviceReq)
		assert.NotNil(t, err)
	})
}

func TestAdminService_UpdateMerchStocks(t *testing.T) {
	t.Run("should return response body", func(t *testing.T) {
		req := dto.UpdateMerchStocksReq{}
		ret := models.Merchstock{}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UpdateMerchStocks", &req).Return(&ret, nil)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.UpdateMerchStocks(&req)
		assert.Nil(t, err)
	})

	t.Run("should return response body", func(t *testing.T) {
		req := dto.UpdateMerchStocksReq{}
		errorRepoRes := httperror.AppError{
			Message: "error",
		}
		mockRepo := new(mocks.AdminRepository)
		mockRepo.On("UpdateMerchStocks", &req).Return(nil, &errorRepoRes)
		adminService := services.NewAdminServices(&services.ADSConfig{AdminRepository: mockRepo})

		_, err := adminService.UpdateMerchStocks(&req)
		assert.NotNil(t, err)
	})

}
