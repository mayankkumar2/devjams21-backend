package participation

import (
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
	FindByID(id *uuid.UUID) (*model.Participation, error)
	DeleteParticipation(p *model.Participation) error
	CreateParticipation(eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error)
}
