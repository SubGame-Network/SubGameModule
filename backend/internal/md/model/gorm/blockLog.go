package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"gorm.io/gorm"
)

type SubgameBlockLog struct {
	gorm.Model
	Num        uint64 `gorm:"uniqueIndex:block_num; not null"`
	Done       bool   `gorm:"type:boolean; not null"`
	ErrorCount int    `gorm:"int(11);"`
}

func SubGameBlockLogModelToDomain(input *SubgameBlockLog) *domain.SubGameBlockLog {
	return &domain.SubGameBlockLog{
		Num:        input.Num,
		Done:       input.Done,
		ErrorCount: input.ErrorCount,
	}
}

func SubGameBlockLogDomainToModel(input *domain.SubGameBlockLog) *SubgameBlockLog {
	return &SubgameBlockLog{
		Num:        input.Num,
		Done:       input.Done,
		ErrorCount: input.ErrorCount,
	}
}
