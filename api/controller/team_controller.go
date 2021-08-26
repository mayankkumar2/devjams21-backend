package controller

import (
	"fmt"
	"net/http"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateTeamCodeController(ctx *gin.Context) {
	payload := new(schema.FindTeamRequest)
	if err := ctx.BindQuery(payload); err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	teamID, _ := uuid.Parse(payload.ID)

	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)
	m, err := db.TeamService.GetTeamMember(ctx, &teamID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrUnauthorizedNotTeamMember, ctx)
		} else {
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}

	if !m.IsLeader {
		views.ErrorView(e.ErrUnauthorizedNotTeamLeader, ctx)
		return
	}
	t, err := db.TeamService.FindByID(ctx, &teamID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	if err := db.TeamService.UpdateTeamCode(ctx, t); err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{})
}

func GetTeamController(ctx *gin.Context) {
	payload := new(schema.FindTeamRequest)
	if err := ctx.BindQuery(payload); err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	teamID, _ := uuid.Parse(payload.ID)

	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)
	m, err := db.TeamService.GetTeamMember(ctx, &teamID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrUnauthorizedNotTeamMember, ctx)
		} else {
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}

	if !m.IsAccepted {
		views.ErrorView(e.ErrUnauthorizedNotTeamMember, ctx)
		fmt.Println("####################")
		return
	}

	t, err := db.TeamService.FindByID(ctx, &teamID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	members, err := db.TeamService.GetMembers(ctx, t.ID)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"team": gin.H{
			"team_details": t,
			"user_details": m,
			"members":      members,
		},
	})
}

func CreateTeamController(ctx *gin.Context) {
	payload := new(schema.CreateTeamRequest)
	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}
	userValue, exists := ctx.Get("user")
	if !exists {
		views.ErrorView(e.ErrUnexpected, ctx)
		return
	}
	usr := userValue.(*model.User)
	t, err := db.TeamService.CreateTeam(ctx, usr, payload.TeamName)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusCreated, "success", t)
}

func JoinTeamController(ctx *gin.Context) {
	payload := new(schema.JoinTeamRequest)
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

	t, err := db.TeamService.FindByJoinCode(ctx, payload.Code)
	if err != nil {
		sentry.CaptureException(err)
		views.ErrorView(err, ctx)
		return
	}

	err = db.TeamService.JoinTeam(ctx, t, usr)
	if err != nil {
		sentry.CaptureException(err)
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrTeamNotFound, ctx)
		} else {
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"team": t,
	})
}

func LeaveTeamController(ctx *gin.Context) {
	payload := new(schema.TeamIDRequest)
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
	m, err := db.TeamService.GetTeamMember(ctx, payload.ID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrUnauthorizedNotTeamMember, ctx)
		} else {
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	if m.IsLeader {
		views.ErrorView(e.ErrTeamLeaderMandatory, ctx)
		return
	}
	t, err := db.TeamService.FindByID(ctx, payload.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrTeamNotFound, ctx)
		} else {
			sentry.CaptureException(err)
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	err = db.TeamService.RemoveFromTeam(ctx, t, usr)
	if err != nil {
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"team": t,
	})
}

func RemoveMemberController(ctx *gin.Context) {
	payload := new(schema.TeamXMemberRequest)
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
	m, err := db.TeamService.GetTeamMember(ctx, payload.ID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrUnauthorizedNotTeamMember, ctx)
		} else {
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	if !m.IsLeader {
		views.ErrorView(e.ErrUnauthorizedNotTeamLeader, ctx)
		return
	}
	t, err := db.TeamService.FindByID(ctx, payload.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrTeamNotFound, ctx)
		} else {
			sentry.CaptureException(err)
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	u, err := db.UserService.FindByID(ctx, payload.ID)
	if err != nil {
		views.ErrorView(e.ErrUserNotFound, ctx)
		return
	}

	err = db.TeamService.RemoveFromTeam(ctx, t, u)
	if err != nil {
		views.ErrorView(err, ctx)
		return
	}
	views.DataView(ctx, http.StatusOK, "user removed", gin.H{})
}

func AcceptMemberRequestController(ctx *gin.Context) {
	payload := new(schema.TeamXMemberRequest)
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
	m, err := db.TeamService.GetTeamMember(ctx, payload.ID, usr.ID)
	if err != nil {
		sentry.CaptureException(err)
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrUnauthorizedNotTeamMember, ctx)
		} else {
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	if !m.IsLeader {
		views.ErrorView(e.ErrUnauthorizedNotTeamLeader, ctx)
		return
	}
	t, err := db.TeamService.FindByID(ctx, payload.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			views.ErrorView(e.ErrTeamNotFound, ctx)
		} else {
			sentry.CaptureException(err)
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	u, err := db.UserService.FindByID(ctx, payload.ID)
	if err != nil {
		views.ErrorView(e.ErrUserNotFound, ctx)
		return
	}
	err = db.TeamService.AcceptJoinRequest(ctx, t, u.ID)
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			views.ErrorView(e.ErrJoinRequestNotFound, ctx)
		} else {
			sentry.CaptureException(err)
			views.ErrorView(e.ErrUnexpected, ctx)
		}
		return
	}
	views.DataView(ctx, http.StatusOK, "user join request accepted", gin.H{})
}

func CheckTeamFullController(ctx *gin.Context) {
	payload := new(struct {
		TeamID  *uuid.UUID `json:"team_id"`
		EventID *uuid.UUID `json:"event_id"`
	})

	if err := ctx.BindJSON(payload); err != nil {
		sentry.CaptureException(err)
		return
	}

	event, err := db.EventService.GetEvent(ctx, payload.EventID)

	if err != nil {
		views.ErrorView(err, ctx)
		sentry.CaptureException(err)
		return
	}

	memberLimit := event.MemberLimit
	memberLowerLimit := event.MemberLowerLimit

	members, err := db.TeamService.GetMembers(ctx, payload.TeamID)

	if err != nil {
		views.ErrorView(err, ctx)
		sentry.CaptureException(err)
		return
	}

	memberCount := len(members)
	isFull := int(memberLimit) == memberCount
	isEmpty := int(memberLowerLimit) == memberCount

	views.DataView(ctx, http.StatusOK, "success", gin.H{
		"is_full":  isFull,
		"is_empty": isEmpty,
	})

}
