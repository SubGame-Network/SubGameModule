package main

import (
	"time"

	"fmt"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/route"
	"github.com/SubGame-Network/SubGameModuleService/internal/logger"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"go.uber.org/zap"
)

func main() {
	// 背景工作
	BackgroundWorker()

	config := provider.NewConfig()
	router := route.SetupRouter(config)
	router.Run(config.Gin.Host + ":" + config.Gin.Port)
}

func BackgroundWorker() {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	MDSvc, err := provider.NewMDService()
	if err != nil {
		zap.S().Error(err)
	}
	if MDSvc == nil {
		return
	}

	SubGameEventListener(MDSvc)
	SubGameBlockListener(MDSvc)
	StakeTimeoutNotify(MDSvc)
}

func SubGameEventListener(MDSvc domain.MDService) {
	go func() {
		// 監聽發生錯誤就12秒後重新監聽
		for ; true; <-time.Tick(12 * time.Second) {
			err := MDSvc.SubGameEventListener()
			if err != nil {
				zap.S().Errorw("SubGameEventListenerError", err)

				// 重新建立ws連線
				MDSvc, _ = provider.NewMDService()
			}
		}
	}()

	// 修復漏聽的區塊
	go func() {
		fmt.Println("=== SubGame Event Fix Start ===")
		for ; true; <-time.Tick(60 * time.Second) {
			err := MDSvc.SubGameFixBlockGap()
			if err != nil {
				// 重新建立ws連線
				MDSvc, _ = provider.NewMDService()
			}
		}
	}()
}

func SubGameBlockListener(MDSvc domain.MDService) {
	go func() {
		fmt.Println("=== 監聽區塊落後 Start ===")
		for ; true; <-time.Tick(30 * time.Minute) {
			err := MDSvc.MonitorBlockSlow()
			if err != nil {
				zap.S().Errorw("SubGameBlockListenerError", "err", err)
			}
		}
	}()
}

func StakeTimeoutNotify(MDSvc domain.MDService) {
	go func() {
		fmt.Println("=== 監聽提領交易上鏈逾時 Start ===")
		for ; true; <-time.Tick(10 * time.Minute) {
			err := MDSvc.StakeTimeoutNotify()
			if err != nil {
				zap.S().Errorw("NotifyStakePendingError", "err", err)
			}
		}
	}()
}
