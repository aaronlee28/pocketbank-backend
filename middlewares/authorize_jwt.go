package middlewares

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httperror.UnauthorizedError()
		}
		return config.Config.JWTSecret, nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	if config.Config.ENV == "testing" {
		fmt.Println("disable JWT authorization on dev env")
		return
	}

	authHeader := c.GetHeader("Authorization")

	s := strings.Split(authHeader, "Bearer ")

	unauthorizedErr := httperror.UnauthorizedError()
	if len(s) < 2 {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	encodedToken := s[1]
	token, err := validateToken(encodedToken)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	userJson, err := json.Marshal(claims["user"])
	var user models.User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	c.Set("user", user)
}

func AuthorizeAdmin(c *gin.Context) {

	if config.Config.ENV == "testing" {
		fmt.Println("disable Admin authorization on dev env")
		return
	}

	authHeader := c.GetHeader("Authorization")

	s := strings.Split(authHeader, "Bearer ")

	forbiddenErr := httperror.ForbiddenError()
	encodedToken := s[1]
	token, _ := validateToken(encodedToken)
	claims, _ := token.Claims.(jwt.MapClaims)
	scopeJson, _ := json.Marshal(claims["Scope"])
	var scope interface{}
	json.Unmarshal(scopeJson, &scope)
	if scope != "user admin" {
		c.AbortWithStatusJSON(forbiddenErr.StatusCode, forbiddenErr)
		return
	}
	return
}
