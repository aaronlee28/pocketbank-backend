package handlers

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *Handler) Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	contact := c.PostForm("contact")
	password := c.PostForm("password")
	referralNumber := c.PostForm("referralNumber")
	referralNumberInt, _ := strconv.Atoi(referralNumber)

	param := &dto.RegReq{
		Name:           name,
		Email:          email,
		Contact:        contact,
		Password:       password,
		ReferralNumber: referralNumberInt,
	}
	result, err := a.AuthService.Register(param)
	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)

		return
	}

	successResponse := httpsuccess.CreatedSuccess("Created", result)
	c.JSON(http.StatusCreated, successResponse)
}

func (a *Handler) SignIn(c *gin.Context) {
	payload, _ := c.Get("payload")
	signin, _ := payload.(*dto.AuthReq)

	result, err1 := a.AuthService.SignIn(signin)

	if err1 != nil {
		e := c.Error(err1)
		c.JSON(http.StatusBadRequest, e)
		return
	}
	if result == nil {
		e2 := c.Error(httperror.BadRequestError("users not found", ""))
		c.JSON(http.StatusBadRequest, e2)
		return
	}
	fmt.Println(result)
	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) GetCode(c *gin.Context) {
	payload, _ := c.Get("payload")
	email, _ := payload.(*dto.CodeReq)
	result, err := a.AuthService.GetCode(email)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.CreatedSuccess("Created", result)
	c.JSON(http.StatusCreated, successResponse)

}

func (a *Handler) ChangePassword(c *gin.Context) {
	payload, _ := c.Get("payload")
	data, _ := payload.(*dto.ChangePReq)

	result, err := a.AuthService.ChangePassword(data)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}
