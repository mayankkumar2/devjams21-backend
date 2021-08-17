package schema

import "github.com/google/uuid"

type CreateChallengeRequest struct {
	Meta    map[string]interface{} `json:"meta"`
	EventID *uuid.UUID             `json:"event_id" gorm:"type:uuid"`
}

type GetChallengeRequest struct {
	ID *uuid.UUID `json:"id" gorm:"type:uuid"`
}

type UpdateChallengeRequest struct {
	ID      *uuid.UUID             `json:"id" gorm:"type:uuid"`
	Meta    map[string]interface{} `json:"meta" gorm:"type:json"`
	EventID *uuid.UUID             `json:"event_id" gorm:"type:uuid"`
}

type DeleteChallengeRequest struct {
	ID *uuid.UUID `json:"id" gorm:"type:uuid"`
}
