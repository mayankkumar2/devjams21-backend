package team

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/GDGVIT/devjams21-backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) FindByID(ctx context.Context, id *uuid.UUID) (*model.Team, error) {
	var t model.Team
	return &t, r.DB.WithContext(ctx).Where("id = ?", id.String()).Find(&t).Error
}

func (r *repo) UpdateTeamCode(ctx context.Context, team *model.Team) error {
	err := r.DB.WithContext(ctx).
		Table("teams").
		Where("id = ?", team.ID).
		Update("join_code", util.RandStringRunes(16)).
		Error
	for err != gorm.ErrRecordNotFound && err != nil {
		err = r.DB.WithContext(ctx).
			Table("teams").
			Where("id = ?", team.ID).
			Update("join_code", util.RandStringRunes(16)).
			Error
	}
	return err
}

func (r *repo) GetMembers(ctx context.Context, id *uuid.UUID) ([]model.TeamXUser, error) {
	var t []model.TeamXUser
	err := r.DB.WithContext(ctx).Joins("User").Find(&t, "team_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *repo) GetTeamMember(ctx context.Context, teamId, userId *uuid.UUID) (*model.TeamXUser, error) {
	var t = new(model.TeamXUser)
	err := r.DB.WithContext(ctx).
		First(t, "team_id = ? AND user_id = ?", teamId, userId).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *repo) CreateTeam(ctx context.Context, usr *model.User, teamName string) (*model.Team, error) {
	t := new(model.Team)
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		t.TeamName = teamName
		t.JoinCode = util.RandStringRunes(16)
		err := tx.Create(t).Error
		if err != nil {
			return err
		}
		teamMember := &model.TeamXUser{
			UserID:     usr.ID,
			TeamID:     t.ID,
			IsLeader:   true,
			IsAccepted: true,
		}
		return tx.Create(teamMember).Error
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
