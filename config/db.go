package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func DB() (*gorm.DB, error) {

	if db != nil {
		return db, nil
	}

	mode := gin.Mode()

	config := &gorm.Config{}

	if mode == "debug" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	dsn := "host=localhost user=jannik password=gorm dbname=ktor port=5432 sslmode=disable TimeZone=Europe/Berlin"

	dbI, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), config)
	db = dbI

	return db, err

}
