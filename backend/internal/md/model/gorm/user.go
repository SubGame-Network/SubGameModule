package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
)

type User struct {
	ID          uint64 `gorm:"auto_increment primary_key"`
	NickName    string `gorm:"type:varchar(255);"`
	Country     string `gorm:"type:varchar(255);"`
	Address     string `gorm:"type:varchar(255);"`
	Email       string `gorm:"type:varchar(255);"`
	CountryCode string `gorm:"type:varchar(255);"`
	Phone       string `gorm:"type:varchar(255);"`
	Nonce       string `gorm:"type:varchar(255);"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func UserModelToDomain(input *User) *domain.User {
	return &domain.User{
		ID:          input.ID,
		NickName:    input.NickName,
		Country:     input.Country,
		Address:     input.Address,
		Email:       input.Email,
		CountryCode: input.CountryCode,
		Phone:       input.Phone,
		Nonce:       input.Nonce,
		UpdatedAt:   input.UpdatedAt,
		CreatedAt:   input.CreatedAt,
	}
}

func UserDomainToModel(input *domain.User) *User {
	return &User{
		ID:          input.ID,
		NickName:    input.NickName,
		Country:     input.Country,
		Address:     input.Address,
		Email:       input.Email,
		CountryCode: input.CountryCode,
		Phone:       input.Phone,
		Nonce:       input.Nonce,
		UpdatedAt:   input.UpdatedAt,
		CreatedAt:   input.CreatedAt,
	}
}
