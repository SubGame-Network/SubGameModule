package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"gorm.io/gorm"
)

type MDRepository struct {
	db     *gorm.DB
	redis  domain.GoRedis
	config *config.Config
}

func NewMDRepository(db *gorm.DB, config *config.Config) domain.MDRepository {
	return &MDRepository{
		db:     db,
		config: config,
	}
}
