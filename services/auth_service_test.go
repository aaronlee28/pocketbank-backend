package services_test

import (

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

func TestAuthService_SignIn(t *testing.T) {
	t.Run("should return response body when given email and password", func(t *testing.T) {
		email := "test@gmail.com"

		request := dto.AuthReq{
			Email:    email,
			Password: "test",
		}
		ea := &email
		user := models.User{
			Email:    ea,
			Password: "test",
		}

		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("MatchingCredential", "test@gmail.com", "test").Return(&user, nil, true)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		idToken, err := authService.SignIn(&request)
		assert.Nil(t, err)
		assert.NotNil(t, idToken)
	})
	t.Run("should return error when isUserActive == false", func(t *testing.T) {
		email := "test@gmail.com"

		request := dto.AuthReq{
			Email:    email,
			Password: "test",
		}

		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("MatchingCredential", "test@gmail.com", "test").Return(nil, nil, false)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		idToken, err := authService.SignIn(&request)
		assert.Nil(t, idToken)
		assert.NotNil(t, err)
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		email := "test@gmail.com"

		request := dto.AuthReq{
			Email:    email,
			Password: "test",
		}
		response := httperror.AppError{
			Message: "test error",
		}
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("MatchingCredential", "test@gmail.com", "test").Return(nil, response, true)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		idToken, err := authService.SignIn(&request)
		assert.Nil(t, idToken)
		assert.NotNil(t, err)
	})
}

func TestAuthService_GetCode(t *testing.T) {
	t.Run("should return response body when given email and password", func(t *testing.T) {
		email := "test"
		getCodeRequest := dto.CodeReq{
			Email: email,
		}
		ea := &email
		user := models.User{
			Email: ea,
		}
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("GetCode", "test").Return(&user, 0, nil)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		res, err := authService.GetCode(&getCodeRequest)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should return error when email is incorrect", func(t *testing.T) {
		email := ""
		getCodeRequest := dto.CodeReq{
			Email: email,
		}
		var user *models.User
		user = nil
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("GetCode", "").Return(user, 1, nil)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		res, err := authService.GetCode(&getCodeRequest)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestAuthService_ChangePassword(t *testing.T) {
	t.Run("should return response body when given email and password", func(t *testing.T) {
		changePassReq := dto.ChangePReq{
			Email:       "email@gmail.cxom",
			NewPassword: "password",
			Code:        12345,
		}
		expectedRet := dto.ChangePRes{Success: "Successfully Change Password"}
		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("ChangePassword", &changePassReq).Return(0)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		ret, err := authService.ChangePassword(&changePassReq)
		assert.Nil(t, err)
		assert.NotNil(t, ret)
		assert.Equal(t, &expectedRet, ret)
	})

	t.Run("should return error when repo returns 1", func(t *testing.T) {
		changePassReq := dto.ChangePReq{
			Email:       "email@gmail.cxom",
			NewPassword: "password",
			Code:        12345,
		}

		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("ChangePassword", &changePassReq).Return(1)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		_, err := authService.ChangePassword(&changePassReq)
		assert.ErrorContains(t, err, "Email is not found")
	})

	t.Run("should return error when repo returns 2", func(t *testing.T) {
		changePassReq := dto.ChangePReq{
			Email:       "email@gmail.cxom",
			NewPassword: "password",
			Code:        12345,
		}

		mockRepo := new(mocks.AuthRepository)
		mockRepo.On("ChangePassword", &changePassReq).Return(2)
		authService := services.NewAuthService(&services.ASConfig{
			AuthRepository: mockRepo,
		})

		_, err := authService.ChangePassword(&changePassReq)
		assert.ErrorContains(t, err, "Code Invalid")
	})
}
