package leaderboard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/google/uuid"
)

type Repository interface {
	CreateScore(ctx context.Context, id *uuid.UUID, score uint, msg string) error
	GetLeaderBoard(ctx context.Context) ([]schema.LeaderboardResponse, error)
}
