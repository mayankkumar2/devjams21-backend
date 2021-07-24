package db

import (
	"github.com/GDGVIT/devjams21-backend/pkg/service/team"
	user2 "github.com/GDGVIT/devjams21-backend/pkg/service/user"
	"gorm.io/gorm"
)

var (
	UserService user2.Service = nil
	TeamService team.Service  = nil
)

func initializeServices(db *gorm.DB) {
	UserService = user2.NewService(user2.NewUserRepo(db))
	TeamService = team.NewService(team.NewRepo(db))
}
