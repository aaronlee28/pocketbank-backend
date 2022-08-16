package services

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/assignment-05-golang-backend/config"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/assignment-05-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/assignment-05-golang-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/assignment-05-golang-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/assignment-05-golang-backend/repositories"
	"net/http"
	"time"
)

type AuthService interface {
	Register(req *dto.AuthReq) (*dto.RegRes, error)
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
	User *models.User `json:"user"`
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

func (a *authService) Register(req *dto.AuthReq) (*dto.RegRes, error) {

	u := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := a.authRepository.Register(u)
	if err != nil {
		return nil, error(httperror.BadRequestError("Failed to register account", ""))
	}
	res := &dto.RegRes{Success: "Success"}
	return res, nil
}

func (a *authService) SignIn(req *dto.AuthReq) (*dto.TokenRes, error) {
	user, err := a.authRepository.MatchingCredential(req.Email, req.Password)
	if err != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Unauthorized",
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
