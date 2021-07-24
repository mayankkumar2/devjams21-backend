package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	ID        *uuid.UUID    `json:"id,omitempty" gorm:"type:uuid"`
	CreatedAt *time.Time    `json:"created_at,omitempty"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
	DeletedAt *sql.NullTime `json:"deleted_at,omitempty"`
}
