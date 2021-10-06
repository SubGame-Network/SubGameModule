package route

import (
	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler/admin/adminAccount"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler/front"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	if config.Gin.Mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.InitLogger)
	// r.Use(middleware.Recovery)

	r.Static("/api/assets", "./assets")

	api := r.Group("/api")
	{
		adminV1 := api.Group("/admin/v1")
		adminV1.Use(middleware.AdminAuthMiddleware)
		{
			adminV1.POST("/login", adminAccount.Login)
			adminV1.PATCH("/password", adminAccount.UpdatePassword)
		}

		frontV1 := api.Group("/v1")
		{

			frontV1.POST("/user/email/send", front.UserEmail)
			frontV1.POST("/user/join", front.UserJoin)

			// Module
			frontV1.GET("/module", front.Module)
			frontV1.GET("/module/:id", front.ModuleDetail)
			frontV1.Use(middleware.FrontAuthMiddleware)
			{
				frontV1.POST("/contact", front.Contact)
				frontV1.GET("/user", front.User)
				frontV1.GET("/stake/record", front.StakeRecord)
				frontV1.PATCH("/user/name", front.UserEditName) // 後面新增可編輯國碼、電話
			}
		}
	}
	return r
}
