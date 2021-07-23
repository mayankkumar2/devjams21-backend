package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var db *gorm.DB

func dbConnect() (*gorm.DB, error) {
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		panic("error: DATABASE_URL environment URL not present or is empty in the environment")
	}
}

func DB() {
	var err error
	db, err = dbConnect()
	if err != nil {
		panic(err)
	}
	initializeServices()
}

func ConnectionHealth() error  {
	db, err := db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
