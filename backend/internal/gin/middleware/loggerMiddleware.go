package middleware

import (
	"github.com/SubGame-Network/SubGameModuleService/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitLogger(c *gin.Context) {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync() // flushes buffer, if any
	c.Next()
}
