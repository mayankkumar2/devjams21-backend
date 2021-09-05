package submission

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	FindById(ctx context.Context, id *uuid.UUID) (*model.Submission, error)
	UpdateSubmission(ctx context.Context, subId *uuid.UUID, meta model.JSON) error
}

type svc struct {
	repo Repository
}

func (s *svc) FindById(ctx context.Context, id *uuid.UUID) (*model.Submission, error) {
	return s.repo.FindById(ctx, id)
}

func (s *svc) UpdateSubmission(ctx context.Context, subId *uuid.UUID, meta model.JSON) error {
	return s.repo.UpdateSubmission(ctx, subId, meta)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
