package repositories

import (
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type TransactionRepository interface {
	TopupSavings(trans *models.Transaction, id int) (*models.Transaction, error, error)

	Payment(trans *models.Transaction, id int) (*models.Transaction, error)

	UpdateInterestAndTaxSavings()
	RunCronJobs()

	TopupDeposit(trans *dto.TopupDepositReq, id int) (*models.Deposit, error)
	TopUpQr(trans *dto.TopUpQr, id int) (*dto.TopUpQr, error)
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
	var m *models.Merchandise
	err2 := w.db.Where("user_id = ?", id).First(&sv)
	newBalance := sv.Balance + trans.Amount
	w.db.Model(&sv).Update("balance", newBalance)
	w.db.Where("user_id = ?", id).First(&m)
	totalTransfer := m.TotalTransfer + trans.Amount
	w.db.Model(&m).Update("total_transfer", totalTransfer)
	if totalTransfer > 1000000 {
		w.db.Model(&m).Update("pen", true)
	}
	if totalTransfer > 5000000 {
		w.db.Model(&m).Update("umbrella", true)

	}
	if totalTransfer > 10000000 {
		w.db.Model(&m).Update("card_holder", true)
	}
	addTransaction := &models.Transaction{
		SenderWalletNumber:   trans.SenderWalletNumber,
		ReceiverWalletNumber: sv.SavingsNumber,
		Amount:               trans.Amount,
		Type:                 trans.Type,
		Description:          trans.Description,
		Status:               "Success",
	}
	err1 := db.Get().Create(&addTransaction)

	return addTransaction, err1.Error, err2.Error
}

func (w *transactionRepository) Payment(trans *models.Transaction, id int) (*models.Transaction, error) {
	var senderSavings *models.Savings
	var receiverName *models.User
	var senderName *models.User
	var receiverSavings *models.Savings
	var checkBalance float32
	var addBalance float32

	w.db.Where("user_id = ?", id).First(&senderSavings)
	err := w.db.Where("savings_number= ?", trans.ReceiverWalletNumber).First(&receiverSavings).Error
	receiverId := receiverSavings.UserID
	w.db.Where("id = ?", senderSavings.UserID).First(&senderName)
	w.db.Where("id = ?", receiverId).First(&receiverName)

	addFailedPayment := &models.Transaction{
		SenderWalletNumber:   senderSavings.SavingsNumber,
		ReceiverWalletNumber: trans.ReceiverWalletNumber,
		Amount:               trans.Amount,
		Type:                 trans.Type,
		Status:               "Failed",
		Description:          "Failed Transaction",
	}
	checkBalance = senderSavings.Balance - trans.Amount
	if checkBalance < 0 {
		addFailedPayment.Description = "Insufficient Balance"
		db.Get().Create(&addFailedPayment)

		return addFailedPayment, nil
	}
	if err != nil {
		addFailedPayment.Description = "Destination Account Not Found"
		db.Get().Create(&addFailedPayment)

		return addFailedPayment, nil
	}
	addSuccessfulPayment := &models.Transaction{
		SenderWalletNumber:   senderSavings.SavingsNumber,
		ReceiverWalletNumber: trans.ReceiverWalletNumber,
		SenderName:           senderName.Name,
		ReceiverName:         receiverName.Name,
		Amount:               trans.Amount,
		Type:                 trans.Type,
		Status:               "Success",
		Description:          trans.Description,
	}
	revertBalance := receiverSavings.Balance
	addBalance = receiverSavings.Balance + trans.Amount

	err3 := w.db.Model(&senderSavings).Update("balance", checkBalance).Error
	if err3 != nil {
		db.Get().Create(&addFailedPayment)
		return addFailedPayment, err3
	}
	err4 := w.db.Model(&receiverSavings).Update("balance", addBalance).Error
	if err4 != nil {
		db.Get().Create(&addFailedPayment)
		w.db.Model(&senderSavings).Update("balance", revertBalance)
		return addFailedPayment, err4
	}

	db.Get().Create(&addSuccessfulPayment)
	return addSuccessfulPayment, nil
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
				Type:                 "Interest",
				Description:          "Interest",
			}
			db.Get().Create(&addInterestTransaction)

			addTaxTransaction := &models.Transaction{
				SenderWalletNumber:   s.SavingsNumber,
				ReceiverWalletNumber: 2,
				Amount:               addInterest,
				Type:                 "Tax on Interest",
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
			isOneMonth := int(difference.Hours() / 24 / 30)

			if isOneMonth == s.Duration {
				w.db.Where("user_id", s.UserID).First(&sv)
				addInterest := sv.Balance + s.Interest
				w.db.Model(&sv).Update("balance", addInterest)
				w.db.Model(&s).Update("updated_at", time.Now().UTC())
				addTransaction1 := &models.Transaction{
					SenderWalletNumber:   3,
					ReceiverWalletNumber: sv.SavingsNumber,
					Amount:               addInterest,
					Type:                 "Deposit Interest",
					Description:          "Deposit Interest",
				}
				db.Get().Create(&addTransaction1)

			}
		}
		if s.AutoDeposit == false {
			difference := time.Now().UTC().Sub(s.UpdatedAt)
			isOneMonth := int64(difference.Hours() / 24 / 30)
			if isOneMonth == 1 && time.Time.IsZero(s.DeletedAt) == true {
				w.db.Where("user_id", s.UserID).First(&sv)
				addInterest := sv.Balance + s.Interest + s.Balance
				w.db.Model(&sv).Update("balance", addInterest)
				w.db.Model(&s).Update("deleted_at", time.Now())
				addTransaction2 := &models.Transaction{
					SenderWalletNumber:   4,
					ReceiverWalletNumber: sv.SavingsNumber,
					Amount:               addInterest,
					Type:                 "Deposit Withdrawal",
					Description:          "Deposit Withdrawal",
				}
				db.Get().Create(&addTransaction2)
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

func (w *transactionRepository) TopupDeposit(trans *dto.TopupDepositReq, id int) (*models.Deposit, error) {
	var sv *models.Savings
	err1 := new(error)
	w.db.Where("user_id = ?", id).First(&sv)
	if sv.Balance-trans.Amount < 0 {
		return nil, *err1
	}
	addDeposit := &models.Deposit{
		UserID:        id,
		Balance:       trans.Amount,
		Tax:           0.2,
		DepositNumber: 3 + rand.Intn(99999-10000) + 10000 + id,
		Duration:      trans.Duration,
		AutoDeposit:   trans.AutoDeposit,
	}
	if trans.Amount < 10000000 {
		addDeposit.InterestRate = 0.06
	} else {
		addDeposit.InterestRate = 0.08
	}
	interestBeforeTax := (trans.Amount * addDeposit.InterestRate) / 12
	interestAfterTax := interestBeforeTax * (1 - addDeposit.Tax)
	addDeposit.Interest = interestAfterTax
	db.Get().Create(&addDeposit)
	w.db.Model(&addDeposit).Update("deleted_at", nil)

	newBalance := sv.Balance - trans.Amount
	w.db.Model(&sv).Update("balance", newBalance)

	addTransaction := &models.Transaction{
		SenderWalletNumber:   sv.SavingsNumber,
		ReceiverWalletNumber: addDeposit.DepositNumber,
		Amount:               trans.Amount,
		Type:                 "Deposit",
		Status:               "Success",
		Description:          "Deposit",
	}
	db.Get().Create(&addTransaction)

	return addDeposit, nil
}

func (w *transactionRepository) TopUpQr(trans *dto.TopUpQr, id int) (*dto.TopUpQr, error) {
	var sv *models.Savings
	err := w.db.Where("savings_number = ?", id).First(&sv).Error
	if err != nil {
		return nil, err
	}
	newBalance := sv.Balance + trans.Amount
	w.db.Model(&sv).Update("balance", newBalance)

	addTransaction := &models.Transaction{
		SenderWalletNumber:   trans.SenderWalletNumber,
		ReceiverWalletNumber: sv.SavingsNumber,
		Amount:               trans.Amount,
		Type:                 "External Top Up",
		Description:          trans.Description,
		Status:               "Success",
	}
	db.Get().Create(&addTransaction)

	return trans, err
}
