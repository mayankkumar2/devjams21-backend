package messageBoard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON, exp time.Time) error
	
}
