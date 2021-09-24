package leaderboard

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

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) GetLeaderBoard(ctx context.Context) ([]schema.LeaderboardResponse, error) {
	m := make([]schema.LeaderboardResponse, 0, 100)
	return m, r.DB.WithContext(ctx).Table(" (SELECT sum(score) as scr, user_id FROM scores GROUP BY user_id) sc").
		Joins("JOIN users ON users.id = sc.user_id").
		Select("users.first_name as first_name, users.last_name as last_name,sc.scr as scr, users.photo_url as photo_url").
		Order("sc.scr DESC").
		Find(&m).Error
}

func (r *repo) CreateScore(ctx context.Context, id *uuid.UUID, score uint, msg string) error {
	return r.DB.WithContext(ctx).Create(&model.Score{
		Score:   score,
		UserID:  id,
		Message: msg,
	}).Error
}
