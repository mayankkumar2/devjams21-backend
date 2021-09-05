package model

import (
	"github.com/google/uuid"
	"time"
)

type TeamXUser struct {
	UserID     *uuid.UUID `json:"user_id" gorm:"primaryKey"`
	TeamID     *uuid.UUID `json:"team_id" gorm:"primaryKey"`
	IsLeader   bool       `json:"is_leader"`
	IsAccepted bool       `json:"is_accepted"`
	CreatedAt  *time.Time `json:"created_at"`
	//Team       *Team      `json:"team,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User *User `json:"user,omitempty"`
}
