package domain

import (
	"time"
)

type Contact struct {
	ID        uint64
	UserId    uint64
	Type      string
	Contact   string
	CreatedAt time.Time
}
