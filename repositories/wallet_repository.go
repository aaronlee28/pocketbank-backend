package repositories

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type WalletRepository interface {
	TransactionHistory(q *Query, id int) (*[]models.Transaction, error)
	UserDetails(id int) (*dto.UserDetailsRes, error)
	DepositInfo(id int) (*[]models.Deposit, error)
	PaymentHistory(id int) (*[]models.Transaction, error)
	FavoriteContact(favoriteid int, selfid int) (*models.Favoritecontact, error)
	FavoriteContactList(id int) (*[]models.Favoritecontact, error)
	ChangeUserDetails(data *dto.ChangeUserDetailsReqRes, id int) (*dto.ChangeUserDetailsReqRes, error)
}

type walletRepository struct {
	db *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

type Query struct {
	SortBy     string
	Sort       string
	Limit      string
	Page       string
	Search     string
	FilterTime string
	MinAmount  string
	MaxAmount  string
}

func NewWalletRepository(c *WRConfig) walletRepository {
	return walletRepository{db: c.DB}
}

func (w *walletRepository) TransactionHistory(q *Query, id int) (*[]models.Transaction, error) {
	var trans *[]models.Transaction
	var account *models.Savings
	limit, _ := strconv.Atoi(q.Limit)
	page, _ := strconv.Atoi(q.Page)
	search := "%" + q.Search + "%"
	offset := (limit * page) - limit
	w.db.Where("user_id = ?", id).First(&account)
	err := w.db.Limit(limit).Offset(offset).Order(q.SortBy+" "+q.Sort).Where("sender_wallet_number = ? OR receiver_wallet_number = ? ", account.SavingsNumber, account.SavingsNumber).Where("UPPER(description) like UPPER(?)", search).Where("created_at >= ? at time zone 'UTC' - interval '"+q.FilterTime+"' day", time.Now()).Where("amount BETWEEN ? and ?", q.MinAmount, q.MaxAmount).Find(&trans).Error

	return trans, err
}

func (w *walletRepository) UserDetails(id int) (*dto.UserDetailsRes, error) {
	var user *models.User
	var sv *models.Savings
	err := w.db.Where("id = ?", id).First(&user).Error
	w.db.Where("user_id = ?", id).First(&sv)
	ret := &dto.UserDetailsRes{
		Name:           user.Name,
		Email:          user.Email,
		Contact:        user.Contact,
		ProfilePicture: user.ProfilePicture,
		ReferralNumber: user.ReferralNumber,
		AccountNumber:  sv.SavingsNumber,
	}

	return ret, err
}

func (w *walletRepository) DepositInfo(id int) (*[]models.Deposit, error) {
	//var user *models.User
	var ds *[]models.Deposit
	err := w.db.Where("user_id = ? ", id).Find(&ds).Error

	return ds, err
}

func (w *walletRepository) PaymentHistory(id int) (*[]models.Transaction, error) {
	var trans *[]models.Transaction
	var account *models.Savings
	w.db.Where("user_id = ?", id).First(&account)
	err := w.db.Where("TYPE = 'Transfer'").Find(&trans).Error

	return trans, err
}

func (w *walletRepository) FavoriteContact(favoriteid int, selfid int) (*models.Favoritecontact, error) {
	var user *models.User
	err := w.db.Where("id = ?", favoriteid).First(&user).Error

	if err == nil {
		addFavoriteContact := &models.Favoritecontact{
			UserID:         selfid,
			FavoriteUserID: user.Id,
		}
		db.Get().Create(&addFavoriteContact)
		return addFavoriteContact, err

	}
	return nil, err
}

func (w *walletRepository) FavoriteContactList(id int) (*[]models.Favoritecontact, error) {
	var fc *[]models.Favoritecontact
	err := w.db.Where("user_id = ?", id).Find(&fc).Error

	return fc, err
}

func (w *authRepository) ChangeUserDetails(data *dto.ChangeUserDetailsReqRes, id int) (*dto.ChangeUserDetailsReqRes, error) {
	var user *models.User
	err := w.db.Where("id = ?", id).First(&user).Error
	fmt.Println(data)
	//errorNumber := 0
	//email := data.Email
	//err := a.db.Where("email = ?", email).First(&user).Error
	//if err != nil {
	//	errorNumber = 1
	//	return errorNumber
	//}

	//if data.Code == user.Code {
	//
	//	hash, _ := hashPassword(data.NewPassword)
	//	a.db.Model(&user).Update("password", hash)
	//	a.db.Model(&user).Update("code", nil)
	//	return errorNumber
	//}
	//errorNumber = 2
	//return errorNumber
	return data, err
}
