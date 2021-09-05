package participation

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id *uuid.UUID) (*model.Participation, error)
	DeleteParticipation(ctx context.Context, p *model.Participation) error
	CreateParticipation(ctx context.Context, eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error)
	GetParticipationTeams(ctx context.Context, eventID *uuid.UUID) ([]model.Team, error)
	IsUserParticipatingInEvent(ctx context.Context, eventId, userId *uuid.UUID) (*int64, error)
	ParticipationByEventAndUser(ctx context.Context, eventId, userId *uuid.UUID) (*model.Participation, error)
}
