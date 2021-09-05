package submission

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindById(ctx context.Context, id *uuid.UUID) (*model.Submission, error) {
	sub := new(model.Submission)
	return sub, r.DB.WithContext(ctx).Model(&model.Submission{}).Where("id = ?", id).First(sub).Error
}

func (r *repo) UpdateSubmission(ctx context.Context, subId *uuid.UUID, meta model.JSON) error {
	return r.DB.WithContext(ctx).
		Table("submissions").
		Where("id = ?", subId).
		Update("meta", meta).
		Error
}
