package gorm

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Admin struct {
	ID        uint64    `gorm:"auto_increment primary_key"`
	UUID      uuid.UUID `gorm:"type:varchar(40); unique_index:uuid"`
	Account   string    `gorm:"type:varchar(255); not null"`
	Password  string    `gorm:"type:varchar(255); not null"`
	Status    bool      `gorm:"type:boolean; default:true"`
	Name      string    `gorm:"type:varchar(255);"`
	GroupUUID string    `gorm:"type:varchar(255);"`
	Email     string    `gorm:"type:varchar(255);"`
	UpdatedAt time.Time `gorm:"null"`
	CreatedAt time.Time `gorm:"null"`
}

type AdminLog struct {
	ID         uint64    `gorm:"auto_increment primary_key"`
	UUID       uuid.UUID `gorm:"type:varchar(40); not null"`
	Account    string    `gorm:"type:varchar(255); not null"`
	LogType    string    `gorm:"type:varchar(255); not null"`
	BeforeData string    `gorm:"type:mediumtext; index:,class:FULLTEXT;"`
	AfterData  string    `gorm:"type:mediumtext; index:,class:FULLTEXT;"`
	UpdatedAt  time.Time
	CreatedAt  time.Time
}