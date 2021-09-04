package db

import (
	"github.com/GDGVIT/devjams21-backend/pkg/service/challenge"
	"github.com/GDGVIT/devjams21-backend/pkg/service/event"
	participation "github.com/GDGVIT/devjams21-backend/pkg/service/partiticipation"
	"github.com/GDGVIT/devjams21-backend/pkg/service/submission"
	"github.com/GDGVIT/devjams21-backend/pkg/service/team"
	user2 "github.com/GDGVIT/devjams21-backend/pkg/service/user"
	"github.com/GDGVIT/devjams21-backend/pkg/service/user_info"
	"gorm.io/gorm"
)

var (
	UserService          user2.Service         = nil
	TeamService          team.Service          = nil
	ChallengeService     challenge.Service     = nil
	EventService         event.Service         = nil
	SubmissionService    submission.Service    = nil
	ParticipationService participation.Service = nil
	UserInfoService      user_info.Service     = nil
)

func initializeServices(db *gorm.DB) {
	UserService = user2.NewService(user2.NewUserRepo(db))
	TeamService = team.NewService(team.NewRepo(db))
	ChallengeService = challenge.NewService(challenge.NewRepo(db))
	EventService = event.NewService(event.NewRepo(db))
	SubmissionService = submission.NewService(submission.NewRepo(db))
	ParticipationService = participation.NewService(participation.NewRepo(db))
	UserInfoService = user_info.NewService(user_info.NewRepo(db))
}
