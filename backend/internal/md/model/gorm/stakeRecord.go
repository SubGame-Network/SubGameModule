package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/shopspring/decimal"
)

type StakeRecord struct {
	ID               uint64          `gorm:"auto_increment primary_key"`
	UserId           uint64          `gorm:"type:int(11);"`
	ModuleId         uint64          `gorm:"type:int(11);"`
	ProgramId        uint64          `gorm:"type:int(11);"`
	StakeSGB         decimal.Decimal `gorm:"type:decimal(30,8);"`
	PeriodOfUseMonth int             `gorm:"type:int(11);"`
	StartTime        time.Time
	EndTime          time.Time
	NFTHash          string          `gorm:"type:varchar(255); uniqueindex;"`
	Address          string          `gorm:"type:varchar(255);"`
	TxHash           string          `gorm:"type:varchar(255); uniqueindex;"`
	BlockNum         int             `gorm:"type:int(11);"`
	Fee              decimal.Decimal `gorm:"type:decimal(30,8);"`
	TxStatus         uint8           `gorm:"type:int(11); index;"`
	Nonce            string          `gorm:"type:varchar(255);"`
	DoneAt           *time.Time
	NotifyTime       *time.Time
	UpdatedAt        time.Time
	CreatedAt        time.Time
}

func StakeRecordModelToDomain(input *StakeRecord) *domain.StakeRecord {
	return &domain.StakeRecord{
		ID:               input.ID,
		UserId:           input.UserId,
		ModuleId:         input.ModuleId,
		ProgramId:        input.ProgramId,
		StakeSGB:         input.StakeSGB,
		PeriodOfUseMonth: input.PeriodOfUseMonth,
		StartTime:        input.StartTime,
		EndTime:          input.EndTime,
		NFTHash:          input.NFTHash,
		Address:          input.Address,
		TxHash:           input.TxHash,
		BlockNum:         input.BlockNum,
		Fee:              input.Fee,
		TxStatus:         input.TxStatus,
		Nonce:            input.Nonce,
		DoneAt:           input.DoneAt,
		NotifyTime:       input.NotifyTime,
		UpdatedAt:        input.UpdatedAt,
		CreatedAt:        input.CreatedAt,
	}
}

func StakeRecordDomainToModel(input *domain.StakeRecord) *StakeRecord {
	return &StakeRecord{
		ID:               input.ID,
		UserId:           input.UserId,
		ModuleId:         input.ModuleId,
		ProgramId:        input.ProgramId,
		StakeSGB:         input.StakeSGB,
		PeriodOfUseMonth: input.PeriodOfUseMonth,
		StartTime:        input.StartTime,
		EndTime:          input.EndTime,
		NFTHash:          input.NFTHash,
		Address:          input.Address,
		TxHash:           input.TxHash,
		BlockNum:         input.BlockNum,
		Fee:              input.Fee,
		TxStatus:         input.TxStatus,
		Nonce:            input.Nonce,
		DoneAt:           input.DoneAt,
		NotifyTime:       input.NotifyTime,
		UpdatedAt:        input.UpdatedAt,
		CreatedAt:        input.CreatedAt,
	}
}
