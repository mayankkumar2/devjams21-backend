package db

import (
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func migrateModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Team{},
		&model.TeamXUser{},
		&model.Event{},
		&model.Challenge{},
		&model.Submission{},
		&model.Participation{},
	)
	if err != nil {
		sentry.CaptureException(err)
		logrus.Errorln(err)
	}
}
