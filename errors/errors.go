package errors

import "errors"

var (
	ErrMethodNotAllowed   = errors.New("error: Method is not allowed")
	ErrInvalidToken       = errors.New("error: Invalid Authorization token")
	ErrUserExists         = errors.New("error: User already exists")
	ErrBadPayloadFormat   = errors.New("error: Invalid payload schema")
	ErrUserInvalidIDToken = errors.New("error: Invalid id_token string")
	ErrUserCreateErr      = errors.New("error: Unexpected error while creating user")
	ErrUnexpected		= errors.New("error: unexpected error occurred")
)