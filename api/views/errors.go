package views

import (
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrView struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var (
	ErrMethodNotAllowed = errors.New("error: Method is not allowed")
	ErrInvalidToken     = errors.New("error: Invalid Authorization token")
	ErrUserExists       = errors.New("error: User already exists")
)

var ErrHTTPStatusMap = map[string]int{
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrInvalidToken.Error():     http.StatusBadRequest,
	ErrUserExists.Error():       http.StatusConflict,
}

func Wrap(err error, w http.ResponseWriter) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	// If error code is not found
	// like a default case
	if code == 0 {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)

	errView := ErrView{
		Message: msg,
		Status:  code,
	}
	log.WithFields(log.Fields{
		"message": msg,
		"code":    code,
	}).Error("Error occurred")

	json.NewEncoder(w).Encode(errView)
}
