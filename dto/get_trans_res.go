package dto

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"time"
)

type TransRes struct {
	SenderWalletNumber   int       `json:"from"`
	ReceiverWalletNumber int       `json:"to"`
	Amount               float32   `json:"amount"`
	Description          string    `json:"description"`
	SourceOfFund         string    `json:"source_of_fund"`
	CreatedAt            time.Time `json:"date"`
}

func (_ *TransRes) FromTransaction(t *models.Transaction) *TransRes {
	return &TransRes{
		SenderWalletNumber:   t.SenderWalletNumber,
		ReceiverWalletNumber: t.ReceiverWalletNumber,
		Amount:               t.Amount,
		Description:          t.Description,
		CreatedAt:            t.CreatedAt,
	}
}
