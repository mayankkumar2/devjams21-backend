package messageBoard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON, exp time.Time) error
}

type svc struct {
	repo Repository
}

func (s *svc) CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON, exp time.Time) error {
	return s.repo.CreateMessage(ctx, usrId, message, meta, exp)
}

func NewService(r Repository) Service {
	return &svc{
		repo: r,
	}
}
