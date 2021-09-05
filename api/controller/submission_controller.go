package controller

import (
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func SaveSubmission(ctx *gin.Context) {
	payload := new(struct {
		ParticipationID *uuid.UUID `json:"participation_id"`
		Meta            model.JSON `json:"submission"`
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
		views.ErrorView(e.ErrRecordNotFound, ctx)
		return
	}

	_, err = db.TeamService.GetTeamMember(ctx, p.TeamID, usr.ID)
	if err == gorm.ErrRecordNotFound {
		views.ErrorView(e.ErrRecordNotFound, ctx)
		return
	} else if err != nil {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}

	event, err := db.EventService.GetEvent(ctx, p.EventID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	if time.Now().Unix() < event.Start.Unix() {
		views.ErrorView(e.ErrEventYetToStart, ctx)
		return
	} else if event.End.Unix() < time.Now().Unix() {
		views.ErrorView(e.ErrEventOver, ctx)
		return
	}

	err = db.SubmissionService.UpdateSubmission(ctx, p.SubmissionID, payload.Meta)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}
