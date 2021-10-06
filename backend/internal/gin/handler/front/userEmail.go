package front

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

func UserEmail(c *gin.Context) {
	var req UserEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v, req: %v", err, c.Request.Body)
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	if !domain.RegexEmail.MatchString(req.Email) {
		handler.Failed(c, domain.ErrorEmailFormat, "")
		return
	}

	req.Email = strings.ToLower(req.Email)
	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	user, err := repo.GetUserByEmail(req.Email)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	if user != nil {
		handler.Failed(c, domain.ErrorAlreadyExistsDB, "")
		return
	}

	// 資料都匹配且還沒驗正完，準備發信
	// 產生驗證碼
	rand.Seed(time.Now().UnixNano()) // set seed
	min := 100000
	max := 999999
	verifyCode := rand.Intn(max-min) + min

	// set redis
	redis, err := provider.NewRedis()
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	config := provider.NewConfig()

	// 驗證碼
	key := domain.GetRedisKeyUserEmailVerify(req.Email)
	ttl := 60 * 10 // 10分鐘
	if err = redis.SetEx(key, verifyCode, ttl); err != nil {
		zap.S().Warn(err)
		handler.Failed(c, domain.ErrorUnauthorized, "")
		return
	}

	smtpService, err := provider.NewSmtpService()
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	status := domain.RedisStatusEmailVerify
	// // 判斷IP是否發送達上限
	// ok1 := smtpService.CheckEmailSendingLimitOK(req.Email, status)
	// if !ok1 {
	// 	handler.Failed(c, domain.ErrorEmailSendingLimit, "")
	// 	return
	// }
	// 判斷60秒內是否有重複送
	ok2 := smtpService.CheckEmailSendingSecLimitOk(req.Email, status)
	if !ok2 {
		handler.Failed(c, domain.ErrorEmailSendingin60secLimit, "")
		return
	}

	email := fmt.Sprintf("support%s", config.Email.Domain)
	ReplyTo := fmt.Sprintf("support%s", config.Email.Domain)
	sender := domain.Sender{Name: "SubgameModule", Email: email, ReplyTo: ReplyTo}
	content, err := ioutil.ReadFile("internal/smtp/smtp_tamplate/template_email_verify_code.html")
	if err != nil {
		zap.S().Warn("ReadFile Error", err)
		handler.Failed(c, domain.ErrorServer, "ReadFile Error")
		return
	}

	// 參數取代
	htmlContent := string(content)
	htmlContent = strings.ReplaceAll(htmlContent, "${VerifyCode}", strconv.FormatInt(int64(verifyCode), 10))

	// 發送人
	to := req.Email

	// 舊SMTP服務發信
	if err := smtpService.Send("Please verify your email", []string{to}, &sender, htmlContent); err != nil {
		senderJson, _ := json.Marshal(sender)
		zap.S().Warn("ForgetPassword smtp寄信失敗 Error", err)
		zap.S().Infow(`發送信息如下：`,
			"To", string(to),
			"sender", string(senderJson))
		handler.Failed(c, domain.ErrorServer, "smtp發送失敗")
		return
	}

	// 紀錄發送次數
	smtpService.SetEmailSendingLimit(status, req.Email)
	smtpService.SetEmailSendingSecLimit(status, req.Email)

	handler.Success(c, gin.H{})
}
