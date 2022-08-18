package repositories

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type TransactionRepository interface {
	TopupSavings(trans *models.Transaction, id int) (*models.Transaction, error, error)

	Transfer(trans *models.Transaction, id int) (*models.Transaction, error, error, error)

	UpdateInterestAndTaxSavings()
	RunCronJobs()

	TopupDeposit(trans *models.Transaction, id int) (*models.Transaction, error, error)
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
func (w *transactionRepository) TopupSavings(trans *models.Transaction, id int) (*models.Transaction, error, error) {
	var sv *models.Savings
	err2 := w.db.Where("user_id = ?", id).First(&sv)
	newBalance := sv.Balance + trans.Amount
	w.db.Model(&sv).Update("balance", newBalance)
	addTransaction := &models.Transaction{
		SenderWalletNumber:   trans.SenderWalletNumber,
		ReceiverWalletNumber: sv.SavingsNumber,
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

func (w *transactionRepository) UpdateInterestAndTaxSavings() {
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

func (w *transactionRepository) WithdrawDeposit() {
	var ds *[]models.Deposit
	var sv *models.Savings
	w.db.Find(&ds)

	for _, s := range *ds {
		if s.AutoDeposit == true {

			difference := time.Now().UTC().Sub(s.UpdatedAt)
			isOneMonth := int64(difference.Hours() / 24 / 30)

			if isOneMonth == 1 {
				w.db.Where("user_id", s.UserID).First(&sv)
				addInterest := sv.Balance + s.Interest
				w.db.Model(&sv).Update("balance", addInterest)
				w.db.Model(&s).Update("updated_at", time.Now().UTC())
			}
		}
		if s.AutoDeposit == false {
			difference := time.Now().UTC().Sub(s.UpdatedAt)
			isOneMonth := int64(difference.Hours() / 24 / 30)
			if isOneMonth == 1 {
				w.db.Where("user_id", s.UserID).First(&sv)
				addInterest := sv.Balance + s.Interest + s.Balance
				w.db.Model(&sv).Update("balance", addInterest)
				w.db.Delete(&s)
				fmt.Println("im here")

			}
		}
	}
}

func (w *transactionRepository) RunCronJobs() {
	c := cron.New(cron.WithLocation(time.UTC))
	_, _ = c.AddFunc("@daily", func() { w.UpdateInterestAndTaxSavings() })
	_, _ = c.AddFunc("@daily", func() { w.WithdrawDeposit() })
	c.Start()

}

func (w *transactionRepository) TopupDeposit(trans *models.Transaction, id int) (*models.Transaction, error, error) {
	var sv *models.Savings
	w.db.Where("user_id = ?", id).First(&sv)

	addDeposit := &models.Deposit{
		UserID:        id,
		Balance:       trans.Amount,
		Tax:           0.2,
		DepositNumber: 3 + rand.Intn(99999-10000) + 10000 + id,
	}
	if trans.Amount < 10000000 {
		addDeposit.InterestRate = 0.06
	} else {
		addDeposit.InterestRate = 0.08
	}
	interestBeforeTax := (trans.Amount * addDeposit.InterestRate) / 12
	interestAfterTax := interestBeforeTax * (1 - addDeposit.Tax)
	addDeposit.Interest = interestAfterTax
	err2 := db.Get().Create(&addDeposit)

	newBalance := sv.Balance - trans.Amount
	w.db.Model(&sv).Update("balance", newBalance)

	addTransaction := &models.Transaction{
		SenderWalletNumber:   sv.SavingsNumber,
		ReceiverWalletNumber: addDeposit.DepositNumber,
		Amount:               trans.Amount,
		Description:          "Deposit",
	}
	err1 := db.Get().Create(&addTransaction)

	return addTransaction, err1.Error, err2.Error
}
