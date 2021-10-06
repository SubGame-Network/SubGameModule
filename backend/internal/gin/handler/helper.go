package handler

import (
	"net/http"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func Failed(c *gin.Context, err domain.ErrorFormat, customMessage string) {
	message := err.Message
	if customMessage != "" {
		message = customMessage
	}

	switch err {
	case domain.ErrorServer:
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    err.Code,
			"message": message,
		})
	case domain.ErrorUnauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    err.Code,
			"message": message,
		})
	case domain.ErrorForbidden:
		c.JSON(http.StatusForbidden, gin.H{
			"code":    err.Code,
			"message": message,
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    err.Code,
			"message": message,
		})
	}
}
