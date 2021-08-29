package event

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	GetEventByTeamID(ctx context.Context, teamID *uuid.UUID) (*model.Event, error)
	CreateEvent(ctx context.Context, payload *schema.CreateEventRequest) (*model.Event, error)
	GetEvent(ctx context.Context, ID *uuid.UUID) (*model.Event, error)
	UpdateEvent(ctx context.Context, payload *schema.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, ID *uuid.UUID) error
}

type svc struct {
	repo Repository
}

func (s *svc) GetEventByTeamID(ctx context.Context, teamID *uuid.UUID) (*model.Event, error) {
	return s.repo.GetEventByTeamID(ctx, teamID)
}

func (s *svc) CreateEvent(ctx context.Context, payload *schema.CreateEventRequest) (*model.Event, error) {
	return s.repo.CreateEvent(ctx, payload)
}

func (s *svc) GetEvent(ctx context.Context, ID *uuid.UUID) (*model.Event, error) {
	return s.repo.GetEvent(ctx, ID)
}

func (s *svc) UpdateEvent(ctx context.Context, payload *schema.UpdateEventRequest) error {
	return s.repo.UpdateEvent(ctx, payload)
}

func (s *svc) DeleteEvent(ctx context.Context, ID *uuid.UUID) error {
	return s.repo.DeleteEvent(ctx, ID)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
