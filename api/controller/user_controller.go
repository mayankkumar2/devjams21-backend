package controller

import (
	"github.com/GDGVIT/devjams21-backend/api/middleware"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/firebaseUtil"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserController(ctx *gin.Context) {
	var payload = new(schema.CreateUserRequest)


	if err := ctx.BindJSON(payload); err != nil {
		return
	}

	usrRec, err := firebaseUtil.GetUserDetail(ctx, payload.IdToken)
	if err != nil {
		views.ErrorView(e.ErrUserInvalidIDToken, ctx)
		return
	}

	usr, err := db.UserService.CreateUser(ctx, usrRec, payload)
	if err != nil {
		views.ErrorView(e.ErrUserExists, ctx)
		return
	}

	tok, exp, err := middleware.Token(usr.ID)
	if err != nil {
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"user": usr,
		"token": tok,
		"expiry": exp,
	})
}

func UserProfileController(ctx *gin.Context)  {
	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)

	views.DataView(ctx, http.StatusOK, "success", usr)
}