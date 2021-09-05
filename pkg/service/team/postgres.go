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
	return &t, r.DB.WithContext(ctx).Where("id = ?", id.String()).First(&t).Error
}

func (r *repo) FetchTeamMembers(ctx context.Context, teamId *uuid.UUID) ([]model.TeamXUser, error) {
	var tm = make([]model.TeamXUser, 0, 100)
	return tm, r.DB.WithContext(ctx).Where("team_id = ?", teamId).Joins("User").Find(&tm).Error
}

func (r *repo) FindByJoinCode(ctx context.Context, code string) (*model.Team, error) {
	t := new(model.Team)
	return t, r.DB.WithContext(ctx).
		First(t, "join_code = ?", code).Error
}

func (r *repo) JoinTeam(ctx context.Context, team *model.Team, usr *model.User) error {
	m := &model.TeamXUser{
		UserID:     usr.ID,
		TeamID:     team.ID,
		IsLeader:   false,
		IsAccepted: true,
	}
	return r.DB.WithContext(ctx).Create(m).Error
}

func (r *repo) RemoveFromTeam(ctx context.Context, team *model.Team, usr *model.User) error {
	return r.DB.WithContext(ctx).Where("user_id = ? AND team_id = ?", usr.ID, team.ID).
		Delete(&model.TeamXUser{}).Error
}

func (r *repo) AcceptJoinRequest(ctx context.Context, team *model.Team, userID *uuid.UUID) error {
	return r.DB.WithContext(ctx).Table("team_x_users").
		Where("team_id = ? AND user_id = ?", team.ID, userID).
		Update("IsAccepted", true).
		Error
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

func (r *repo) GetTeamMember(ctx context.Context, teamId *uuid.UUID, userId *uuid.UUID) (*model.TeamXUser, error) {
	var t = new(model.TeamXUser)
	err := r.DB.WithContext(ctx).
		First(t, "team_id = ? AND user_id = ?", teamId, userId).Error
	if err != nil {
		return nil, err
	}
	return t, err
}
func (r *repo) UpdateTeamName(ctx context.Context, teamId *uuid.UUID, teamName string) error {
	return r.DB.WithContext(ctx).Model(&model.Team{}).Where("id = ?", teamId).
		Update("TeamName", teamName).Error
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
