package handlers_test

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/server"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/testutils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
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

func TestHandler_AdminUserTransaction(t *testing.T) {
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
		response := dto.TransRes{
			SenderWalletNumber:   0,
			ReceiverWalletNumber: 0,
			SenderName:           "test",
			ReceiverName:         "test",
			Amount:               0,
			Type:                 "test",
			Status:               "test",
			Description:          "test",
			CreatedAt:            time.Time{},
		}
		var responseArray []dto.TransRes
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("AdminUserTransaction", request, 0).Return(&responseArray, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/usertransaction/:id", testutils.MakeRequestBody(request))
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
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("AdminUserTransaction", request, 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/usertransaction/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_AdminUserDetails(t *testing.T) {
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
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("AdminUserDetails", 0).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/userdetails/:id", testutils.MakeRequestBody(nil))
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
		mockService.On("AdminUserDetails", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/userdetails/:id", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_AdminUserReferralDetails(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.UserReferralDetailsRes{
			Count:       0,
			ListOfUsers: nil,
		}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("AdminUserReferralDetails", 0).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/userreferraldetails/:id", testutils.MakeRequestBody(nil))
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
		mockService.On("AdminUserReferralDetails", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/userreferraldetails/:id", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_ChangeUserStatus(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("ChangeUserStatus", 0).Return(nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/changeuserstatus/:id", testutils.MakeRequestBody(nil))
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
		mockService.On("ChangeUserStatus", 0).Return(response)
		req, _ := http.NewRequest(http.MethodPatch, "/changeuserstatus/:id", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_UserDepositInfo(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := dto.UserDepositInfo{
			UserID:      0,
			Balance:     0,
			AllDeposits: nil,
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UserDepositInfo", 0).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/userdepositinfo/:id", testutils.MakeRequestBody(nil))
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
		mockService.On("UserDepositInfo", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/userdepositinfo/:id", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_UserRate(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.ChangeInterestRateReq{InterestRate: 0.2}

		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UserRate", 0, request).Return(nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/userrate/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.ChangeInterestRateReq{InterestRate: 0.2}

		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UserRate", 0, request).Return(response)
		req, _ := http.NewRequest(http.MethodPatch, "/userrate/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_Merchandise(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := models.Merchandise{
			Id:             0,
			UserID:         0,
			TotalTransfer:  0,
			Pen:            false,
			Umbrella:       false,
			CardHolder:     false,
			SendPen:        "test",
			SendUmbrella:   "test",
			SendCardHolder: "test",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("Merchandise", 0).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/merchandise/:id", testutils.MakeRequestBody(nil))
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
		mockService.On("Merchandise", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/merchandise/:id", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_UsersRate(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.ChangeInterestRateReq{InterestRate: 0.2}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UsersRate", request).Return(nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/usersrate", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.ChangeInterestRateReq{InterestRate: 0.2}

		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UsersRate", request).Return(response)
		req, _ := http.NewRequest(http.MethodPatch, "/usersrate", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_CreatePromotion(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.PromotionReq{
			Title: "",
			Photo: "",
		}
		response := &dto.PromotionReq{
			Title: "",
			Photo: "",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       request,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("CreatePromotion", request).Return(response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/promotion", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.PromotionReq{
			Title: "",
			Photo: "",
		}
		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("CreatePromotion", request).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPost, "/promotion", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_GetPromotion(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := models.Promotion{
			Id:    0,
			Title: "test",
			Photo: "test",
		}
		var responseArray []models.Promotion
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("GetPromotion").Return(&responseArray, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/promotion", testutils.MakeRequestBody(nil))
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
		mockService.On("GetPromotion").Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/promotion", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_UpdatePromotion(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.PatchPromotionReq{
			Title: "",
			Photo: "",
		}
		response := &dto.PatchPromotionReq{
			Title: "",
			Photo: "",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UpdatePromotion", 0, request).Return(response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/updatepromotion/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})
	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.PatchPromotionReq{
			Title: "",
			Photo: "",
		}
		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UpdatePromotion", 0, request).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPatch, "/updatepromotion/:id", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_DeletePromotion(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := models.Promotion{
			Id:    0,
			Title: "test",
			Photo: "test",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("DeletePromotion", 0).Return(&response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/promotion/:id", testutils.MakeRequestBody(nil))
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
		mockService.On("DeletePromotion", 0).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPatch, "/promotion/:id", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_EligibleMerchandiseList(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := models.Merchandise{
			Id:             0,
			UserID:         0,
			TotalTransfer:  0,
			Pen:            false,
			Umbrella:       false,
			CardHolder:     false,
			SendPen:        "test",
			SendUmbrella:   "test",
			SendCardHolder: "test",
		}
		var responseArray []models.Merchandise
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("EligibleMerchandiseList").Return(&responseArray, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/eligiblemerchandiselist", testutils.MakeRequestBody(nil))
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
		mockService.On("EligibleMerchandiseList").Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/eligiblemerchandiselist", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_MerchandiseStatus(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.MerchandiseStatus{
			UserID:      0,
			MerchToSend: "test",
			Status:      "test",
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("MerchandiseStatus", request).Return(nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPost, "/merchandisestatus", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.MerchandiseStatus{
			UserID:      0,
			MerchToSend: "test",
			Status:      "test",
		}
		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("MerchandiseStatus", request).Return(response)
		req, _ := http.NewRequest(http.MethodPost, "/merchandisestatus", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_UpdateMerchStocks(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		request := &dto.UpdateMerchStocksReq{
			Name:  "test",
			Stock: 1,
		}
		response := &models.Merchstock{
			Id:         0,
			Name:       "test",
			StockCount: 1,
		}
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       response,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UpdateMerchStocks", request).Return(response, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodPatch, "/updatemerchstocks", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(res), rec.Body.String())
	})

	t.Run("should return error when error is not nil", func(t *testing.T) {
		request := &dto.UpdateMerchStocksReq{
			Name:  "test",
			Stock: 1,
		}
		response := httperror.AppError{
			Message: "Test Error",
		}
		responseError := ("{\"error\":\"Test Error\"}")
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("UpdateMerchStocks", request).Return(nil, response)
		req, _ := http.NewRequest(http.MethodPatch, "/updatemerchstocks", testutils.MakeRequestBody(request))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}

func TestHandler_GetMerchStock(t *testing.T) {
	t.Run("should return successful when payload is given", func(t *testing.T) {
		response := models.Merchstock{
			Id:         1,
			Name:       "test",
			StockCount: 1,
		}
		var responseArray []models.Merchstock
		responseArray = append(responseArray, response)
		responseSuccess := httpsuccess.AppSuccess{
			StatusCode: 200,
			Message:    "Ok",
			Data:       responseArray,
		}
		mockService := new(mocks.AdminService)
		router := &server.RouterConfig{AdminService: mockService}
		mockService.On("GetMerchStock").Return(&responseArray, nil)

		res, _ := json.Marshal(&responseSuccess)
		req, _ := http.NewRequest(http.MethodGet, "/getmerchstock", testutils.MakeRequestBody(nil))
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
		mockService.On("GetMerchStock").Return(nil, response)
		req, _ := http.NewRequest(http.MethodGet, "/getmerchstock", testutils.MakeRequestBody(nil))
		_, rec := testutils.ServeReq(router, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, responseError, rec.Body.String())
	})
}
