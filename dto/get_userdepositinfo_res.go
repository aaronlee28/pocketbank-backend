package dto

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"

type UserDepositInfo struct {
	UserID      int               `json:"userId"`
	Balance     float32           `json:"balance"`
	AllDeposits *[]models.Deposit `json:"allDeposits"`
}
