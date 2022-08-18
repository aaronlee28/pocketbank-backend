package dto

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
)

type DepositInfoRes struct {
	Balance       float32 `json:"balance"`
	DepositNumber int     `json:"deposit_number"`
}

func (_ *DepositInfoRes) FromDepositInfo(t *models.Deposit) *DepositInfoRes {
	return &DepositInfoRes{
		Balance:       t.Balance,
		DepositNumber: t.DepositNumber,
	}
}
