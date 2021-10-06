package front

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserResponse struct {
	NickName    string `json:"nickName"`
	Country     string `json:"country"` // 城市
	Address     string `json:"address"`
	Email       string `json:"email"`
	CountryCode string `json:"countryCode"` // 國碼
	Phone       string `json:"phone"`
}

func User(c *gin.Context) {
	userAddress := c.GetString("userAddress")
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

	handler.Success(c, UserResponse{
		NickName:    user.NickName,
		Country:     user.Country,
		Address:     user.Address,
		Email:       user.Email,
		CountryCode: user.CountryCode,
		Phone:       user.Phone,
	})
	return
}
