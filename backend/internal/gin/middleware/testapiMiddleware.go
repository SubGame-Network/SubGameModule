package middleware

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
)

func APITestMiddleware(c *gin.Context) {
	config := provider.NewConfig()

	if config.Gin.Mode != "DEBUG" {
		handler.Failed(c, domain.ErrorForbidden, domain.ErrorForbidden.Message)
		c.Abort()
		return
	}

	c.Next()
}
