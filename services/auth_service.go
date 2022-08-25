package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type AuthService interface {
	Register(req *dto.RegReq) (*dto.RegRes, error)
	SignIn(req *dto.AuthReq) (*dto.TokenRes, error)
	GetCode(email *dto.CodeReq) (*dto.CodeRes, error)
	ChangePassword(data *dto.ChangePReq) (*dto.ChangePRes, error)
}

type authService struct {
	authRepository repositories.AuthRepository
	appConfig      config.AppConfig
}

type ASConfig struct {
	AuthRepository repositories.AuthRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *ASConfig) *authService {
	return &authService{
		authRepository: c.AuthRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	Scope string
	User  *models.User `json:"user"`
}

func (a *authService) generateJWTToken(user *models.User) (*dto.TokenRes, error) {

	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}
	claims := &idTokenClaims{

		jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			ExpiresAt: &timeExpire,
			IssuedAt:  &timeNow,
		},
		user.Role,
		&models.User{Id: user.Id, Email: user.Email},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hmacSampleSecret := "very-secret"
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	if err != nil {
		return nil, err
	}

	res := &dto.TokenRes{IDToken: tokenString}
	return res, nil
}

func (a *authService) Register(req *dto.RegReq) (*dto.RegRes, error) {

	u := &models.User{
		Name:     req.Name,
		Email:    &req.Email,
		Password: req.Password,
		Contact:  req.Contact,
	}
	checkReferral := req.ReferralNumber
	_, err := a.authRepository.Register(u, checkReferral)
	if err != nil {
		if strings.Contains(err.Error(), "users_contact_uindex") {
			return nil, error(httperror.BadRequestError("Contact is already registered", "401"))
		}
		if strings.Contains(err.Error(), "users_email_uindex") {
			return nil, error(httperror.BadRequestError("Email is already registered", "401"))
		}
		if strings.Contains(err.Error(), "record not found") {
			return nil, error(httperror.BadRequestError("Referral code not registered", "401"))
		}

	}
	res := &dto.RegRes{Email: req.Email, Name: req.Name, Contact: req.Contact}
	return res, nil
}

func (a *authService) SignIn(req *dto.AuthReq) (*dto.TokenRes, error) {
	user, err, isUserActive := a.authRepository.MatchingCredential(req.Email, req.Password)
	if isUserActive == false {
		return nil, httperror.AppError{
			StatusCode: http.StatusForbidden,
			Code:       "User Account Inactive",
			Message:    "User Account Inactive",
		}
	}
	if err != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Incorrect email or password",
		}
	}
	token, err := a.generateJWTToken(user)
	return token, err
}

func (a *authService) GetCode(email *dto.CodeReq) (*dto.CodeRes, error) {
	stringEmail := email.Email
	user, code, err := a.authRepository.GetCode(stringEmail)
	if err != nil || user == nil {
		return nil, error(httperror.BadRequestError("Email is not found", "400"))
	}
	returncode := &dto.CodeRes{
		Code: code,
	}
	return returncode, err
}

func (a *authService) ChangePassword(data *dto.ChangePReq) (*dto.ChangePRes, error) {
	errNumber := a.authRepository.ChangePassword(data)
	if errNumber == 1 {
		return nil, error(httperror.BadRequestError("Email is not found", "400"))
	}
	if errNumber == 2 {
		return nil, error(httperror.BadRequestError("Code Invalid", "401"))

	}
	returnres := &dto.ChangePRes{
		Success: "Successfully Change Password",
	}
	return returnres, nil
}
