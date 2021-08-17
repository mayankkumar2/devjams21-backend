package errors

import "errors"

var (
	ErrMethodNotAllowed          = errors.New("error: Method is not allowed")
	ErrInvalidToken              = errors.New("error: Invalid Authorization token")
	ErrUserExists                = errors.New("error: User already exists")
	ErrBadPayloadFormat          = errors.New("error: Invalid payload schema")
	ErrUserInvalidIDToken        = errors.New("error: Invalid id_token string")
	ErrUserCreateErr             = errors.New("error: Unexpected error while creating user")
	ErrUnexpected                = errors.New("error: unexpected error occurred")
	ErrUnauthorizedNotTeamMember = errors.New("error: Unauthorized to access team data")
	ErrUnauthorizedNotTeamLeader = errors.New("error: Unauthorized to manipulate resource")
	ErrTeamNotFound              = errors.New("error: Team not found")
	ErrTeamLeaderMandatory       = errors.New("error: Every team must have a team leader")
	ErrUserNotFound              = errors.New("error: User not found")
	ErrJoinRequestNotFound       = errors.New("error: Team join request not found")
	ErrRecordNotFound            = errors.New("error: Record not found")
)
