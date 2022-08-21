package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/httpsuccess"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Handler) UsersList(c *gin.Context) {

	result, err := a.AdminService.UsersList()

	if err != nil {
		e := c.Error(err)
		c.JSON(http.StatusBadRequest, e)
		return
	}
	
	successResponse := httpsuccess.OkSuccess("Ok", result)
	c.JSON(http.StatusOK, successResponse)
}
