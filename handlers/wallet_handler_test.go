package handlers_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/server"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/testutils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestHandler_TransactionHistory(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &repositories.Query{
			SortBy:     "",
			Sort:       "",
			Limit:      "",
			Page:       "",
			Search:     "",
			FilterTime: "",
			MinAmount:  "",
			MaxAmount:  "",
			Type:       "",
		}
		response := dto.TransHistoryRes{
			TotalLength:  0,
			Transactions: nil,
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("TransactionHistory", request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/transactionhistory", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &repositories.Query{
			SortBy:     "",
			Sort:       "",
			Limit:      "",
			Page:       "",
			Search:     "",
			FilterTime: "",
			MinAmount:  "",
			MaxAmount:  "",
			Type:       "",
		}
		response := httperror.AppError{
			Message: "Bad Request",
		}

		responseError := ("{\"error\":\"Bad Request\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("TransactionHistory", request, 0).Return(nil, &response)
		req, _ := http.NewRequest(http.MethodGet, "/transactionhistory", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_UserDetails(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.UserDetailsRes{
			Name:           "test",
			Email:          "test",
			Contact:        "test",
			ProfilePicture: "test",
			ReferralNumber: 0,
			SavingsNumber:  0,
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("UserDetails", 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/userdetails", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		response := httperror.AppError{
			Message: "Test Error",
		}

		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("UserDetails", 0).Return(nil, &response)
		req, _ := http.NewRequest(http.MethodGet, "/userdetails", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_DepositInfo(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.DepositInfoRes{
			Balance:       0,
			DepositNumber: 0,
			InterestRate:  0,
			Interest:      0,
			Duration:      0,
			UpdatedAt:     time.Time{},
			DeletedAt:     time.Time{},
			AutoDeposit:   false,
		}
		var responseArray []dto.DepositInfoRes
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("DepositInfo", 0).Return(&responseArray, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/depositinfo", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		response := httperror.AppError{
			Message: "Test Error",
		}

		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("DepositInfo", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/depositinfo", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_SavingsInfo(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.SavingsRes{
			UserID:   0,
			Balance:  0,
			Interest: 0,
			Tax:      0,
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("SavingsInfo", 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/savingsinfo", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
	t.Run("should return error when error is not nil", func(t *testing.T) {
		response := httperror.AppError{
			Message: "Test Error",
		}

		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("SavingsInfo", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/savingsinfo", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_FavoriteContact(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := dto.FavoriteContactReq{FavoriteAccountNumber: 0}

		response := dto.FavoriteContactRes{
			UserID:                0,
			FavoriteAccountNumber: 0,
			Favorite:              false,
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("FavoriteContact", &request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/favoritecontact", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := dto.FavoriteContactReq{FavoriteAccountNumber: 0}
		response := httperror.AppError{
			Message: "Test Error",
		}

		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("FavoriteContact", &request, 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPost, "/favoritecontact", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})

}

func TestHandler_FavoriteContactList(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.FavoriteContactRes{
			UserID:                0,
			FavoriteAccountNumber: 0,
			Favorite:              false,
		}
		var responseArray []dto.FavoriteContactRes
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("FavoriteContactList", 0).Return(&responseArray, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/favoritecontactlist", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		response := httperror.AppError{
			Message: "Test Error",
		}

		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("FavoriteContactList", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/favoritecontactlist", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_ChangeUserDetails(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.ChangeUserDetailsReqRes{
			Name:           "",
			Email:          "",
			Contact:        "",
			ProfilePicture: "",
			ReferralNumber: 0,
			AccountNumber:  0,
		}
		response := dto.ChangeUserDetailsReqRes{
			Name:           "0",
			Email:          "0",
			Contact:        "0",
			ProfilePicture: "0",
			ReferralNumber: 0,
			AccountNumber:  0,
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("ChangeUserDetails", request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/userdetails", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.ChangeUserDetailsReqRes{
			Name:           "",
			Email:          "",
			Contact:        "",
			ProfilePicture: "",
			ReferralNumber: 0,
			AccountNumber:  0,
		}
		response := httperror.AppError{
			Message: "Test Error",
		}

		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.WalletService)
		router := &server.RouterConfig{WalletService: mockService}
		mockService.On("ChangeUserDetails", request, 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPatch, "/userdetails", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}
