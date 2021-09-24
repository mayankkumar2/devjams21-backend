package db

import (
	"github.com/GDGVIT/devjams21-backend/pkg/service/challenge"
	"github.com/GDGVIT/devjams21-backend/pkg/service/event"
	"github.com/GDGVIT/devjams21-backend/pkg/service/leaderboard"
	"github.com/GDGVIT/devjams21-backend/pkg/service/messageBoard"
	participation "github.com/GDGVIT/devjams21-backend/pkg/service/partiticipation"
	"github.com/GDGVIT/devjams21-backend/pkg/service/submission"
	"github.com/GDGVIT/devjams21-backend/pkg/service/team"
	user2 "github.com/GDGVIT/devjams21-backend/pkg/service/user"
	"gorm.io/gorm"
)

var (
	UserService          user2.Service         = nil
	TeamService          team.Service          = nil
	ChallengeService     challenge.Service     = nil
	EventService         event.Service         = nil
	SubmissionService    submission.Service    = nil
	ParticipationService participation.Service = nil
	MessageBoardService  messageBoard.Service  = nil
	LeaderboardService   leaderboard.Service   = nil
)

func initializeServices(db *gorm.DB) {
	UserService = user2.NewService(user2.NewUserRepo(db))
	TeamService = team.NewService(team.NewRepo(db))
	ChallengeService = challenge.NewService(challenge.NewRepo(db))
	EventService = event.NewService(event.NewRepo(db))
	SubmissionService = submission.NewService(submission.NewRepo(db))
	ParticipationService = participation.NewService(participation.NewRepo(db))
	MessageBoardService = messageBoard.NewService(messageBoard.NewRepo(db))
	LeaderboardService = leaderboard.NewService(leaderboard.NewRepo(db))
}
