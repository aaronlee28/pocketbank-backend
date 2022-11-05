package httpsuccess

import (
	"net/http"
)

type AppSuccess struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (success AppSuccess) Success() string {
	return success.Message
}

func OkSuccess(message string, data interface{}) AppSuccess {
	return AppSuccess{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
}

func CreatedSuccess(message string, data interface{}) AppSuccess {
	return AppSuccess{
		StatusCode: http.StatusCreated,
		Message:    message,
		Data:       data,
	}
}
