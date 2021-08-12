package challenge

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	CreateChallenge(ctx context.Context, payload *schema.CreateChallengeRequest) (*model.Challenge, error)
	GetChallenge(ctx context.Context, EventID *uuid.UUID) (*model.Challenge, error)
	UpdateChallenge(ctx context.Context, challenge *model.Challenge, payload *schema.UpdateChallengeRequest) error
	DeleteChallenge(ctx context.Context, ID *uuid.UUID) error
}

type svc struct {
	repo Repository
}

func (s *svc) CreateChallenge(ctx context.Context, payload *schema.CreateChallengeRequest) (*model.Challenge, error) {
	return s.repo.CreateChallenge(ctx, payload)
}

func (s *svc) GetChallenge(ctx context.Context, ID *uuid.UUID) (*model.Challenge, error) {
	return s.repo.GetChallenge(ctx, ID)
}

func (s *svc) UpdateChallenge(ctx context.Context, challenge *model.Challenge, payload *schema.UpdateChallengeRequest) error {
	return s.repo.UpdateChallenge(ctx, challenge, payload)
}

func (s *svc) DeleteChallenge(ctx context.Context, ID *uuid.UUID) error {
	return s.repo.DeleteChallenge(ctx, ID)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
