package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        *uuid.UUID    `json:"id,omitempty" gorm:"type:uuid"`
	CreatedAt *time.Time    `json:"created_at,omitempty"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}

func (u *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	id := uuid.New()
	u.ID = &id
	return nil
}
