package handlers_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/server"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/testutils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandler_Register(t *testing.T) {
	t.Run("should return user when user is registering", func(t *testing.T) {
		request := dto.RegReq{
			Name:           "",
			Email:          "",
			Contact:        "",
			Password:       "",
			ReferralNumber: 0,
			Photo:          nil,
		}
		response := &dto.RegRes{
			Name:    "test",
			Email:   "test",
			Contact: "test",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 201,
			Message:    "Created",
			Data:       response,
		}
		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("Register", &request).Return(response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/register", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.RegReq{
			Name:           "",
			Email:          "",
			Contact:        "",
			Password:       "",
			ReferralNumber: 0,
			Photo:          nil,
		}
		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("Register", request).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPost, "/register", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})

}

func TestHandler_SignIn(t *testing.T) {
	t.Run("should return token when sign in with id", func(t *testing.T) {
		request := dto.AuthReq{
			Email:    "test@gmail.com",
			Password: "testpassword",
		}
		response := dto.TokenRes{
			IDToken: "testing token",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}

		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("SignIn", &request).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/signin", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
	t.Run("should return error when signin with invalid credentials", func(t *testing.T) {
		request := dto.AuthReq{
			Email:    "test@gmail.com",
			Password: "testpassword",
		}
		response := httperror.AppError{
			Message: "User Account Inactive",
		}
		m := map[string]string{}
		m["error"] = "User Account Inactive"

		responseError := m

		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("SignIn", &request).Return(nil, response)

		res, _ := json.Marshal(&responseError)
		req, _ := http.NewRequest(http.MethodPost, "/signin", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when result is nil", func(t *testing.T) {
		request := dto.AuthReq{
			Email:    "test@gmail.com",
			Password: "testpassword",
		}

		m := map[string]string{}
		m["error"] = "users not found"

		responseError := m

		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("SignIn", &request).Return(nil, nil)

		res, _ := json.Marshal(&responseError)
		req, _ := http.NewRequest(http.MethodPost, "/signin", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
}

func TestHandler_GetCode(t *testing.T) {
	t.Run("should return code when email is given", func(t *testing.T) {
		request := dto.CodeReq{Email: "test@test.com"}
		response := dto.CodeRes{Code: 12345}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 201,
			Message:    "Created",
			Data:       response,
		}

		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("GetCode", &request).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/getcode", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when email is present", func(t *testing.T) {
		request := dto.CodeReq{Email: ""}
		response := httperror.AppError{
			StatusCode: 400,
			Code:       "BAD_REQUEST",
			Message:    "Key: 'AuthReq.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'AuthReq.Password' Error:Field validation for 'Password' failed on the 'required' tag",
		}
		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("SignIn", &request).Return(&response, nil)

		res, _ := json.Marshal(&response)
		req, _ := http.NewRequest(http.MethodPost, "/signin", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := dto.CodeReq{Email: "test@test.com"}
		response := httperror.AppError{
			Message: "User Account Inactive",
		}
		m := map[string]string{}
		m["error"] = "User Account Inactive"

		responseError := m

		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("GetCode", &request).Return(nil, response)

		res, _ := json.Marshal(&responseError)
		req, _ := http.NewRequest(http.MethodPost, "/getcode", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
}

func TestHandler_ChangePassword(t *testing.T) {
	t.Run("should return success when request exist", func(t *testing.T) {
		request := dto.ChangePReq{
			Email:       "test@test.com",
			NewPassword: "test",
			Code:        12345,
		}
		response := dto.ChangePRes{Success: "success"}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}

		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("ChangePassword", &request).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/changepassword", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when request is empty", func(t *testing.T) {
		request := dto.ChangePReq{
			Email:       "",
			NewPassword: "",
		}
		response := error(httperror.BadRequestError("Code Invalid", "401"))
		responseError := ("{\"error\":\"Code Invalid\"}")
		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("ChangePassword", &request).Return(nil, response)

		req, _ := http.NewRequest(http.MethodPatch, "/changepassword", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())

	})

	t.Run("should return error when request is empty", func(t *testing.T) {
		request := dto.ChangePReq{
			NewPassword: "",
			Code:        12345,
		}
		response := error(httperror.BadRequestError("Email is not found", "400"))
		responseError := ("{\"error\":\"Email is not found\"}")
		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("ChangePassword", &request).Return(nil, response)

		req, _ := http.NewRequest(http.MethodPatch, "/changepassword", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}
