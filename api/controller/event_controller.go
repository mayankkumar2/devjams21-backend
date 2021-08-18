package controller

import (
	"fmt"
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
	payload := new(schema.GetEventRequest)
	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	event, err := db.EventService.GetEvent(ctx, payload.ID)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", event)
}
