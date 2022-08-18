package repositories

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"time"
)

type TransactionRepository interface {
	Topup(trans *models.Transaction, id int) (*models.Transaction, error, error)

	Transfer(trans *models.Transaction, id int) (*models.Transaction, error, error, error)

	UpdateInterestAndTax()
	RunCronJobs()
}

type transactionRepository struct {
	db *gorm.DB
}

type TRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(c *TRConfig) transactionRepository {
	return transactionRepository{db: c.DB}
}
func (w *transactionRepository) Topup(trans *models.Transaction, id int) (*models.Transaction, error, error) {
	var wallet *models.Wallet
	err2 := w.db.Where("user_id = ?", id).First(&wallet)
	newBalance := wallet.Balance + trans.Amount
	w.db.Model(&wallet).Update("balance", newBalance)
	addTransaction := &models.Transaction{
		SenderWalletNumber:   wallet.WalletNumber,
		ReceiverWalletNumber: wallet.WalletNumber,
		Amount:               trans.Amount,
		Description:          trans.Description,
	}
	err1 := db.Get().Create(&addTransaction)

	return addTransaction, err1.Error, err2.Error
}

func (w *transactionRepository) Transfer(trans *models.Transaction, id int) (*models.Transaction, error, error, error) {
	var senderWallet *models.Wallet
	var receiverWallet *models.Wallet
	var checkBalance float32
	var addBalance float32
	err := fmt.Errorf("")
	_ = w.db.Where("user_id = ?", id).First(&senderWallet)
	checkBalance = senderWallet.Balance - trans.Amount
	if checkBalance < 0 {
		return nil, err, nil, nil
	}
	_ = w.db.Where("wallet_number = ?", trans.ReceiverWalletNumber).First(&receiverWallet)
	addBalance = receiverWallet.Balance + trans.Amount
	if receiverWallet.UserID == 0 {
		return nil, nil, err, nil
	}

	w.db.Model(&senderWallet).Update("balance", checkBalance)
	w.db.Model(&receiverWallet).Update("balance", addBalance)
	addTransaction := &models.Transaction{
		SenderWalletNumber:   senderWallet.WalletNumber,
		ReceiverWalletNumber: receiverWallet.WalletNumber,
		Amount:               trans.Amount,
		Description:          trans.Description,
	}
	_ = db.Get().Create(&addTransaction)

	return addTransaction, nil, nil, nil
}

func (w *transactionRepository) UpdateInterestAndTax() {
	var svs *[]models.Savings
	w.db.Find(&svs)
	for _, s := range *svs {
		if s.Balance > 0 {
			interest := (s.Balance * s.Interest) / (12 * 30)
			addInterest := s.Balance + interest
			w.db.Model(&s).Update("balance", addInterest)
			taxonInterest := interest * s.Tax
			payTax := s.Balance - taxonInterest
			w.db.Model(&s).Update("balance", payTax)

			addInterestTransaction := &models.Transaction{
				SenderWalletNumber:   1,
				ReceiverWalletNumber: s.SavingsNumber,
				Amount:               addInterest,
				Description:          "Interest",
			}
			db.Get().Create(&addInterestTransaction)

			addTaxTransaction := &models.Transaction{
				SenderWalletNumber:   s.SavingsNumber,
				ReceiverWalletNumber: 2,
				Amount:               addInterest,
				Description:          "Tax on Interest",
			}
			db.Get().Create(&addTaxTransaction)

		}

	}
}

//
func (w *transactionRepository) RunCronJobs() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	c := cron.New(cron.WithLocation(loc))
	c.AddFunc("@daily", func() { w.UpdateInterestAndTax() })
	c.Start()

}
