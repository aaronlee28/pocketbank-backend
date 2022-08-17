package middlewares

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func UpdateInterestAndTax() {
	var sv *models.Savings
	db.Where("saving_number = ? ", 20536).First(&sv)
	fmt.Println("my saving number is", sv.SavingsNumber)
}
