package model

import "github.com/google/uuid"

type Challenge struct {
	BaseModel
	Meta    JSON       `json:"meta" gorm:"type:json"`
	EventID *uuid.UUID `json:"event_id"`
}
