package db

import (
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

func dbConnect() (*gorm.DB, error) {
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
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
