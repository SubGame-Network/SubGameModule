package service

import (
	"encoding/json"
	"strconv"
	"time"

	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

func (svc *NotifyService) NotifySlackError(text string) {
	if svc.Config.Notify.ServiceName == "" {
		return
	}

	attachment := slack.Attachment{
		Color: "danger",
		Title: svc.Config.Notify.ServiceName,
		Text:  text,
		Ts:    json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(svc.Config.Notify.SlackWebhook, &msg)
	if err != nil {
		zap.S().Warn(err)
	}
}

func (svc *NotifyService) NotifyLog(text string) error {
	zap.S().Info(text)

	if svc.Config.Notify.ServiceName != "SubGameModuleService-prod" {
		// 測試站，凌晨3點前 不發通知到slack
		if time.Now().Hour() < 3 {
			return nil
		}
	}

	// slack
	attachment := slack.Attachment{
		Color: "good",
		Title: svc.Config.Notify.ServiceName,
		Text:  text,
		Ts:    json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.PostWebhook(svc.Config.Notify.SlackLogWebhook, &msg)
	if err != nil {
		zap.S().Warn(err)
	}

	// telegram
	if svc.Config.Notify.TelegramApiToken == "" || svc.Config.Notify.TelegramChatId == 0 {
		return nil
	}
	bot, err := tgbotapi.NewBotAPI(svc.Config.Notify.TelegramApiToken)
	if err != nil {
		zap.S().Warn(err)
	}
	if err == nil && svc.Config.Notify.TelegramChatId != 0 {
		msg := tgbotapi.NewMessage(svc.Config.Notify.TelegramChatId, fmt.Sprintf("%s", text))
		_, err := bot.Send(msg)
		if err != nil {
			zap.S().Warn(err)
		}
	}

	return nil
}
