package leaderboard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/google/uuid"
)

type Service interface {
	CreateScore(ctx context.Context, id *uuid.UUID, score uint, msg string) error
	GetLeaderBoard(ctx context.Context) ([]schema.LeaderboardResponse,error)
}

type svc struct {
	repo Repository
}

func (s *svc) GetLeaderBoard(ctx context.Context) ([]schema.LeaderboardResponse, error) {
	return s.repo.GetLeaderBoard(ctx)
}

func (s *svc) CreateScore(ctx context.Context, id *uuid.UUID, score uint, msg string) error {
	return s.repo.CreateScore(ctx, id, score, msg)
}

func NewService(r Repository) Service {
	return &svc{
		repo: r,
	}
}


