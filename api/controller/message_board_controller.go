package controller

import (
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func SendMessageToTeam(ctx *gin.Context)  {
	payload := new(struct{
		TeamId *uuid.UUID `json:"team_id"`
		Message string `json:"message"`
		Meta model.JSON `json:"meta"`
		ExpiresAt time.Time `json:"expires_at"`
	})

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	txu, err := db.TeamService.FetchTeamMembers(ctx, payload.TeamId)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	ids := make([]*uuid.UUID, 0, 100)
	for _, v := range txu {
		ids = append(ids, v.UserID)
	}

	err = db.MessageBoardService.CreateMessage(ctx, ids, payload.Message, payload.Meta, payload.ExpiresAt)
	if err != nil {
		views.ErrorView(err, nil)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}


func SendMessageToOne(ctx *gin.Context)  {
	payload := new(struct{
		UserId *uuid.UUID `json:"user_id"`
		Message string `json:"message"`
		Meta model.JSON `json:"meta"`
		ExpiresAt time.Time `json:"expires_at"`
	})

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	_, err := db.UserService.FindByID(ctx, payload.UserId)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	ids := []*uuid.UUID{
		payload.UserId,
	}

	err = db.MessageBoardService.CreateMessage(ctx, ids, payload.Message, payload.Meta, payload.ExpiresAt)
	if err != nil {
		views.ErrorView(err, nil)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}
