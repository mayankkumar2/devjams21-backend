package messageBoard

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository{
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateMessage(ctx context.Context, usrId []*uuid.UUID, message string, meta model.JSON) error {
	x := make([]model.MessageBoard, 0, 100)
	for _, v := range usrId {
		x = append(x, model.MessageBoard{
			UserID:  v,
			Message: message,
			Meta:   meta,
		})
	}
	return r.DB.WithContext(ctx).Create(&x).Error
}