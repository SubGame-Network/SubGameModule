package front

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContactRequest struct {
	Type    string `json:"type" binding:"required"`
	Contact string `json:"contact" binding:"required"`
}

func Contact(c *gin.Context) {
	userAddress := c.GetString("userAddress")

	var req ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v, req: %v", err, c.Request.Body)
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	user, err := repo.GetUserByAddress(userAddress)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	if user == nil {
		handler.Failed(c, domain.ErrorUserNotFound, "")
		return
	}

	// 新增contact
	err = repo.InsertContact(domain.Contact{
		Type:    req.Type,
		Contact: req.Contact,
		UserId:  user.ID,
	})
	if err != nil {
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	handler.Success(c, gin.H{})
}
