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
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:        c.AuthService,
		TransactionService: c.TransactionService,
		WalletService:      c.WalletService,
	})
	router.Static("/docs", "swaggerui")
	router.NoRoute(middlewares.WrongEndpoint())
	router.Use(h.RunCronJobs)
	router.GET("/getcode", middlewares.RequestValidator(&dto.CodeReq{}), h.GetCode)
	router.PATCH("/changepassword", middlewares.RequestValidator(&dto.ChangePReq{}), h.ChangePassword)
	router.POST("/register", middlewares.RequestValidator(&dto.RegReq{}), h.Register)
	router.POST("/signin", middlewares.RequestValidator(&dto.AuthReq{}), h.SignIn)
	router.Use(middlewares.AuthorizeJWT)
	router.POST("/topupsavings", middlewares.RequestValidator(&dto.TopupSavingsReq{}), h.TopupSavings)
	router.GET("/transactionhistory", h.TransactionHistory)
	router.POST("/payment", middlewares.RequestValidator(&dto.PaymentReq{}), h.Payment)
	router.GET("/userdetails", h.UserDetails)
	router.POST("/topupdeposit", middlewares.RequestValidator(&dto.TopupDepositReq{}), h.TopupDeposit)
	router.GET("/depositinfo", h.DepositInfo)
	router.GET("/paymenthistory", h.PaymentHistory)
	router.POST("/favoritecontact", middlewares.RequestValidator(&dto.FavoriteContactReq{}), h.FavoriteContact)
	return router
}
