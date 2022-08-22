package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/dto"
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

func (a *Handler) Merchandise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ret, err := a.AdminService.Merchandise(id)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", ret)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) UserDepositInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ret, err := a.AdminService.UserDepositInfo(id)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", ret)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) UserRate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	payload, _ := c.Get("payload")
	data, _ := payload.(*dto.ChangeInterestRateReq)

	err := a.AdminService.UserRate(id, data)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", err)
	c.JSON(http.StatusOK, successResponse)
}

func (a *Handler) UsersRate(c *gin.Context) {
	payload, _ := c.Get("payload")
	data, _ := payload.(*dto.ChangeInterestRateReq)

	err := a.AdminService.UsersRate(data)

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}

	successResponse := httpsuccess.OkSuccess("Ok", err)
	c.JSON(http.StatusOK, successResponse)
}
