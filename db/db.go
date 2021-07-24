package db

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
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
		sentry.CaptureException(err)
		panic(err)
	}
	migrateModels(db)
	initializeServices(db)
}

func ConnectionHealth(ctx *gin.Context) error {
	db, err := db.DB()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	c, _ := context.WithTimeout(ctx, time.Minute)
	return db.PingContext(c)
}
