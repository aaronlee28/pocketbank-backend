package handlers_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
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
		mockService := new(mocks.AuthService)
		router := &server.RouterConfig{AuthService: mockService}
		mockService.On("SignIn", &request).Return(&response, nil)

		_, _ = json.Marshal(&response)
		req, _ := http.NewRequest(http.MethodPost, "/signin", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Contains(t, string(res), rec.Body.String())

	})
}
