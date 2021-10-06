package config

import (
	"log"

	"os"
	"strings"

	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

type Config struct {
	Gin     *GinConfig
	DB      *DBConfig
	Jwt     *JwtConfig
	Redis   *RedisConfig
	SubGame *SubGame
	Other   *OtherConfig
	Notify  *Notify
	Email   *EmailConfig
	AwsSES  *AwsSES
}

type GinConfig struct {
	Host string
	Port string
	Mode string
}

type DBConfig struct {
	Dialect  string
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Flag     string
}

type JwtConfig struct {
	Issuer               string
	Secret               string
	Refresh_token_exp    int
	Access_token_exp_sec int
}

type RedisConfig struct {
	Host           string
	Port           int
	Database       int
	Auth           string
	Max_idle       int
	Max_active     int
	Idle_timeout   int
	Notify_active  int
	Polling_active int
}

type SubGame struct {
	RPC               string
	SS58Prefix        string
	ChainRuntimeTypes []byte
	EventKey          string
	MDOwnerPubkey     string
	MDOwnerSeed       string
}

type OtherConfig struct {
	EmailVerifyLink string
}
type Notify struct {
	ServiceName      string
	SlackWebhook     string
	SlackLogWebhook  string
	TelegramApiToken string
	TelegramChatId   int64
}

type EmailConfig struct {
	API_URL                 string
	API_KEY                 string
	Domain                  string
	AgentVerifyMailDayLimit string
}

type AwsSES struct {
	Region    string
	AccessId  string
	SecretKey string
	Sender    string
	CharSet   string
}

func NewConfig() *Config {
	configPath := "./"
	runPath, _ := os.Getwd()
	matchPathStatus := false
	pathArr := strings.Split(runPath, "/")
	for i := len(pathArr) - 1; i > 0; i-- {
		configPath += "../"
		if pathArr[i] == "cmd" || pathArr[i] == "test" || pathArr[i] == "migration" {
			matchPathStatus = true
			break
		}
	}
	if !matchPathStatus {
		configPath = "./"
	}
	configPath += "config"

	// subgame runtime types
	chainRuntimeTypes, err := ioutil.ReadFile(fmt.Sprintf("%s/subgame.json", configPath))
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.WatchConfig()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Gin: &GinConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
			Mode: viper.GetString("server.mode"),
		},
		DB: &DBConfig{
			Dialect:  viper.GetString("db.dialect"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.database"),
			Flag:     viper.GetString("db.flag"),
		},
		Jwt: &JwtConfig{
			Issuer:               viper.GetString("jwt.issuer"),
			Secret:               viper.GetString("jwt.secret"),
			Refresh_token_exp:    viper.GetInt("jwt.refresh_token_exp"),
			Access_token_exp_sec: viper.GetInt("jwt.access_token_exp_sec"),
		},
		Redis: &RedisConfig{
			Host:           viper.GetString("redis.host"),
			Port:           viper.GetInt("redis.port"),
			Database:       viper.GetInt("redis.database"),
			Auth:           viper.GetString("redis.auth"),
			Max_idle:       viper.GetInt("redis.max_idle"),
			Max_active:     viper.GetInt("redis.max_active"),
			Idle_timeout:   viper.GetInt("redis.idle_timeout"),
			Notify_active:  viper.GetInt("redis.notify_active"),
			Polling_active: viper.GetInt("redis.polling_active"),
		},
		SubGame: &SubGame{
			RPC:               viper.GetString("subgame.rpc"),
			SS58Prefix:        viper.GetString("subgame.ss58_prefix"),
			ChainRuntimeTypes: chainRuntimeTypes,
			EventKey:          viper.GetString("subgame.event_key"),
			MDOwnerPubkey:     viper.GetString("subgame.md_owner_pubkey"),
			MDOwnerSeed:       viper.GetString("subgame.md_owner_seed"),
		},
		Other: &OtherConfig{
			EmailVerifyLink: viper.GetString("other.email_verify_link"),
		},
		Notify: &Notify{
			ServiceName:      viper.GetString("notify.service_name"),
			SlackWebhook:     viper.GetString("notify.slack_webhook"),
			SlackLogWebhook:  viper.GetString("notify.slack_log_webhook"),
			TelegramApiToken: viper.GetString("notify.telegram_api_token"),
			TelegramChatId:   viper.GetInt64("notify.telegram_chat_id"),
		},
		Email: &EmailConfig{
			API_URL:                 viper.GetString("email.api_url"),
			API_KEY:                 viper.GetString("email.api_key"),
			Domain:                  viper.GetString("email.domain"),
			AgentVerifyMailDayLimit: viper.GetString("email.agent_verify_mail_day_limit"),
		},
		AwsSES: &AwsSES{
			Region:    viper.GetString("aws_ses.region"),
			AccessId:  viper.GetString("aws_ses.access_id"),
			SecretKey: viper.GetString("aws_ses.secret_key"),
			Sender:    viper.GetString("aws_ses.sender"),
			CharSet:   viper.GetString("aws_ses.char_set"),
		},
	}
}
