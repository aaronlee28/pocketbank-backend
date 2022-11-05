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

func TestHandler_TopupSavings(t *testing.T) {
	t.Run("should return successful when payload given", func(t *testing.T) {
		request := dto.TopupSavingsReq{
			Amount:             100000,
			SenderWalletNumber: 1,
			Description:        "",
		}
		response := dto.TopupSavingsRes{
			Amount:      100000,
			Description: "",
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 201,
			Message:    "Created",
			Data:       response,
		}

		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("TopupSavings", &request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)

		req, _ := http.NewRequest(http.MethodPost, "/topupsavings", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when amount is less than 50000", func(t *testing.T) {
		request := dto.TopupSavingsReq{
			Amount:             49999,
			SenderWalletNumber: 1,
			Description:        "",
		}
		response := httperror.AppError{
			Message: "Minimum Amount is Rp.50000",
		}
		responseError := ("{\"error\":\"Minimum Amount is Rp.50000\"}")
		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("TopupSavings", &request, 0).Return(nil, response)

		req, _ := http.NewRequest(http.MethodPost, "/topupsavings", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_Payment(t *testing.T) {
	t.Run("should return successful when payload given", func(t *testing.T) {
		request := dto.PaymentReq{
			ReceiverAccount: 1,
			Amount:          100000,
			Description:     "",
		}
		response := dto.PaymentRes{
			SenderAccount:   2,
			ReceiverAccount: 1,
			SenderName:      "test",
			ReceiverName:    "test",
			Amount:          100000,
			Status:          "Success",
			Description:     "",
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 201,
			Message:    "Created",
			Data:       response,
		}

		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("Payment", &request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)

		_, _ = json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/payment", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when amount is less than 1000", func(t *testing.T) {
		request := dto.PaymentReq{
			ReceiverAccount: 12345,
			Amount:          999,
			Description:     "",
		}
		response := httperror.AppError{
			Message: "Minimum Amount is Rp.1000",
		}
		responseError := ("{\"error\":\"Minimum Amount is Rp.1000\"}")
		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("Payment", &request, 0).Return(nil, response)

		req, _ := http.NewRequest(http.MethodPost, "/payment", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})

}

func TestHandler_TopupDeposit(t *testing.T) {
	t.Run("should return successful when payload given", func(t *testing.T) {
		request := dto.TopupDepositReq{
			Amount:      1000000,
			Duration:    1,
			AutoDeposit: false,
		}
		response := dto.DepositRes{
			Amount:      1000000,
			Duration:    1,
			AutoDeposit: false,
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 201,
			Message:    "Created",
			Data:       response,
		}

		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("TopupDeposit", &request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)

		_, _ = json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/topupdeposit", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when amount is less than 1000000", func(t *testing.T) {
		request := dto.TopupDepositReq{
			Amount:      1000000,
			Duration:    1,
			AutoDeposit: false,
		}
		response := httperror.AppError{
			Message: "Minimum Amount is Rp.1000000",
		}
		responseError := ("{\"error\":\"Minimum Amount is Rp.1000000\"}")
		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("TopupDeposit", &request, 0).Return(nil, response)

		req, _ := http.NewRequest(http.MethodPost, "/topupdeposit", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})

}

func TestHandler_TopUpQr(t *testing.T) {
	t.Run("should return response when proper payload is given", func(t *testing.T) {
		request := dto.TopUpQr{
			SenderWalletNumber: 12345,
			Amount:             100000,
			Description:        "",
		}
		response := dto.TopUpQr{
			SenderWalletNumber: 12345,
			Amount:             100000,
			Description:        "",
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 201,
			Message:    "Created",
			Data:       response,
		}

		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("TopUpQr", &request, 0).Return(&response, nil)
		res, _ := json.Marshal(&responseSuccess)

		_, _ = json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/topupqr/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when user is not found", func(t *testing.T) {
		request := dto.TopUpQr{
			Amount:      100000,
			Description: "",
		}
		response := httperror.AppError{
			Message: "User Not Found",
		}
		responseError := ("{\"error\":\"User Not Found\"}")
		mockService := new(mocks.TransactionService)
		router := &server.RouterConfig{TransactionService: mockService}
		mockService.On("TopUpQr", &request, 0).Return(nil, response)

		req, _ := http.NewRequest(http.MethodPost, "/topupqr/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}
