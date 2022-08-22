package server

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/handlers"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/middlewares"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthService        services.AuthService
	WalletService      services.WalletService
	TransactionService services.TransactionService
	AdminService       services.AdminService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:        c.AuthService,
		TransactionService: c.TransactionService,
		WalletService:      c.WalletService,
		AdminService:       c.AdminService,
	})
	router.Static("/docs", "swaggerui")
	router.NoRoute(middlewares.WrongEndpoint())
	router.Use(h.RunCronJobs)
	router.GET("/getcode", middlewares.RequestValidator(&dto.CodeReq{}), h.GetCode)
	router.PATCH("/changepassword", middlewares.RequestValidator(&dto.ChangePReq{}), h.ChangePassword)
	router.POST("/register", h.Register)
	router.POST("/signin", middlewares.RequestValidator(&dto.AuthReq{}), h.SignIn)
	router.Use(middlewares.AuthorizeJWT)
	router.POST("/topupsavings", middlewares.RequestValidator(&dto.TopupSavingsReq{}), h.TopupSavings)
	router.GET("/transactionhistory", h.TransactionHistory)
	router.POST("/payment", middlewares.RequestValidator(&dto.PaymentReq{}), h.Payment)
	router.GET("/userdetails", h.UserDetails)
	router.PATCH("/userdetails", h.ChangeUserDetails)
	router.POST("/topupdeposit", middlewares.RequestValidator(&dto.TopupDepositReq{}), h.TopupDeposit)
	router.GET("/depositinfo", h.DepositInfo)
	router.GET("/paymenthistory", h.PaymentHistory)
	router.POST("/favoritecontact", middlewares.RequestValidator(&dto.FavoriteContactReq{}), h.FavoriteContact)
	router.GET("/favoritecontactlist", h.FavoriteContactList)
	//admin router
	router.Use(middlewares.AuthorizeAdmin)
	router.GET("/userslist", h.AdminUsersList)
	router.GET("/usertransaction/:id", h.AdminUserTransaction)
	router.GET("/userdetails/:id", h.AdminUserDetails)
	router.GET("/userreferraldetails/:id", h.AdminUserReferralDetails)
	router.PATCH("/ChangeUserDetailserstatus/:id", h.ChangeUserStatus)
	router.GET("/merchandise/:id", h.Merchandise)
	router.GET("/userdepositinfo/:id", h.UserDepositInfo)
	router.PATCH("/userrate/:id", middlewares.RequestValidator(&dto.ChangeInterestRateReq{}), h.UserRate)
	router.PATCH("/usersrate", middlewares.RequestValidator(&dto.ChangeInterestRateReq{}), h.UsersRate)

	router.POST("/promotion", h.CreatePromotion)
	router.GET("/promotion", h.GetPromotion)
	router.PATCH("/promotion", h.UpdatePromotion)
	router.DELETE("/promotion/:id", h.DeletePromotion)

	router.GET("/eligiblemerchandiselist", h.EligibleMerchandiseList)
	router.POST("/merchandisestatus", middlewares.RequestValidator(&dto.MerchandiseStatus{}), h.MerchandiseStatus)
	router.PATCH("/updatemerchstocks", middlewares.RequestValidator(&dto.UpdateMerchStocksReq{}), h.UpdateMerchStocks)
	return router
}
