package service

import (
	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
)

type NotifyService struct {
	Config *config.Config
}

func NewNotifyServer(config *config.Config) (domain.NotifyService, error) {
	return &NotifyService{
		Config: config,
	}, nil
}
