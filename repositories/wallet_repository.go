package repositories

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
	"strconv"
)

type WalletRepository interface {
	Transaction(q *Query, id int) (*[]models.Transaction, error)
	UserDetails(id int) (*dto.UserDetailsRes, error)
	DepositInfo(id int) (*[]models.Deposit, error)
}

type walletRepository struct {
	db *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

type Query struct {
	SortBy       string
	Sort         string
	Limit        string
	Page         string
	Search       string
	FilterTime   string
	FilterAmount string
}

func NewWalletRepository(c *WRConfig) walletRepository {
	return walletRepository{db: c.DB}
}

func (w *walletRepository) Transaction(q *Query, id int) (*[]models.Transaction, error) {
	var trans *[]models.Transaction
	var account *models.Savings
	limit, _ := strconv.Atoi(q.Limit)
	page, _ := strconv.Atoi(q.Page)
	search := "%" + q.Search + "%"
	offset := (limit * page) - limit
	w.db.Where("user_id = ?", id).First(&account)
	err := w.db.Limit(limit).Offset(offset).Order(q.SortBy+" "+q.Sort).Where("sender_wallet_number = ? OR receiver_wallet_number = ? ", account.SavingsNumber, account.SavingsNumber).Where("UPPER(description) like UPPER(?)", search).Find(&trans).Error

	return trans, err
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

func (w *walletRepository) DepositInfo(id int) (*[]models.Deposit, error) {
	//var user *models.User
	var ds *[]models.Deposit
	err := w.db.Where("user_id = ? ", id).Find(&ds).Error

	return ds, err
}
