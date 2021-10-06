package domain

import (
	"time"
)

type User struct {
	ID          uint64
	NickName    string
	Country     string // 城市
	Address     string
	Email       string
	CountryCode string // 國碼
	Phone       string
	Nonce       string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
