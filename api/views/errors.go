package views

import (
	"database/sql/driver"
	e "github.com/GDGVIT/devjams21-backend/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)


var ErrHTTPStatusMap = map[string]int{
	e.ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	e.ErrInvalidToken.Error():     http.StatusBadRequest,
	e.ErrUserExists.Error():       http.StatusConflict,
	driver.ErrBadConn.Error(): http.StatusServiceUnavailable,
	gorm.ErrInvalidDB.Error(): http.StatusServiceUnavailable,
}

func ErrorView(err error, c *gin.Context) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	if code == 0 {
		code = http.StatusInternalServerError
	}

	log.WithFields(log.Fields{
		"message": msg,
		"code":    code,
	}).Error("Error occurred")

	c.JSON(code, gin.H{
		"code": code,
		"error": true,
		"message": msg,
	})
}

