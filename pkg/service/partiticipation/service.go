package participation

import (
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	FindByID(id *uuid.UUID) (*model.Participation, error)
	DeleteParticipation(p *model.Participation) error
	CreateParticipation(eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error)
}

type svc struct {
	repo Repository
}

func (s *svc) FindByID(id *uuid.UUID) (*model.Participation, error) {
	return s.repo.FindByID(id)
}

func (s *svc) DeleteParticipation(p *model.Participation) error {
	return s.repo.DeleteParticipation(p)
}

func (s *svc) CreateParticipation(eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error) {
	return s.repo.CreateParticipation(eventId, userID, teamName)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
