package services_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_Register(t *testing.T) {
	t.Run("should return response body when given email and password", func(t *testing.T) {
		email := "test"
		request := dto.RegReq{
			Name:           "test",
			Contact:        "test",
			Password:       "test",
			Email:          email,
			ReferralNumber: 0,
			Photo:          nil,
		}
		ea := &email
		var user = models.User{
			Id:             0,
			Name:           "test",
			Contact:        "test",
			Password:       "test",
			Code:           0,
			Email:          ea,
			ReferralNumber: nil,
			IsActive:       false,
		}
		response := dto.RegRes{
			Email:   "test",
			Name:    "test",
			Contact: "test",
		}

		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("Register", &user, 0).Return(&user, nil)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})
		res, err := authService.Register(&request)
		assert.Nil(t, err)
		assert.Equal(t, res, &response)
	})

	t.Run("should return error when repo returns users_contact_uindex", func(t *testing.T) {
		email := ""

		request := dto.RegReq{
			Name:     "",
			Contact:  "",
			Password: "",
			Email:    "",
		}
		ea := &email

		var user = models.User{
			Name:     "",
			Contact:  "",
			Password: "",
			Email:    ea,
			IsActive: false,
		}
		response := httperror.AppError{
			Message: "users_contact_uindex",
		}
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("Register", &user, 0).Return(nil, response)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})
		res, err := authService.Register(&request)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
	t.Run("should return error when repo returns users_email_uindex", func(t *testing.T) {
		email := ""

		request := dto.RegReq{
			Name:     "",
			Contact:  "",
			Password: "",
			Email:    "",
		}
		ea := &email

		var user = models.User{
			Name:     "",
			Contact:  "",
			Password: "",
			Email:    ea,
			IsActive: false,
		}
		response := httperror.AppError{
			Message: "users_email_uindex",
		}
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("Register", &user, 0).Return(nil, response)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})
		res, err := authService.Register(&request)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
	t.Run("should return error when repo returns record not found", func(t *testing.T) {
		email := ""

		request := dto.RegReq{
			Name:     "",
			Contact:  "",
			Password: "",
			Email:    "",
		}
		ea := &email

		var user = models.User{
			Name:     "",
			Contact:  "",
			Password: "",
			Email:    ea,
			IsActive: false,
		}
		response := httperror.AppError{
			Message: "record not found",
		}
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("Register", &user, 0).Return(nil, response)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})
		res, err := authService.Register(&request)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
