package repositories

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
)

type AuthRepository interface {
	MatchingCredential(email, password string) (*models.User, error)
	Register(user *models.User, cr int) (*models.User, error)
	GetCode(email string) (*models.User, int, error)
	ChangePassword(data *dto.ChangePReq) int
}

type authRepository struct {
	db *gorm.DB
}

type ARConfig struct {
	DB *gorm.DB
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func NewAuthRepository(c *ARConfig) authRepository {
	return authRepository{db: c.DB}
}

func (a *authRepository) Register(user *models.User, cr int) (*models.User, error) {
	var checkUser *models.User
	var referralBonus *models.Savings
	err := a.db.Where("referral_number = ?", cr).First(&checkUser).Error
	if err != nil {
		return nil, err
	}

	hash, _ := hashPassword(user.Password)
	user.Password = hash
	user.EligibleMerchandise = false
	user.ReferralNumber = rand.Intn(99999-9999) + 9999
	res := db.Get().Create(&user)
	a.db.Model(&user).Update("code", nil)

	w := &models.Wallet{
		UserID:        user.Id,
		WalletNumber:  1 + rand.Intn(99999-10000) + 10000 + user.Id,
		SavingsNumber: 200000 + user.Id,
		DepositNumber: 300000 + user.Id,
	}
	db.Get().Create(&w)

	s := &models.Savings{
		UserID: user.Id,
	}
	db.Get().Create(&s)

	a.db.Where("user_id = ?", checkUser.Id).First(&referralBonus)
	referralPrice := referralBonus.Balance + 20000
	a.db.Model(&referralBonus).Update("balance", referralPrice)

	return user, res.Error
}
func (a *authRepository) MatchingCredential(email, password string) (*models.User, error) {
	var user *models.User
	err := a.db.Where("email = ?", email).First(&user).Error
	hashed := user.Password
	check := checkPasswordHash(password, hashed)
	if err != nil || check == false {
		return nil, err
	}

	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, err
	}
	return user, err
}

func (a *authRepository) GetCode(email string) (*models.User, int, error) {
	var user *models.User
	generateCode := rand.Intn(99999-9999) + 9999
	err := a.db.Where("email = ?", email).First(&user).Error

	a.db.Model(&user).Update("code", generateCode)

	return user, generateCode, err
}

func (a *authRepository) ChangePassword(data *dto.ChangePReq) int {
	var user *models.User
	errorNumber := 0
	email := data.Email
	err := a.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		errorNumber = 1
		return errorNumber
	}

	if data.Code == user.Code {

		hash, _ := hashPassword(data.NewPassword)
		a.db.Model(&user).Update("password", hash)
		a.db.Model(&user).Update("code", nil)
		return errorNumber
	}
	errorNumber = 2
	return errorNumber
}
