package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type RecordResult uint8

var (
	Failed  RecordResult = 0
	Success RecordResult = 1
	Pending RecordResult = 2
)

type StakeRecord struct {
	ID               uint64
	UserId           uint64
	UserName         string // join user talbe
	ModuleId         uint64
	ProgramId        uint64
	StakeSGB         decimal.Decimal
	PeriodOfUseMonth int
	StartTime        time.Time
	EndTime          time.Time
	NFTHash          string
	Address          string
	TxHash           string
	BlockNum         int
	Fee              decimal.Decimal
	TxStatus         uint8
	Nonce            string
	DoneAt           *time.Time
	NotifyTime       *time.Time
	UpdatedAt        time.Time
	CreatedAt        time.Time
}
