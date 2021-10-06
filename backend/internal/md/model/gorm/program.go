package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/shopspring/decimal"
)

type Program struct {
	ID          uint64          `gorm:"auto_increment primary_key"`
	PeriodOfUse int             `gorm:"type:varchar(255);"`
	Amount      decimal.Decimal `gorm:"type:decimal(30,8);"`
	CreatedAt   time.Time
}

func ProgramModelToDomain(input *Program) *domain.Program {
	return &domain.Program{
		ID:          input.ID,
		PeriodOfUse: input.PeriodOfUse,
		Amount:      input.Amount,
		CreatedAt:   input.CreatedAt,
	}
}

func ProgramDomainToModel(input *domain.Program) *Program {
	return &Program{
		ID:          input.ID,
		PeriodOfUse: input.PeriodOfUse,
		Amount:      input.Amount,
		CreatedAt:   input.CreatedAt,
	}
}
