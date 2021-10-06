//+build wireinject

package provider

import (
	"log"
	"sync"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/database"
	"github.com/SubGame-Network/SubGameModuleService/internal/jwt"
	"github.com/SubGame-Network/SubGameModuleService/internal/redis"
	"github.com/google/wire"
	"gorm.io/gorm"

	adminRepo "github.com/SubGame-Network/SubGameModuleService/internal/adminAccount/repository/gorm"
	MDRepo "github.com/SubGame-Network/SubGameModuleService/internal/md/repository/gorm"
	MDSvc "github.com/SubGame-Network/SubGameModuleService/internal/md/service"
	notifySvc "github.com/SubGame-Network/SubGameModuleService/internal/notify/service"
	smtp "github.com/SubGame-Network/SubGameModuleService/internal/smtp/service"
)

var db *gorm.DB
var dbOnce sync.Once

func NewDB() (*gorm.DB, error) {
	var err error
	if db == nil {
		dbOnce.Do(func() {
			log.Println("connect db")
			db, err = database.DatabaseConnection(NewConfig().DB)
			if err != nil {
				return
			}
			log.Println("connect db success")
		})
	}
	return db, err
}

var cg *config.Config
var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		log.Println("read config")
		cg = config.NewConfig()
		log.Println("read config success")
	})
	return cg
}

func NewJwt() (domain.JwtUtil, error) {
	panic(wire.Build(jwt.NewJwt, NewConfig))
}

func NewRedis() (domain.GoRedis, error) {
	panic(wire.Build(redis.NewGoRedis, NewConfig))
}

func NewAdminRepo() (domain.AdminRepository, error) {
	panic(wire.Build(adminRepo.NewAdminRepository, NewDB))
}

func NewMDService() (domain.MDService, error) {
	panic(wire.Build(MDSvc.NewMDService, NewConfig, NewMDRepo, NewNotifyService, NewRedis))
}

func NewMDRepo() (domain.MDRepository, error) {
	panic(wire.Build(MDRepo.NewMDRepository, NewDB, NewConfig))
}

func NewNotifyService() (domain.NotifyService, error) {
	panic(wire.Build(notifySvc.NewNotifyServer, NewConfig))
}

func NewSmtpService() (smtp.SmtpService, error) {
	panic(wire.Build(smtp.NewSmtpService, NewConfig, NewRedis))
}
