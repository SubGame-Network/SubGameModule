package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
)

type Contact struct {
	ID        uint64 `gorm:"auto_increment primary_key"`
	UserId    uint64 `gorm:"type:int(11);"`
	Type      string `gorm:"type:varchar(255);"`
	Contact   string `gorm:"type:varchar(255);"`
	CreatedAt time.Time
}

func ContactModelToDomain(input *Contact) *domain.Contact {
	return &domain.Contact{
		ID:        input.ID,
		UserId:    input.UserId,
		Type:      input.Type,
		Contact:   input.Contact,
		CreatedAt: input.CreatedAt,
	}
}

func ContactDomainToModel(input *domain.Contact) *Contact {
	return &Contact{
		ID:        input.ID,
		UserId:    input.UserId,
		Type:      input.Type,
		Contact:   input.Contact,
		CreatedAt: input.CreatedAt,
	}
}
