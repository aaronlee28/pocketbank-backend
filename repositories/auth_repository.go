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
	MatchingCredential(email, password string) (*models.User, error, bool)
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

	if cr != 0 {
		err := a.db.Where("referral_number = ?", cr).First(&checkUser).Error
		if err != nil {
			return nil, err
		}
	}

	hash, _ := hashPassword(user.Password)
	user.Password = hash
	user.Role = "user"
	user.IsActive = true
	user.ReferralNumber = nil
	err := db.Get().Create(&user).Error

	if err != nil {
		return nil, err
	}
	makeNewReferralCode := true
	for makeNewReferralCode == true {
		checkReferralNumber := rand.Intn(99999-9999) + 9999
		er := a.db.Model(&user).Update("referral_number", checkReferralNumber).Error
		if er == nil {
			makeNewReferralCode = false
		}
	}

	s := &models.Savings{
		UserID:        user.Id,
		SavingsNumber: 1 + rand.Intn(99999-10000) + 10000 + user.Id,
		Tax:           0.05,
		Interest:      0.2,
	}
	db.Get().Create(&s)

	m := &models.Merchandise{
		UserID:        user.Id,
		TotalTransfer: 0,
		Pen:           false,
		Umbrella:      false,
		CardHolder:    false,
	}
	db.Get().Create(&m)

	//has to be last because referral code might be there but failed to create the account for other reasons
	if cr != 0 {
		a.db.Where("user_id = ?", checkUser.Id).First(&referralBonus)
		referralPrice := referralBonus.Balance + 20000
		a.db.Model(&referralBonus).Update("balance", referralPrice)

		addTransaction := &models.Transaction{
			SenderWalletNumber:   5,
			ReceiverWalletNumber: referralBonus.SavingsNumber,
			Amount:               20000,
			Type:                 "Referral Payment",
			Description:          "Referral Payment",
		}
		db.Get().Create(&addTransaction)
		addReferral := &models.Referral{
			ReferralNumber: cr,
			UsedByUserID:   user.Id,
		}
		db.Get().Create(&addReferral)

	}

	return user, nil
}
func (a *authRepository) MatchingCredential(email, password string) (*models.User, error, bool) {
	var user *models.User
	err := a.db.Where("email = ?", email).First(&user).Error
	if user.IsActive == false {
		return nil, nil, false
	}
	hashed := user.Password
	check := checkPasswordHash(password, hashed)
	if err != nil || check == false {
		return nil, err, true
	}

	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, err, true
	}
	return user, err, true
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
