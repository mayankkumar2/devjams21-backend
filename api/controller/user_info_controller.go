package controller

import (
	"net/http"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func CreateUserInfoController(ctx *gin.Context) {
	payload := new(schema.CreateUserInfoRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	user_info, err := db.UserInfoService.CreateUserInfo(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
	}

	views.DataView(ctx, http.StatusOK, "success", user_info)
}

func DeleteUserInfoController(ctx *gin.Context) {
	payload := new(schema.DeleteUserInfoRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	err := db.UserInfoService.DeleteUserInfo(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}

func UpdateUserInfoController(ctx *gin.Context) {
	payload := new(schema.UpdateUserInfoRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	err := db.UserInfoService.UpdateUserInfo(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
	}

	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}

func GetUserInfoController(ctx *gin.Context) {
	payload := new(schema.GetUserInfoRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	user_info, err := db.UserInfoService.GetUserInfo(ctx, payload)

	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
	}

	views.DataView(ctx, http.StatusOK, "success", user_info)
}
