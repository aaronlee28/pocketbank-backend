package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Handler) Topup(c *gin.Context) {

	payload, _ := c.Get("payload")
	payload2, _ := c.Get("user")
	param, _ := payload.(*dto.TopupReq)
	user, _ := payload2.(models.User)
	userid := user.Id

	result, err := a.walletService.Topup(param, userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) Transaction(c *gin.Context) {
	payload, _ := c.Get("user")
	user, _ := payload.(models.User)
	userid := user.Id

	query := &repositories.Query{
		SortBy: c.Query("sortBy"),
		Sort:   c.Query("sort"),
		Limit:  c.Query("limit"),
		Page:   c.Query("page"),
		Search: c.Query("search"),
	}

	result, err := a.walletService.Transaction(query, userid)

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

	result, err := a.walletService.Transfer(param, userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) UserDetails(c *gin.Context) {

	payload, _ := c.Get("user")
	user, _ := payload.(models.User)
	userid := user.Id

	result, err := a.walletService.UserDetails(userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) UpdateInterestAndTax(c *gin.Context) {

	a.walletService.UpdateInterestAndTax()

	c.JSON(http.StatusOK, nil)

}

func (a *Handler) RunCronJobs(c *gin.Context) {

	a.walletService.RunCronJobs()

	c.JSON(http.StatusOK, nil)

}
