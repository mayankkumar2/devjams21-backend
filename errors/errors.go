package errors

import "errors"

var (
	ErrMethodNotAllowed = errors.New("error: Method is not allowed")
	ErrInvalidToken     = errors.New("error: Invalid Authorization token")
	ErrUserExists       = errors.New("error: User already exists")
)