package user

import (
	"context"
	"time"

	"firebase.google.com/go/auth"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindMessages(ctx context.Context, userID *uuid.UUID) ([]model.MessageBoard, error) {
	msg := make([]model.MessageBoard, 0, 100)
	return msg, r.DB.WithContext(ctx).Where("user_id = ? AND expires_at > ?", userID, time.Now()).Table("message_boards").
		Order("created_at DESC").
		Find(&msg).Error
}

func (r *repo) CreateUser(ctx context.Context, record *auth.UserRecord, req *schema.CreateUserRequest) (*model.User, error) {
	db := r.DB.WithContext(ctx)
	usr := &model.User{
		Name:           record.DisplayName,
		UID:            record.UID,
		Email:          record.Email,
		RegNo:          req.Meta.RegNo,
		College:        req.Meta.College,
		PhotoUrl:       record.PhotoURL,
		PhoneNumber:    req.Meta.PhoneNumber,
		Gender:         req.Meta.Gender,
		Degree:         req.Meta.Degree,
		Stream:         req.Meta.Stream,
		GraduationYear: req.Meta.GraduationYear,
		Age:            req.Meta.Age,
		Address:        req.Meta.Address,
		TShirtSize:     req.Meta.TShirtSize,
	}
	return usr, db.Create(&usr).Error
}

func (r *repo) FindByID(ctx context.Context, id *uuid.UUID) (*model.User, error) {
	usr := new(model.User)
	return usr, r.DB.WithContext(ctx).
		First(usr, "id = ?", id.String()).Error
}



func (r *repo) FindByUID(ctx context.Context, uid string) (*model.User, error) {
	var usr = new(model.User)
	return usr, r.DB.WithContext(ctx).
		First(usr, "uid = ?", uid).
		Error
}

func (r *repo) GetTeams(ctx context.Context, userID *uuid.UUID) ([]model.Team, error) {
	var t []model.Team
	err := r.DB.WithContext(ctx).Table("teams").
		Find(&t, "id IN (SELECT team_id FROM team_x_users WHERE user_id = ?)", userID).
		Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repo) IsLeader(ctx context.Context, userID *uuid.UUID, teamID *uuid.UUID) (bool, error) {
	teamXUser := new(model.TeamXUser)

	err := r.DB.WithContext(ctx).Table("team_x_users").
		First(teamXUser, "user_id = ? AND team_id = ?", userID, teamID).
		Error
	if err != nil {
		return false, err
	}
	return teamXUser.IsLeader, err
}

func (r *repo) UpdateAttributes(ctx context.Context, id *uuid.UUID, p map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}).Updates(p).Error
}

func (r *repo) MyParticipation(ctx context.Context, userId *uuid.UUID) ([]model.Participation, error) {

	m := make([]model.Participation, 0, 100)

	return m, r.DB.WithContext(ctx).
		Joins("Event").
		Joins("Team").
		Find(
			&m,
			"team_id IN (SELECT id FROM teams JOIN team_x_users txu ON teams.id=txu.team_id AND  user_id = ?)",
			userId).Error
}
