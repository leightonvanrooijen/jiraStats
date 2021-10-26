package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type JiraDB struct {
	db *gorm.DB
}

// Opens connection to mysql db
func ConnectDB() *JiraDB {
	db, err := gorm.Open(mysql.Open(
		"leighton:123456@tcp(127.0.0.1:3307)/jira?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Sprint{}, &Issue{})

	return &JiraDB{db: db}
}
