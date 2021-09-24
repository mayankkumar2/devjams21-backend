package messageBoard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON, exp time.Time) error {
	x := make([]model.MessageBoard, 0, 100)
	for _, v := range usrId {
		x = append(x, model.MessageBoard{
			UserID:    v,
			Message:   message,
			Meta:      meta,
			ExpiresAt: exp,
		})
	}
	return r.DB.WithContext(ctx).Create(&x).Error
}
