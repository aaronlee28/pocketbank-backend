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

func TestHandler_AdminUsersList(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.UsersListRes{
			Id:           0,
			Name:         "test",
			ReferralCode: 0,
			IsActive:     false,
		}
		var responseArray []dto.UsersListRes
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("AdminUsersList").Return(&responseArray, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/userslist", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("AdminUsersList").Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/userslist", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}
