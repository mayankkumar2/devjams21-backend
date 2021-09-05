package model

import "github.com/google/uuid"

type MessageBoard struct {
	UserID *uuid.UUID `json:"user_id"`
	Message string `json:"message"`
	Meta  JSON `json:"meta" gorm:"type:JSON"`
}
