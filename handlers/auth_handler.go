package handlers

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Handler) Register(c *gin.Context) {
	payload, _ := c.Get("payload")
	param, _ := payload.(*dto.RegReq)
	result, err := a.authService.Register(param)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) SignIn(c *gin.Context) {
	payload, _ := c.Get("payload")
	signin, _ := payload.(*dto.AuthReq)

	result, err1 := a.authService.SignIn(signin)

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
	//c.(http.StatusOK, result)

}

func (a *Handler) GetCode(c *gin.Context) {
	payload, _ := c.Get("payload")
	email, _ := payload.(*dto.CodeReq)
	result, err := a.authService.GetCode(email)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) ChangePassword(c *gin.Context) {
	payload, _ := c.Get("payload")
	data, _ := payload.(*dto.ChangePReq)

	result, err := a.authService.ChangePassword(data)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}
