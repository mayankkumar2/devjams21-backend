package submission

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
	FindById(ctx context.Context, id *uuid.UUID) (*model.Submission, error)
	UpdateSubmission(ctx context.Context, subId *uuid.UUID, meta model.JSON) error
}
