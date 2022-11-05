package handlers

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"

type Handler struct {
	AuthService        services.AuthService
	TransactionService services.TransactionService
	WalletService      services.WalletService
	AdminService       services.AdminService
}

type HandlerConfig struct {
	AuthService        services.AuthService
	TransactionService services.TransactionService
	WalletService      services.WalletService
	AdminService       services.AdminService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		AuthService:        c.AuthService,
		TransactionService: c.TransactionService,
		WalletService:      c.WalletService,
		AdminService:       c.AdminService,
	}
}
