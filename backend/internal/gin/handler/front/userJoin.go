package front

import (
	"strings"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	SubGameHelp "github.com/SubGame-Network/SubGameModuleService/internal/md/service"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserJoinRequest struct {
	Email       string `json:"email" binding:"required"`
	NickName    string `json:"nickName" binding:"required"`
	Country     string `json:"country"`
	Address     string `json:"address" binding:"required"`
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
	VerifyCode  string `json:"verifyCode" binding:"required"`
}

func UserJoin(c *gin.Context) {
	var req UserJoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v, req: %v", err, c.Request.Body)
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	req.Email = strings.ToLower(req.Email)
	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	if req.Address == "" {
		handler.Failed(c, domain.ErrorNotSubGameAddress, domain.ErrorNotSubGameAddress.Message)
		c.Abort()
		return
	}

	userAddress, err := SubGameHelp.SubGameAddressToPubkey(req.Address)
	if err != nil {
		handler.Failed(c, domain.ErrorNotSubGameAddress, domain.ErrorNotSubGameAddress.Message)
		c.Abort()
		return
	}
	userAddress = strings.TrimPrefix(userAddress, "0x")

	user, err := repo.GetUserByAddress(userAddress)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	if user != nil {
		handler.Failed(c, domain.ErrorAlreadyExistsDB, "")
		return
	}

	redis, err := provider.NewRedis()
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	// 取驗證碼
	key1 := domain.GetRedisKeyUserEmailVerify(req.Email)
	verifyCode, _ := redis.Get(key1)

	// 檢查還沒過期
	if verifyCode == "" {
		handler.Failed(c, domain.ErrorEmailExpired, "")
		return
	}

	// 檢查驗證碼正確
	if req.VerifyCode != verifyCode {
		handler.Failed(c, domain.ErrorVerifyCode, "")
		return
	}

	// 新增user
	err = repo.InsertUser(domain.User{
		NickName:    req.NickName,
		Country:     req.Country,
		Address:     userAddress,
		Email:       req.Email,
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
	})
	if err != nil {
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	// delete redis
	redis.Del(key1)
	handler.Success(c, gin.H{})
}
