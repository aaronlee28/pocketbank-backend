package repositories

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"strconv"
)

type WalletRepository interface {
	Topup(trans *models.Transaction, id int) (*models.Transaction, error, error)
	Transaction(q *Query, id int) (*[]models.Transaction, error)
	Transfer(trans *models.Transaction, id int) (*models.Transaction, error, error, error)
	UserDetails(id int) (*dto.UserDetailsRes, error)
	UpdateInterestAndTax()
	RunCronJobs()
}

type walletRepository struct {
	db *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

type Query struct {
	SortBy string
	Sort   string
	Limit  string
	Page   string
	Search string
}

func NewWalletRepository(c *WRConfig) walletRepository {
	return walletRepository{db: c.DB}
}

func (w *walletRepository) Topup(trans *models.Transaction, id int) (*models.Transaction, error, error) {
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

func (w *walletRepository) Transaction(q *Query, id int) (*[]models.Transaction, error) {
	var trans *[]models.Transaction
	var wallet *models.Wallet
	limit, _ := strconv.Atoi(q.Limit)
	page, _ := strconv.Atoi(q.Page)
	search := "%" + q.Search + "%"
	offset := (limit * page) - limit
	w.db.Where("user_id = ?", id).First(&wallet)
	err := w.db.Limit(limit).Offset(offset).Order(q.SortBy+" "+q.Sort).Where("sender_wallet_number = ? OR receiver_wallet_number = ? ", wallet.WalletNumber, wallet.WalletNumber).Where("UPPER(description) like UPPER(?)", search).Find(&trans).Error

	return trans, err
}

func (w *walletRepository) Transfer(trans *models.Transaction, id int) (*models.Transaction, error, error, error) {
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

func (w *walletRepository) UserDetails(id int) (*dto.UserDetailsRes, error) {
	var user *models.User
	var wallet *models.Wallet
	err := w.db.Where("id = ?", id).First(&user).Error
	w.db.Where("user_id = ?", id).First(&wallet)
	ret := &dto.UserDetailsRes{
		Id:           user.Id,
		Email:        user.Email,
		WalletID:     wallet.Id,
		WalletNumber: wallet.WalletNumber,
		Balance:      wallet.Balance,
	}

	return ret, err
}

func (w *walletRepository) UpdateInterestAndTax() {
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
			fmt.Println("im here yay")
		}

	}
}

//
func (w *walletRepository) RunCronJobs() {
	c := cron.New()
	c.AddFunc("@every 2s", func() { fmt.Println("every 2s") })
	c.Start()

}
