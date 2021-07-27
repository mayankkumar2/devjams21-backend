package challenge

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
	CreateChallenge(ctx context.Context, payload *schema.CreateChallengeRequest) (*model.Challenge, error)
	GetChallenge(ctx context.Context, ID *uuid.UUID) (*model.Challenge, error)
	UpdateChallenge(ctx context.Context, challenge *model.Challenge, payload *schema.UpdateChallengeRequest) error
	DeleteChallenge(ctx context.Context, ID *uuid.UUID) error
}
