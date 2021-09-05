package messageBoard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON) error
}

type svc struct {
	repo Repository
}

func (s *svc) CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON) error {
	return s.repo.CreateMessage(ctx, usrId, message, meta)
}

func NewService(r Repository) Service {
	return &svc{
		repo: r,
	}
}


