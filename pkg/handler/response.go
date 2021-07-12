package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// errorResponse - формат ответа для ошибки
type errorResponse struct {
	Message string `json:"message"`
}

// statusResponse - формат ответа для успешного статуса
type statusResponse struct {
	Status string `json:"status"`
}

// newErrorResponse - отправка ошибки
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
