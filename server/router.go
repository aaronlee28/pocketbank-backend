package server

import (
	dto2 "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	dto "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto/auth"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/handlers"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/middlewares"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthService services.AuthService
	//WalletService services.WalletService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:   c.AuthService,
		WalletService: c.WalletService,
	})
	router.Static("/docs", "swaggerui")
	router.NoRoute(middlewares.WrongEndpoint())
	router.GET("/getcode", middlewares.RequestValidator(&dto2.CodeReq{}), h.GetCode)
	router.PATCH("/changepassword", middlewares.RequestValidator(&dto2.ChangePReq{}), h.ChangePassword)
	router.POST("/register", middlewares.RequestValidator(&dto2.AuthReq{}), h.Register)
	router.POST("/signin", middlewares.RequestValidator(&dto2.AuthReq{}), h.SignIn)
	router.Use(middlewares.AuthorizeJWT)
	router.POST("/topup", middlewares.RequestValidator(&dto.TopupReq{}), h.Topup)
	router.GET("/transaction", h.Transaction)
	router.POST("/transfer", middlewares.RequestValidator(&dto.TransferReq{}), h.Transfer)
	router.GET("/userdetails", middlewares.RequestValidator(&dto.UserDetailsRes{}), h.UserDetails)

	return router
}
