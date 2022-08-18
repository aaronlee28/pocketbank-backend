package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Handler) TopupSavings(c *gin.Context) {

	payload, _ := c.Get("payload")
	payload2, _ := c.Get("user")
	param, _ := payload.(*dto.TopupSavingsReq)
	user, _ := payload2.(models.User)
	userid := user.Id

	result, err := a.TransactionService.TopupSavings(param, userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) Transfer(c *gin.Context) {

	payload, _ := c.Get("payload")
	payload2, _ := c.Get("user")
	param, _ := payload.(*dto.TransferReq)
	user, _ := payload2.(models.User)
	userid := user.Id

	result, err := a.TransactionService.Transfer(param, userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) RunCronJobs(c *gin.Context) {

	a.TransactionService.RunCronJobs()

	c.Next()

}

func (a *Handler) TopupDeposit(c *gin.Context) {

	payload, _ := c.Get("payload")
	payload2, _ := c.Get("user")
	param, _ := payload.(*dto.TopupDepositReq)
	user, _ := payload2.(models.User)
	userid := user.Id

	result, err := a.TransactionService.TopupDeposit(param, userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}
