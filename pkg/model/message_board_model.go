package model

import (
	"github.com/google/uuid"
	"time"
)

type MessageBoard struct {
	UserID    *uuid.UUID `json:"user_id"`
	Message   string     `json:"message"`
	Meta      JSON       `json:"meta" gorm:"type:JSON"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt time.Time  `json:"expires_at"`
}
