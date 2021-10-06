package middleware

import (
	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		notificationService, _ := provider.NewNotificationService()
	// 		slackErr := notificationService.SlackNotifyError(context.Background(), fmt.Sprintf("%v", err))
	// 		if slackErr != nil {
	// 			zap.S().Error(slackErr)
	// 		}
	//
	// 		zap.S().Error(err)
	// 		handler.Failed(c, domain.ErrorServer)
	// 		return
	// 	}
	// }()
	c.Next()
}
