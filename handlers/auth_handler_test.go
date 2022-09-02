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

//func TestHandler_Register(t *testing.T) {
//	t.Run("should return user when user is registering", func(t *testing.T) {
//		request := dto.RegReq{
//			Name:           "test",
//			Email:          "test@test.com",
//			Contact:        "0123456789",
//			Password:       "test",
//			ReferralNumber: 12345,
//			Photo:          []uint8(nil),
//		}
//		response := dto.RegRes{
//			Name:    "test",
//			Email:   "test@test.com",
//			Contact: "0123456789",
//		}
//		mockService := new(mocks.AuthService)
//		router := &server.RouterConfig{AuthService: mockService}
//		mockService.On("Register", &request).Return(&response, nil)
//
//		res, _ := json.Marshal(&response)
//		req, _ := http.NewRequest(http.MethodPost, "/register", testutils.MakeRequestBody(request))
//		req.Form = request
//		_, rec := testutils.ServeReq(router, req)
//
//		//assert.Equal(t, http.StatusOK, rec.Code)
//		assert.Equal(t, string(res), rec.Body.String())
//
//	})
//}

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

	t.Run("should return error when signin request is invalid", func(t *testing.T) {
		request := dto.AuthReq{
			Email:    "",
			Password: "",
		}
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

		//assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
}
