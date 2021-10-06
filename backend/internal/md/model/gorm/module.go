package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
)

type Module struct {
	ID          uint64 `gorm:"auto_increment primary_key"`
	Name        string `gorm:"type:varchar(255);"`
	Depiction   string `gorm:"type:string;"`
	ReadmeMdUrl string `gorm:"type:varchar(255);"`
	Tags        string `gorm:"type:varchar(255);"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func ModuleModelToDomain(input *Module) *domain.Module {
	return &domain.Module{
		ID:          input.ID,
		Name:        input.Name,
		Depiction:   input.Depiction,
		ReadmeMdUrl: input.ReadmeMdUrl,
		Tags:        input.Tags,
		UpdatedAt:   input.UpdatedAt,
		CreatedAt:   input.CreatedAt,
	}
}

func ModuleDomainToModel(input *domain.Module) *Module {
	return &Module{
		ID:          input.ID,
		Name:        input.Name,
		Depiction:   input.Depiction,
		ReadmeMdUrl: input.ReadmeMdUrl,
		Tags:        input.Tags,
		UpdatedAt:   input.UpdatedAt,
		CreatedAt:   input.CreatedAt,
	}
}
