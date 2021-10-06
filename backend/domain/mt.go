package domain

import (
	"sync"

	"time"

	"github.com/shopspring/decimal"
)

var MuxWithdraw sync.Mutex

// 鎖住狀態
type MuxStatus bool

var MuxStatusLocked MuxStatus = false
var MuxStatusOpen MuxStatus = true
var MuxWithdrawStatus = MuxStatusOpen

type CoinType uint8

var (
	// SGB
	CoinTypeSGB CoinType = 1
)

type MDService interface {
	// 區塊監聽
	SubGameEventListener() error
	SubGameFixBlockGap() error

	SubGameEventStake(event *SubGameChainEvent, gas decimal.Decimal, blockTimestamp int) error

	// 監控區塊落後
	MonitorBlockSlow() error
	// 監控提領交易上鏈逾時
	StakeTimeoutNotify() error

	SubGameGetOwnerBalance() (decimal.Decimal, error)
}

type MDRepository interface {
	SubGameGetLastBlockLog() (*SubGameBlockLog, error)
	SubGameInsertUpdateBlockLog(input *SubGameBlockLog)
	SubGameDeleteOldBlock(maxBlockNum uint64) error
	SubGameGetBlockGapLog() ([]*SubGameBlockLog, error)

	// user
	InsertUser(input User) error
	UpdateUser(user *User) error
	GetAllUser() ([]*User, error)
	GetAllUserByPage(row, page int, account string, email string, address string) ([]*User, int64, error)
	GetUserByAccountOrEmailOrAddress(account string, email string, address string) (*User, error)
	GetUserByAccount(account string) (*User, error)
	GetUserById(id uint64) (*User, error)
	GetUserByAddress(address string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUserNonce(id uint64, nonce string) error

	EventGetStakeRecordByTxhash(txHash string) (*StakeRecord, error)
	EventUpdateStakeRecord(txHash string, input *StakeRecord) error

	StakeGetUserByAddress(address string) (*User, error)
	StakeInsertRecord(input StakeRecord) error
	StakeRecordByUserIdAndTime(userId uint64, from, to time.Time) ([]*StakeRecord, error)
	StakeAllRecords(from, to time.Time) ([]*StakeRecord, int64, error)
	StakeRecordByID(id uint64) (*StakeRecord, error)
	StakeUpdateRecordByID(id uint64, input StakeRecord) error
	StakeGetTimeout() ([]*StakeRecord, error)
	StakeUpdateNotifyTime(id uint64) error
	StakeAllRecordsByPage(userId uint64, row, page int, NftHash string, status *int, PeriodOfUseMonth int) ([]*StakeRecord, int64, error)
	StakeTotalStakedByUserId(userId uint64) (decimal.Decimal, error)
	StakeTotalWithrawnByUserId(userId uint64) (decimal.Decimal, error)

	// module
	InsertModule(input Module) error
	UpdateModule(module *Module) error
	GetAllModule() ([]*Module, error)
	GetModuleById(id uint64) (*Module, error)

	// program
	InsertProgram(input Program) error
	GetAllProgram() ([]*Program, error)

	// contact
	InsertContact(input Contact) error
	UpdateContact(module *Contact) error
	GetAllContact() ([]*Contact, error)
	GetContactById(id uint64) (*Contact, error)
}
