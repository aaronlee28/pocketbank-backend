package dto

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"time"
)

type PaymentHistoryRes struct {
	SenderWalletNumber   int       `json:"from"`
	ReceiverWalletNumber int       `json:"to"`
	Amount               float32   `json:"amount"`
	Description          string    `json:"description"`
	CreatedAt            time.Time `json:"date"`
}

func (_ *PaymentHistoryRes) FromTransaction(t *models.Transaction) *PaymentHistoryRes {
	return &PaymentHistoryRes{
		SenderWalletNumber:   t.SenderWalletNumber,
		ReceiverWalletNumber: t.ReceiverWalletNumber,
		Amount:               t.Amount,
		Description:          t.Description,
		CreatedAt:            t.CreatedAt,
	}
}
