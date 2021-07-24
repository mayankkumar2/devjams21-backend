package user

import (
	"context"
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

func (r *repo) CreateUser(ctx context.Context, record *auth.UserRecord, req *schema.CreateUserRequest) (*model.User, error) {
	db := r.DB.WithContext(ctx)
	usr := &model.User{
		Name:    record.DisplayName,
		UID:     record.UID,
		Email:   record.Email,
		RegNo:   req.Meta.RegNo,
		College: req.Meta.College,
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
	return usr, r.DB.WithContext(ctx).First(usr, "uid = ?", uid).Error
}

func (r *repo) UpdateAttributes(ctx context.Context, id *uuid.UUID, p map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}).Updates(p).Error
}
