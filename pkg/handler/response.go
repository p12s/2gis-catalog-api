package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// errorResponse - формат ответа для ошибки
type errorResponse struct {
	Message string `json:"message"`
}

// newErrorResponse - отправка ошибки
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
