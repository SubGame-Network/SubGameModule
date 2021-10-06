package front

import (
	"bytes"
	"io/ioutil"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserEditNameRequest struct {
	NickName    string `json:"nickName"`
	Country     string `json:"country"`     // 城市
	CountryCode string `json:"countryCode"` // 國碼
	Phone       string `json:"phone"`
}

func UserEditName(c *gin.Context) {
	userAddress := c.GetString("userAddress")

	var req UserEditNameRequest
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v, req: %v", err, string(reqBody))
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

	// 更改名字
	user.NickName = req.NickName
	user.Country = req.Country
	user.CountryCode = req.CountryCode
	user.Phone = req.Phone
	err = repo.UpdateUser(user)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	handler.Success(c, gin.H{})
}
