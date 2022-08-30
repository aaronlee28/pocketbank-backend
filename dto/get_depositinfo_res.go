package dto

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"time"
)

type DepositInfoRes struct {
	Balance       float32   `json:"balance"`
	DepositNumber int       `json:"depositNumber"`
	InterestRate  float32   `json:"interestRate"`
	Interest      float32   `json:"interest"`
	Duration      int       `json:"duration"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
	AutoDeposit   bool      `json:"autoDeposit"`
}

func (_ *DepositInfoRes) FromDepositInfo(t *models.Deposit) *DepositInfoRes {
	return &DepositInfoRes{
		Balance:       t.Balance,
		DepositNumber: t.DepositNumber,
		InterestRate:  t.InterestRate,
		Interest:      t.Interest,
		Duration:      t.Duration,
		UpdatedAt:     t.UpdatedAt,
		DeletedAt:     t.DeletedAt,
		AutoDeposit:   t.AutoDeposit,
	}
}
