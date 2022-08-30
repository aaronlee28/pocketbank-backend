package repositories

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
	"reflect"
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
	UserRate(id int, data *dto.ChangeInterestRateReq) error
	UsersRate(data *dto.ChangeInterestRateReq) error
	CreatePromotion(data *dto.PromotionReq) (*models.Promotion, error)
	GetPromotion() (*[]models.Promotion, error)
	UpdatePromotion(id int, data *dto.PatchPromotionReq) (*dto.PatchPromotionReq, error)
	DeletePromotion(id int) (*models.Promotion, error)
	EligibleMerchandiseList() (*[]models.Merchandise, error)
	MerchandiseStatus(data *dto.MerchandiseStatus) (error, error, int)
	UpdateMerchStocks(data *dto.UpdateMerchStocksReq) (*models.Merchstock, error)
	GetMerchStock() (*[]models.Merchstock, error)
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
		Email:          *user.Email,
		Contact:        user.Contact,
		ProfilePicture: user.ProfilePicture,
		ReferralNumber: *user.ReferralNumber,
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
	err := w.db.Where("user_id = ?", id).Where("deleted_at is null").Order("created_at desc").Find(&m).Error
	return m, err
}

func (w *adminRepository) UserRate(id int, data *dto.ChangeInterestRateReq) error {
	var d *[]models.Deposit
	err := w.db.Where("user_id = ?", id).Find(&d).Error

	for _, deposit := range *d {

		w.db.Model(deposit).Update("interest_rate", data.InterestRate)

	}
	return err
}

func (w *adminRepository) UsersRate(data *dto.ChangeInterestRateReq) error {
	var d *[]models.Deposit
	err := w.db.Find(&d).Error

	for _, deposit := range *d {

		w.db.Model(deposit).Update("interest_rate", data.InterestRate)

	}
	return err
}

func (w *adminRepository) CreatePromotion(data *dto.PromotionReq) (*models.Promotion, error) {
	p := &models.Promotion{
		Title: data.Title,
		Photo: data.Photo,
	}

	err := db.Get().Create(&p).Error

	return p, err
}

func (w *adminRepository) GetPromotion() (*[]models.Promotion, error) {
	var p *[]models.Promotion
	err := w.db.Where("deleted_at is null").Find(&p).Error
	return p, err
}

func (w *adminRepository) UpdatePromotion(id int, data *dto.PatchPromotionReq) (*dto.PatchPromotionReq, error) {

	var p *models.Promotion
	err := w.db.Where("id = ?", id).First(&p).Error
	pho := p.Photo
	v := reflect.ValueOf(*data)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		if val != "" && val != 0 {
			change := v.Type().Field(i).Name
			input := v.Field(i).Interface()
			w.db.Model(&p).Update(change, input)
		}
	}
	if data.Photo == "null" {

		w.db.Model(&p).Update("photo", pho)

	}
	return data, err
}

func (w *adminRepository) DeletePromotion(id int) (*models.Promotion, error) {
	var p *models.Promotion
	err := w.db.Where("id = ?", id).Find(&p).Error
	w.db.Model(&p).Update("deleted_at", time.Now())

	return p, err
}

func (w *adminRepository) EligibleMerchandiseList() (*[]models.Merchandise, error) {
	var m *[]models.Merchandise
	err := w.db.Where("pen = true or umbrella = true or card_holder = true").Find(&m).Error

	return m, err
}

func (w *adminRepository) MerchandiseStatus(data *dto.MerchandiseStatus) (error, error, int) {
	var m *models.Merchandise
	var s *models.Merchstock
	err1 := w.db.Where("user_id = ?", data.UserID).First(&m).Error
	send := "send_" + data.MerchToSend
	err2 := w.db.Model(&m).Update(send, data.Status).Error
	if data.Status == "On process" {
		w.db.Where("name = ?", data.MerchToSend).First(&s)

		if s.StockCount <= 0 {
			return nil, nil, 1
		}
		newStock := s.StockCount - 1

		w.db.Model(&s).Update("stock_count", newStock)
	}

	return err1, err2, 0
}

func (w *adminRepository) UpdateMerchStocks(data *dto.UpdateMerchStocksReq) (*models.Merchstock, error) {
	var m *models.Merchstock
	err := w.db.Where("name = ?", data.Name).First(&m).Error
	w.db.Model(&m).Update("stock_count", data.Stock)

	return m, err
}

func (w *adminRepository) GetMerchStock() (*[]models.Merchstock, error) {
	var m *[]models.Merchstock
	err := w.db.Find(&m).Error

	return m, err
}
