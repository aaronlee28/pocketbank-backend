package middlewares

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/assignment-05-golang-backend/httperror"
	"github.com/gin-gonic/gin"
	"reflect"
)

func RequestValidator(model any) gin.HandlerFunc {
	return func(c *gin.Context) {
		modelPtr := reflect.ValueOf(model).Elem()
		modelPtr.Set(reflect.Zero(modelPtr.Type()))
		if err := c.ShouldBindJSON(model); err != nil {
			badRequest := httperror.BadRequestError(err.Error(), "")
			c.AbortWithStatusJSON(badRequest.StatusCode, badRequest)
			return
		}
		c.Set("payload", model)
	}
}
