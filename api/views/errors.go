package views

import (
	"database/sql/driver"
	"github.com/getsentry/sentry-go"
	"net/http"

	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ErrHTTPStatusMap = map[error]int{
	e.ErrMethodNotAllowed:              http.StatusMethodNotAllowed,
	e.ErrInvalidToken:                  http.StatusBadRequest,
	e.ErrUserExists:                    http.StatusConflict,
	driver.ErrBadConn:                  http.StatusServiceUnavailable,
	gorm.ErrInvalidDB:                  http.StatusServiceUnavailable,
	e.ErrBadPayloadFormat:              http.StatusUnprocessableEntity,
	e.ErrUserInvalidIDToken:            http.StatusUnauthorized,
	e.ErrUserCreateErr:                 http.StatusServiceUnavailable,
	gorm.ErrRecordNotFound:             http.StatusNotFound,
	e.ErrUnauthorizedNotTeamMember:     http.StatusUnauthorized,
	e.ErrUnauthorizedNotTeamLeader:     http.StatusUnauthorized,
	e.ErrUnableToCreateParticipation:   http.StatusInternalServerError,
	e.ErrUserAlreadyRegisteredForEvent: http.StatusConflict,
	e.ErrTeamAtCapacity:                http.StatusConflict,
	e.ErrEventRSVPExpired:              http.StatusForbidden,
	e.ErrEventOver:                     http.StatusForbidden,
	e.ErrEventYetToStart:               http.StatusForbidden,
	e.ErrNoChallengeInEvent:            http.StatusConflict,
	e.ErrAgreeTermsCondition:           http.StatusUnauthorized,
	e.ErrAlreadyExists:  http.StatusConflict,
}

func ErrorView(err error, c *gin.Context) {
	msg := err.Error()
	code, exist := ErrHTTPStatusMap[err]

	if !exist {
		code = http.StatusInternalServerError
		sentry.CaptureException(err)
		msg = "unexpected error"
	}

	log.WithFields(log.Fields{
		"message": msg,
		"code":    code,
	}).Error("Error occurred")

	c.JSON(code, gin.H{
		"code":    code,
		"error":   true,
		"message": msg,
	})
}
