package handlers

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"

type Handler struct {
	authService        services.AuthService
	TransactionService services.TransactionService
	WalletService      services.WalletService
}

type HandlerConfig struct {
	AuthService        services.AuthService
	TransactionService services.TransactionService
	WalletService      services.WalletService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		authService:        c.AuthService,
		TransactionService: c.TransactionService,
		WalletService:      c.WalletService,
	}
}
