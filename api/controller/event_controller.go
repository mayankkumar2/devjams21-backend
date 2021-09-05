package controller

import (
	"fmt"
	"github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func CreateEventController(ctx *gin.Context) {
	payload := new(schema.CreateEventRequest)
	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	e, err := db.EventService.CreateEvent(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusCreated, "success", e)
}

func UpdateEventController(ctx *gin.Context) {
	payload := new(schema.UpdateEventRequest)

	if err := ctx.ShouldBindJSON(payload); err != nil {
		sentry.CaptureException(err)
	}
	fmt.Println(payload)
	err := db.EventService.UpdateEvent(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}

func DeleteEventController(ctx *gin.Context) {
	payload := new(schema.DeleteEventRequest)
	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	err := db.EventService.DeleteEvent(ctx, payload.ID)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}

func GetEventController(ctx *gin.Context) {
	eventIdValue := ctx.Param("event_id")
	eventId, err := uuid.Parse(eventIdValue)
	if err != nil {
		views.ErrorView(errors.ErrBadPayloadFormat, ctx)
		return
	}

	event, err := db.EventService.GetEvent(ctx, &eventId)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	isReg := false
	userValue, exists := ctx.Get("user")
	if !exists {
		sentry.CaptureException(err)
		views.ErrorView(errors.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)
	c, err := db.ParticipationService.IsUserParticipatingInEvent(ctx, event.ID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(errors.ErrUnexpected, ctx)
		return
	}
	if *c >= 1 {
		isReg = true
	}
	p, err := db.ParticipationService.ParticipationByEventAndUser(ctx, event.ID, usr.ID)
	if err != gorm.ErrRecordNotFound && err != nil {
		sentry.CaptureException(err)
		views.ErrorView(errors.ErrUnexpected, ctx)
		return
	}

	teamSize := 0
	if p != nil {
		mxu, _ := db.TeamService.FetchTeamMembers(ctx, p.TeamID)
		teamSize = len(mxu)
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"event":         event,
		"is_registered": isReg,
		"participation": p,
		"team_size": teamSize,
	})
}

func GetAllEventsController(ctx *gin.Context) {
	events, err := db.EventService.GetAllEvent(ctx)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(errors.ErrUnexpected, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", events)
}
