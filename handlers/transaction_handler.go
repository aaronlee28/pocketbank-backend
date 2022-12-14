package handlers

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	successResponse := httpsuccess.CreatedSuccess("Created", result)
	c.JSON(http.StatusCreated, successResponse)
}

func (a *Handler) Payment(c *gin.Context) {

	payload, _ := c.Get("payload")
	payload2, _ := c.Get("user")
	param, _ := payload.(*dto.PaymentReq)
	user, _ := payload2.(models.User)
	userid := user.Id
	result, err := a.TransactionService.Payment(param, userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.CreatedSuccess("Created", result)
	c.JSON(http.StatusCreated, successResponse)
}

func (a *Handler) RunCronJobs(c *gin.Context) {
	if config.Config.ENV == "testing" {
		fmt.Println("disable cron")
		return
	}

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

	successResponse := httpsuccess.CreatedSuccess("Created", result)
	c.JSON(http.StatusCreated, successResponse)
}

func (a *Handler) TopUpQr(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	payload, _ := c.Get("payload")
	param, _ := payload.(*dto.TopUpQr)

	result, err := a.TransactionService.TopUpQr(param, id)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.CreatedSuccess("Created", result)
	c.JSON(http.StatusCreated, successResponse)
}
