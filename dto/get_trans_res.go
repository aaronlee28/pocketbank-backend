package dto

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"time"
)

type TransRes struct {
	SenderWalletNumber   int       `json:"from"`
	ReceiverWalletNumber int       `json:"to"`
	SenderName           string    `json:"senderName"`
	ReceiverName         string    `json:"receiverName"`
	Amount               float32   `json:"amount"`
	Type                 string    `json:"type"`
	Status               string    `json:"status"`
	Description          string    `json:"description"`
	CreatedAt            time.Time `json:"date"`
}

func (_ *TransRes) FromTransaction(t *models.Transaction) *TransRes {
	return &TransRes{
		SenderWalletNumber:   t.SenderWalletNumber,
		ReceiverWalletNumber: t.ReceiverWalletNumber,
		SenderName:           t.SenderName,
		ReceiverName:         t.ReceiverName,
		Amount:               t.Amount,
		Type:                 t.Type,
		Status:               t.Status,
		Description:          t.Description,
		CreatedAt:            t.CreatedAt,
	}
}
