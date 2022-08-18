package server

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/services"
)

func Init() {
	authRepository := repositories.NewAuthRepository(&repositories.ARConfig{DB: db.Get()})
	authService := services.NewAuthService(&services.ASConfig{
		AuthRepository: &authRepository,
		AppConfig:      config.Config,
	})

	walletRepository := repositories.NewWalletRepository(&repositories.WRConfig{DB: db.Get()})
	walletService := services.NewWalletServices(&services.WSConfig{WalletRepository: &walletRepository})

	transactionRepository := repositories.NewTransactionRepository(&repositories.TRConfig{DB: db.Get()})
	transactionService := services.NewTransactionServices(&services.TSConfig{TransactionRepository: &transactionRepository})

	router := NewRouter(&RouterConfig{
		AuthService:        authService,
		WalletService:      walletService,
		TransactionService: transactionService,
	})
	err := router.Run()
	if err != nil {
		fmt.Println("server error:", err)
	}
}
