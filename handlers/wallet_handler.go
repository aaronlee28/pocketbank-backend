package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Handler) Transaction(c *gin.Context) {
	payload, _ := c.Get("user")
	user, _ := payload.(models.User)
	userid := user.Id

	query := &repositories.Query{
		SortBy:     c.Query("sortBy"),
		Sort:       c.Query("sort"),
		Limit:      c.Query("limit"),
		Page:       c.Query("page"),
		Search:     c.Query("search"),
		FilterTime: c.Query("filterTime"),
		MinAmount:  c.Query("minAmount"),
		MaxAmount:  c.Query("maxAmount"),
	}

	result, err := a.WalletService.Transaction(query, userid)

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

	result, err := a.WalletService.UserDetails(userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (a *Handler) DepositInfo(c *gin.Context) {

	payload, _ := c.Get("user")
	user, _ := payload.(models.User)
	userid := user.Id

	result, err := a.WalletService.DepositInfo(userid)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	c.JSON(http.StatusOK, result)

}
