package repositories

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"time"
)

type WalletRepository interface {
	TransactionHistory(q *Query, id int) (int, *[]models.Transaction, error)
	UserDetails(id int) (*dto.UserDetailsRes, error)
	DepositInfo(id int) (*[]models.Deposit, error)
	SavingsInfo(id int) (*models.Savings, error)
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
	Type       string
}

func NewWalletRepository(c *WRConfig) walletRepository {
	return walletRepository{db: c.DB}
}

func (w *walletRepository) TransactionHistory(q *Query, id int) (int, *[]models.Transaction, error) {
	var trans *[]models.Transaction
	var account *models.Savings
	limit, _ := strconv.Atoi(q.Limit)
	page, _ := strconv.Atoi(q.Page)
	search := "%" + q.Search + "%"
	ty := "%" + q.Type + "%"
	offset := (limit * page) - limit
	w.db.Where("user_id = ?", id).First(&account)
	query := w.db.Order(q.SortBy+" "+q.Sort).Where("sender_wallet_number = ? OR receiver_wallet_number = ? ", account.SavingsNumber, account.SavingsNumber).Where("UPPER(description) like UPPER(?)", search).Where("created_at >= ? at time zone 'UTC' - interval '"+q.FilterTime+"' day", time.Now()).Where("amount BETWEEN ? and ?", q.MinAmount, q.MaxAmount).Where("type like ?", ty).Find(&trans)

	totalLength := len(*trans)

	w.db.Limit(limit).Offset(offset).Order(q.SortBy+" "+q.Sort).Where("sender_wallet_number = ? OR receiver_wallet_number = ? ", account.SavingsNumber, account.SavingsNumber).Where("UPPER(description) like UPPER(?)", search).Where("created_at >= ? at time zone 'UTC' - interval '"+q.FilterTime+"' day", time.Now()).Where("amount BETWEEN ? and ?", q.MinAmount, q.MaxAmount).Where("type like ?", ty).Find(&trans)
	fmt.Println("length", len(*trans))

	return totalLength, trans, query.Error
}

func (w *walletRepository) UserDetails(id int) (*dto.UserDetailsRes, error) {
	var user *models.User
	var sv *models.Savings
	err := w.db.Where("id = ?", id).First(&user).Error
	w.db.Where("user_id = ?", id).First(&sv)
	res := &dto.UserDetailsRes{
		Name:           user.Name,
		Email:          *user.Email,
		Contact:        user.Contact,
		ProfilePicture: user.ProfilePicture,
		ReferralNumber: *user.ReferralNumber,
		AccountNumber:  sv.SavingsNumber,
	}

	return res, err

}

func (w *walletRepository) DepositInfo(id int) (*[]models.Deposit, error) {
	var ds *[]models.Deposit
	err := w.db.Where("user_id = ? ", id).Where("deleted_at is null").Order("updated_at desc").Find(&ds).Error

	return ds, err
}

func (w *walletRepository) SavingsInfo(id int) (*models.Savings, error) {
	//var user *models.User
	var s *models.Savings
	err := w.db.Where("user_id = ? ", id).First(&s).Error

	return s, err
}

func (w *walletRepository) FavoriteContact(favoriteid int, selfid int) (*models.Favoritecontact, error) {
	var sv *models.Savings
	var fc *models.Favoritecontact
	err := w.db.Where("savings_number = ?", favoriteid).First(&sv).Error
	err2 := w.db.Where("favorite_account_number = ?", favoriteid).First(&fc).Error

	if err == nil && err2 != nil {
		addFavoriteContact := &models.Favoritecontact{
			UserID:                selfid,
			FavoriteAccountNumber: sv.SavingsNumber,
			Favorite:              true,
		}

		db.Get().Create(&addFavoriteContact)
		return addFavoriteContact, err
	}
	if err == nil && err2 == nil {
		if fc.Favorite == true {
			w.db.Model(&fc).Update("favorite", false)
			return fc, nil
		}
		if fc.Favorite == false {
			w.db.Model(&fc).Update("favorite", true)
			return fc, nil
		}
	}
	return nil, err
}

func (w *walletRepository) FavoriteContactList(id int) (*[]models.Favoritecontact, error) {
	var fc *[]models.Favoritecontact
	err := w.db.Where("user_id = ?", id).Find(&fc).Error

	return fc, err
}

func (w *walletRepository) ChangeUserDetails(data *dto.ChangeUserDetailsReqRes, id int) (*dto.ChangeUserDetailsReqRes, error) {
	var user *models.User
	err := w.db.Where("id = ?", id).First(&user).Error
	pho := user.ProfilePicture
	v := reflect.ValueOf(*data)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		if val != "" && val != 0 {
			change := v.Type().Field(i).Name
			input := v.Field(i).Interface()
			err2 := w.db.Model(&user).Update(change, input).Error
			if err2 != nil {
				return nil, err2
			}
		}
	}
	if data.ProfilePicture == "null" {
		w.db.Model(&user).Update("profile_picture", pho)
	}
	return data, err
}
