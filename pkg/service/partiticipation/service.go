package participation

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	// FindByID Fetches an participation object specified by id param
	FindByID(ctx context.Context, id *uuid.UUID) (*model.Participation, error)
	// DeleteParticipation cancels a participation
	DeleteParticipation(ctx context.Context, p *model.Participation) error
	// CreateParticipation creates a participation object
	CreateParticipation(ctx context.Context, eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error)
	// GetParticipationTeams gets array of teams participating in an event
	GetParticipationTeams(ctx context.Context, eventID *uuid.UUID) ([]model.Team, error)
	IsUserParticipatingInEvent(ctx context.Context, eventId, userId *uuid.UUID) (*int64, error)
}

type svc struct {
	repo Repository
}

func (s *svc) IsUserParticipatingInEvent(ctx context.Context, eventId, userId *uuid.UUID) (*int64, error) {
	return s.repo.IsUserParticipatingInEvent(ctx, eventId, userId)
}

func (s *svc) FindByID(ctx context.Context, id *uuid.UUID) (*model.Participation, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *svc) DeleteParticipation(ctx context.Context, p *model.Participation) error {
	return s.repo.DeleteParticipation(ctx, p)
}

func (s *svc) CreateParticipation(ctx context.Context, eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error) {
	return s.repo.CreateParticipation(ctx, eventId, userID, teamName)
}

func (s *svc) GetParticipationTeams(ctx context.Context, eventID *uuid.UUID) ([]model.Team, error) {
	return s.repo.GetParticipationTeams(ctx, eventID)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
