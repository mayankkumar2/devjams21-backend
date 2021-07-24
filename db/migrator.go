package db

import (
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func migrateModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Team{},
		&model.TeamXUser{},
	)
	if err != nil {
		logrus.Errorln(err)
	}
}
