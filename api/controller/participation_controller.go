package controller

import (
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"net/http"

	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTeamsController(ctx *gin.Context) {
	payload := new(struct {
		EventID *uuid.UUID `json:"event_id"`
	})

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}
	teams, err := db.ParticipationService.GetParticipationTeams(ctx, payload.EventID)
	if err != nil {
		views.ErrorView(err, ctx)
		sentry.CaptureException(err)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"teams": teams,
	})
}

func GenerateTeamName(userName string) string {
	return userName+"'s Team"
}

func CreateParticipationController(ctx *gin.Context) {
	payload := new(struct{
		EventID *uuid.UUID `json:"event_id"`
	})
	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)

	c, err := db.ParticipationService.IsUserParticipatingInEvent(ctx, payload.EventID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}

	if *c >= 1 {
		views.ErrorView(e.ErrUserAlreadyRegisteredForEvent, ctx)
		return
	}

	p, err := db.ParticipationService.CreateParticipation(ctx, payload.EventID, usr.ID, GenerateTeamName(usr.Name))
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnableToCreateParticipation, ctx)
		return
	}

	views.DataView(ctx, http.StatusCreated, "created participation", p)
}

