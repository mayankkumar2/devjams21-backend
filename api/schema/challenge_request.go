package schema

import "github.com/google/uuid"

type CreateChallengeRequest struct {
	Meta    map[string]interface{} `json:"meta" gorm:"type:json"`
	EventID *uuid.UUID             `json:"event_id" binding:"uuid"`
}

type GetChallengeRequest struct {
	ID *uuid.UUID `json:"id" binding:"uuid"`
}

type UpdateChallengeRequest struct {
	ID      *uuid.UUID             `json:"id" binding:"uuid"`
	Meta    map[string]interface{} `json:"meta" gorm:"type:json"`
	EventID *uuid.UUID             `json:"event_id" binding:"uuid"`
}

type DeleteChallengeRequest struct {
	ID *uuid.UUID `json:"id" binding:"uuid"`
}
