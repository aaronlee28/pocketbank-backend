package repositories

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AdminRepository interface {
	AdminUsersList() (*[]models.User, error)
	AdminUserTransaction(q *Query, id int) (*[]models.Transaction, error)
	AdminUserDetails(id int) (*dto.UserDetailsRes, error)
	AdminUserReferralDetails(id int) (*[]models.Referral, error)
	ChangeUserStatus(id int) error
	Merchandise(id int) (*models.Merchandise, error)
	UserDepositInfo(id int) (*[]models.Deposit, error)
}

type adminRepository struct {
	db *gorm.DB
}

type ADRConfig struct {
	DB *gorm.DB
}

func NewAdminRepository(c *ADRConfig) adminRepository {
	return adminRepository{db: c.DB}
}

func (w *adminRepository) AdminUsersList() (*[]models.User, error) {
	var u *[]models.User
	err := w.db.Find(&u).Error

	return u, err
}

func (w *adminRepository) AdminUserTransaction(q *Query, id int) (*[]models.Transaction, error) {
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

func (w *adminRepository) AdminUserDetails(id int) (*dto.UserDetailsRes, error) {
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

func (w *adminRepository) AdminUserReferralDetails(id int) (*[]models.Referral, error) {
	var user *models.User
	var rs *[]models.Referral
	err := w.db.Where("id = ?", id).First(&user).Error
	w.db.Where("referral_number = ?", user.ReferralNumber).Find(&rs)

	return rs, err
}

func (w *adminRepository) ChangeUserStatus(id int) error {
	var user *models.User
	err := w.db.Where("id = ?", id).First(&user).Error
	if user.IsActive == false {
		w.db.Model(&user).Update("is_active", true)
		return err
	}

	if user.IsActive == true {
		w.db.Model(&user).Update("is_active", false)
		return err
	}

	return err
}

func (w *adminRepository) Merchandise(id int) (*models.Merchandise, error) {
	var m *models.Merchandise
	err := w.db.Where("user_id = ?", id).First(&m).Error

	return m, err
}

func (w *adminRepository) UserDepositInfo(id int) (*[]models.Deposit, error) {
	var m *[]models.Deposit
	err := w.db.Where("user_id = ?", id).Order("updated_at").Find(&m).Error
	return m, err
}
