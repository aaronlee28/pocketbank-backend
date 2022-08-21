package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *Handler) AdminUsersList(c *gin.Context) {

	result, err := a.AdminService.AdminUsersList()

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) AdminUserTransaction(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
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

	result, err := a.AdminService.AdminUserTransaction(query, id)
	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) AdminUserDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := a.AdminService.AdminUserDetails(id)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) AdminUserReferralDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := a.AdminService.AdminUserReferralDetails(id)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) ChangeUserStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ret := a.AdminService.ChangeUserStatus(id)

	if ret != nil {
		e := c.Error(ret)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", ret)
	c.JSON(http.StatusOK, successResponse)
}
