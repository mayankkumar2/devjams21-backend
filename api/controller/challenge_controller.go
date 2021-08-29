package controller

import (
	"net/http"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func CreateChallengeController(ctx *gin.Context) {
	payload := new(schema.CreateChallengeRequest)
	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	challenge, err := db.ChallengeService.CreateChallenge(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusCreated, "success", challenge)
}

func GetChallengeController(ctx *gin.Context) {
	payload := new(schema.GetChallengeRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	challenge, err := db.ChallengeService.GetChallenge(ctx, payload.ID)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", challenge)

}

func UpdateChallengeController(ctx *gin.Context) {
	payload := new(schema.UpdateChallengeRequest)

	if err := ctx.ShouldBindJSON(payload); err != nil {
		sentry.CaptureException(err)
	}

	err := db.ChallengeService.UpdateChallenge(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})

}

func DeleteChallengeController(ctx *gin.Context) {
	payload := new(schema.DeleteChallengeRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	err := db.ChallengeService.DeleteChallenge(ctx, payload.ID)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})

}
