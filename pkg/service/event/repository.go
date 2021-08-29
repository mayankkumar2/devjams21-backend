package event

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetEventByTeamID(ctx context.Context, teamID *uuid.UUID) (*model.Event, error)
	CreateEvent(ctx context.Context, payload *schema.CreateEventRequest) (*model.Event, error)
	GetEvent(ctx context.Context, ID *uuid.UUID) (*model.Event, error)
	UpdateEvent(ctx context.Context, payload *schema.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, ID *uuid.UUID) error
}
