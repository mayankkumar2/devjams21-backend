package model

import "github.com/google/uuid"

type Score struct {
	Score   uint       `json:"score"`
	UserID  *uuid.UUID `json:"user_id"`
	User    *User      `json:"user,omitempty"`
	Message string     `json:"message"`
}
