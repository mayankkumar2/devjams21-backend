package challenge

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) CreateChallenge(ctx context.Context, payload *schema.CreateChallengeRequest) (*model.Challenge, error) {
	db := r.DB.WithContext(ctx)

	c := &model.Challenge{
		EventID: payload.EventID,
		Meta:    model.JSON(payload.Meta),
	}

	return c, db.Create(&c).Error
}

func (r *repo) GetChallenge(ctx context.Context, ID *uuid.UUID) (*model.Challenge, error) {
	c := new(model.Challenge)
	err := r.DB.WithContext(ctx).First(c, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *repo) UpdateChallenge(ctx context.Context, payload *schema.UpdateChallengeRequest) error {
	return r.DB.WithContext(ctx).Table("challenges").Where("id = ?", payload.ID).Omit("id, event_id").Updates(payload).Error
}

func (r *repo) DeleteChallenge(ctx context.Context, ID *uuid.UUID) error {
	return r.DB.WithContext(ctx).Where("id = ?", ID).Delete(&model.Challenge{}).Error
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
