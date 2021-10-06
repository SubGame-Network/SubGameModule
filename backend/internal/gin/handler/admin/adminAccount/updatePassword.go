package adminAccount

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type requestUpdatePassword struct {
	Password string `json:"password" binding:"required"`
}

func UpdatePassword(c *gin.Context) {
	var req requestUpdatePassword
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v, req: %v", err, string(reqBody))
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	err := checkPassword(req.Password)
	if err != nil {
		zap.S().Infof("err: %v req: %v", err, c.Request)
		handler.Failed(c, domain.ErrorBadRequest, domain.ErrorPasswordRules.Message)
		return
	}

	adminUUID := uuid.FromStringOrNil(c.MustGet("AdminUUID").(string))

	adminRepo, err := provider.NewAdminRepo()
	if err != nil {
		zap.S().Warn(err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	_, err = adminRepo.GetUserByUUID(adminUUID)
	if err != nil {
		zap.S().Info(err)
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	password := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))

	err = adminRepo.UpdatePasswordByUUID(adminUUID, password)
	if err != nil {
		zap.S().Warn(err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	handler.Success(c, gin.H{})
}

func checkPassword(password string) error {

	var CheckPasswordA2Z = regexp.MustCompile(`[A-Z]{1}`)
	var CheckPassworda2z = regexp.MustCompile(`[a-z]{1}`)
	var CheckPassword0to9 = regexp.MustCompile(`[0-9]{1}`)

	if !CheckPasswordA2Z.MatchString(password) || !CheckPassworda2z.MatchString(password) || !CheckPassword0to9.MatchString(password) {
		return fmt.Errorf("Password rules do not match, the password must contain an English uppercase, an English lowercase, and a number")
	}

	return nil
}
