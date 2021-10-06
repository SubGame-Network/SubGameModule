package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Program struct {
	ID          uint64
	PeriodOfUse int
	Amount      decimal.Decimal
	CreatedAt   time.Time
}
