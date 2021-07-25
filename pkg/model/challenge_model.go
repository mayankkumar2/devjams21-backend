package model

import "github.com/google/uuid"

type Challenge struct {
	BaseModel
	Meta map[string]interface{} `json:"meta"`
	EventID *uuid.UUID `json:"event_id"`
}
