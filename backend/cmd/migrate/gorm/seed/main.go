package main

import (
	"fmt"
	"log"

	"crypto/md5"
	"errors"

	"github.com/SubGame-Network/SubGameModuleService/config"
	admin "github.com/SubGame-Network/SubGameModuleService/internal/adminAccount/repository/gorm"
	mt "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/shopspring/decimal"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	config := config.NewConfig()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connection,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln(err)
	}

	// 建立data
	db.Exec("USE " + config.DB.Database)

	err = seedAdmin(db)
	if err != nil {
		log.Fatalln(err)
	}

	err = seedProgram(db)
	if err != nil {
		log.Fatalln(err)
	}

	err = seedModule(db)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")
}

func seedAdmin(db *gorm.DB) error {
	var err error
	var adminModel admin.Admin
	err = db.Where("account = ?", "admin").First(&adminModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.Create(&admin.Admin{
			UUID:     uuid.FromStringOrNil("955bc062-0d1e-4ef8-b3e4-5e3160e21986"),
			Account:  "admin",
			Password: fmt.Sprintf("%x", md5.Sum([]byte("1q2w3e4r"))),
		}).Error
	}
	return err
}

func seedProgram(db *gorm.DB) error {
	var m []*mt.Program
	err := db.Find(&m).Error

	if len(m) == 0 {
		err = db.Create(&mt.Program{
			ID:          1,
			PeriodOfUse: 6,
			Amount:      decimal.NewFromInt(10000),
		}).Error
		err = db.Create(&mt.Program{
			ID:          2,
			PeriodOfUse: 12,
			Amount:      decimal.NewFromInt(15000),
		}).Error
	}
	return err
}

func seedModule(db *gorm.DB) error {
	var m []*mt.Module
	err := db.Find(&m).Error

	if len(m) == 0 {
		err = db.Create(&mt.Module{
			ID:          1,
			Name:        "demo game1",
			Depiction:   "This is a dome pallet, which simulates a game module",
			ReadmeMdUrl: "https://raw.githubusercontent.com/polkadot-js/apps/master/README.md",
		}).Error
		err = db.Create(&mt.Module{
			ID:          2,
			Name:        "demo game2",
			Depiction:   "This is a dome pallet, which simulates a game module",
			ReadmeMdUrl: "https://raw.githubusercontent.com/polkadot-js/apps/master/README.md",
		}).Error
	}
	return err
}
