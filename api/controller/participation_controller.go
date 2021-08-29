package controller

import (
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"gorm.io/gorm"
	"net/http"

	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTeamsController(ctx *gin.Context) {
	eventIdValue := ctx.Param("event_id")
	eventId, err := uuid.Parse(eventIdValue)
	if err != nil {
		views.ErrorView(e.ErrBadPayloadFormat, ctx)
		return
	}

	teams, err := db.ParticipationService.GetParticipationTeams(ctx, &eventId)
	if err != nil {
		views.ErrorView(err, ctx)
		sentry.CaptureException(err)
		return
	}
	for i := range teams {
		m := make([]model.TeamXUser, 0, 100)
		for j := range teams[i].TeamXUser {
			if teams[i].TeamXUser[j].User != nil {
				teams[i].TeamXUser[j].User.RegNo = ""
			}

			if teams[i].TeamXUser[j].IsAccepted {
				m = append(m, teams[i].TeamXUser[j])
			}
		}
		teams[i].TeamXUser = m
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

	p, err := db.ParticipationService.CreateParticipation(ctx, payload.EventID, usr.ID, GenerateTeamName(usr.Email))
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnableToCreateParticipation, ctx)
		return
	}

	views.DataView(ctx, http.StatusCreated, "created participation", p)
}

func DeleteParticipationController(ctx *gin.Context)  {
	payload := new(struct{
		ParticipationID *uuid.UUID `json:"participation_id"`
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

	p, err := db.ParticipationService.FindByID(ctx, payload.ParticipationID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(err, ctx)
			return
		} else {
			sentry.CaptureException(err)
			views.ErrorView(e.ErrUnexpected, ctx)
			return
		}
	}

	teamMember, err := db.TeamService.GetTeamMember(ctx, p.TeamID, usr.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrTeamNotFound, ctx)
			return
		} else {
			sentry.CaptureException(err)
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	if !teamMember.IsLeader {
		views.ErrorView(e.ErrUnauthorizedNotTeamLeader, ctx)
		return
	}
	err = db.ParticipationService.DeleteParticipation(ctx, p)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "deleted participation", p)
}