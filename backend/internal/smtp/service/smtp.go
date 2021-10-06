package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SmtpService interface {
	Send(subject string, toSlice []string, sender *domain.Sender, HtmlContent string) error
	AWSSend(subject string, toSlice []string, HtmlBody string) error
	SetEmailSendingLimit(status string, IPAddress string) error
	CheckEmailSendingLimitOK(IPAddress string, status string) bool
	GetEmailSendingCount(IPAddress string, status string) (int, int)
	SetEmailSendingSecLimit(status string, IPAddress string) error
	CheckEmailSendingSecLimitOk(IPAddress string, status string) bool
	apiSendMail(method string, data interface{}) error
}

type smtp struct {
	config *config.Config
	redis  domain.GoRedis
}

func NewSmtpService(config *config.Config, redisServ domain.GoRedis) SmtpService {
	return &smtp{config, redisServ}
}

func (s *smtp) Send(subject string, toSlice []string, sender *domain.Sender, HtmlContent string) error {
	var to []domain.Email
	// slice to Email struct
	for i := 0; i < len(toSlice); i++ {
		to = append(to, domain.Email{toSlice[i]})
	}

	payload := domain.Payload{
		Sender:      sender,
		To:          to,
		RepleyTo:    domain.Email{sender.ReplyTo},
		HtmlContent: HtmlContent,
		Subject:     "[SubgameModule]" + subject,
	}
	err := s.apiSendMail("POST", payload)
	if err != nil {
		zap.S().Warn("smtp server response error", err)
		return err
	}
	toString := strings.Join(toSlice, ",")
	zap.S().Infow(`Smtp 信件發送成功！`, "To", toString)
	return nil
}

func (s *smtp) AWSSend(subject string, toSlice []string, HtmlBody string) error {
	region := s.config.AwsSES.Region
	accessID := s.config.AwsSES.AccessId
	secretKey := s.config.AwsSES.SecretKey
	sender := s.config.AwsSES.Sender
	charSet := s.config.AwsSES.CharSet

	toAddress := []*string{}
	for _, val := range toSlice {
		toAddress = append(toAddress, aws.String(val))
	}

	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accessID, secretKey, ""),
		},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: toAddress,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(HtmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	_, err = svc.SendEmail(input)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *smtp) apiSendMail(method string, data interface{}) error {
	config := s.config
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, config.Email.API_URL, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", config.Email.API_KEY)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	sucRes := struct {
		MessageId string `json:"messageId"`
	}{}
	if err := json.Unmarshal(body, &sucRes); err != nil || sucRes.MessageId == "" {
		return err
	}
	return nil
}

//*************************************************************
//**************************附加 功能***************************
//*************************************************************
// 1-1. 『紀錄』限制每日發送上限	SetEmailSendingLimit
// 1-2. 『檢查』限制每日發送上限	CheckEmailSendingLimitOK
// 1-2. 『取得』目前發送次數		GetEmailSendingCount
// 1-1. 『設定』限制60秒內重複發送	SetEmailSendingSecLimit
// 1-2. 『檢查』限制60秒內重複發送	CheckEmailSendingSecLimitOk

// ToDo use config []limit_while_list
var (
	whiteList = []string{}
)

// limitWhileList

const (
	EmailSendingLimitPrefix    = "EmailSendingLimit"
	EmailSendingSecLimitPrefix = "EmailSendingSecLimit"
)

func (s *smtp) SetEmailSendingLimit(status string, IPAddress string) error {
	redis := s.redis

	// key1 3次限制的
	key := fmt.Sprintf("%s:%s:%s", EmailSendingLimitPrefix, status, IPAddress)
	// 紀錄每日發送上限
	err := redis.INCR(key)
	if err != nil {
		zap.S().Warn("redis set incr error", err)
		return err
	}

	// 計算距離明日凌晨與現在 時間差
	time1 := time.Now()
	tomorrowDate := time1.Unix() - int64(time1.Second()) - int64(time1.Minute()*60) - int64(time1.Hour()*60*60) + (60 * 60 * 24)
	// 時間差
	ttl := tomorrowDate - time1.Unix()
	if err := redis.Expire(key, int(ttl)); err != nil {
		zap.S().Warn(err.Error())
		return err
	}
	return nil
}

func (s *smtp) CheckEmailSendingLimitOK(IPAddress string, status string) bool {
	for _, ip := range whiteList {
		ipRegexp := regexp.MustCompile(ip)
		if ipRegexp.MatchString(IPAddress) {
			return true
		}
	}

	config := s.config

	// key1 3次限制的
	key1 := fmt.Sprintf("%s:%s:%s", EmailSendingLimitPrefix, status, IPAddress)
	redis := s.redis
	count, _ := redis.Get(key1)
	limit := config.Email.AgentVerifyMailDayLimit
	if count >= limit {
		zap.S().Warn("該Email到達上限" + IPAddress)
		return false
	}
	return true
}

func (s *smtp) GetEmailSendingCount(IPAddress string, status string) (int, int) {
	key := fmt.Sprintf("%s:%s:%s", EmailSendingLimitPrefix, status, IPAddress)
	redis := s.redis
	countStr, _ := redis.Get(key)
	count, _ := strconv.Atoi(countStr)

	config := s.config
	limit, _ := strconv.Atoi(config.Email.AgentVerifyMailDayLimit)
	return count, (limit - count)
}

func (s *smtp) SetEmailSendingSecLimit(status string, IPAddress string) error {
	redis := s.redis
	// key2 60秒內
	key := fmt.Sprintf("%s:%s:%s", EmailSendingSecLimitPrefix, status, IPAddress)
	// 紀錄60內不能重複發送
	if err := redis.SetEx(key, "ping", 60); err != nil {
		return err
	}
	return nil
}

func (s *smtp) CheckEmailSendingSecLimitOk(IPAddress string, status string) bool {
	redis := s.redis
	// key2 60秒內
	key2 := fmt.Sprintf("%s:%s:%s", EmailSendingSecLimitPrefix, status, IPAddress)
	str, _ := redis.Get(key2)
	if str != "" {
		zap.S().Info("60秒內不能重複發送" + IPAddress)
		return false
	}

	return true
}
