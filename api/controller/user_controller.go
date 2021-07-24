package controller

import (
	"github.com/GDGVIT/devjams21-backend/api/middleware"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/firebaseUtil"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserController(ctx *gin.Context) {
	var payload = new(schema.CreateUserRequest)

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	usrRec, err := firebaseUtil.GetUserDetail(ctx, payload.IdToken)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUserInvalidIDToken, ctx)
		return
	}

	usr, err := db.UserService.CreateUser(ctx, usrRec, payload)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUserExists, ctx)
		return
	}

	tok, exp, err := middleware.Token(usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusCreated, "success", gin.H{
		"user":   usr,
		"token":  tok,
		"expiry": exp,
	})
}

func UserProfileController(ctx *gin.Context) {
	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)

	views.DataView(ctx, http.StatusOK, "success", usr)
}

func UserProfileUpdateController(ctx *gin.Context) {
	var payload struct {
		Updates map[string]interface{} `json:"updates"`
	}
	if err := ctx.BindJSON(&payload); err != nil {
		sentry.CaptureException(err)
		return
	}
	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)

	allowedAttributes := []string{"name", "college", "reg_no"}

	updatesPayload := make(map[string]interface{})
	for _, v := range allowedAttributes {
		if k, ok := payload.Updates[v]; ok {
			updatesPayload[v] = k
		}
	}

	err := db.UserService.UpdateAttributes(ctx, usr.ID, updatesPayload)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	views.DataView(ctx, http.StatusOK, "success", nil)
}

func UserLoginController(ctx *gin.Context) {
	var payload = new(struct {
		IdToken string `json:"id_token"`
	})

	if err := ctx.BindJSON(payload); err != nil {
		return
	}

	usrRec, err := firebaseUtil.GetUserDetail(ctx, payload.IdToken)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUserInvalidIDToken, ctx)
		return
	}

	usr, err := db.UserService.FindByUID(ctx, usrRec.UID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	tok, exp, err := middleware.Token(usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"user":   usr,
		"token":  tok,
		"expiry": exp,
	})
}
