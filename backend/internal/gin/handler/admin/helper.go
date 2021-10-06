package admin

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// InsertAdminLog 記錄後台操作記錄
func InsertAdminLog(AdminUUID uuid.UUID, AdminAccount, logType, beforeData, afterData string) {
	adminLogSqlData := domain.AdminLog{
		UUID:       AdminUUID,
		Account:    AdminAccount,
		LogType:    logType,
		BeforeData: beforeData,
		AfterData:  afterData,
	}
	adminRepo, err := provider.NewAdminRepo()
	if err != nil {
		zap.S().Warnw("InsertAdminLogError", "err", err)
	}
	err = adminRepo.CreateAdminLog(adminLogSqlData)
	if err != nil {
		zap.S().Warnw("InsertAdminLogError", "err", err)
	}
}
