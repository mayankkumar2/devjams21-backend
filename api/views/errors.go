package views

import (
	"database/sql/driver"
	"net/http"

	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ErrHTTPStatusMap = map[error]int{
	e.ErrMethodNotAllowed:          http.StatusMethodNotAllowed,
	e.ErrInvalidToken:              http.StatusBadRequest,
	e.ErrUserExists:                http.StatusConflict,
	driver.ErrBadConn:              http.StatusServiceUnavailable,
	gorm.ErrInvalidDB:              http.StatusServiceUnavailable,
	e.ErrBadPayloadFormat:          http.StatusUnprocessableEntity,
	e.ErrUserInvalidIDToken:        http.StatusUnauthorized,
	e.ErrUserCreateErr:             http.StatusServiceUnavailable,
	gorm.ErrRecordNotFound:         http.StatusNotFound,
	e.ErrUnauthorizedNotTeamMember: http.StatusUnauthorized,
	e.ErrUnauthorizedNotTeamLeader: http.StatusUnauthorized,
}

func ErrorView(err error, c *gin.Context) {
	msg := err.Error()
	code := ErrHTTPStatusMap[err]

	if code == 0 {
		code = http.StatusInternalServerError
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
