package domain

import (
	"time"
)

type Module struct {
	ID          uint64
	Name        string
	Depiction   string
	ReadmeMdUrl string
	Tags        string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
