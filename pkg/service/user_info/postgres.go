package user_info

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
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

func (r *repo) CreateUserInfo(ctx context.Context, info *schema.CreateUserInfoRequest) (*model.UserInfo, error) {

	user_info := &model.UserInfo{
		UserID:          info.UserID,
		GithubURL:       info.GithubURL,
		LinkedinURL:     info.LinkedinURL,
		DiscordUsername: info.DiscordUsername,
	}

	return user_info, r.DB.WithContext(ctx).Create(&user_info).Error
}

func (r *repo) DeleteUserInfo(ctx context.Context, info *schema.DeleteUserInfoRequest) error {

	return r.DB.WithContext(ctx).Table("user_infos").Where("user_id = ?", info.UserID).Delete(info.UserID).Error

}

func (r *repo) UpdateUserInfo(ctx context.Context, info *schema.UpdateUserInfoRequest) error {

	return r.DB.WithContext(ctx).Table("user_infos").Where("user_id = ?", info.UserID).Updates(info).Error

}

func (r *repo) GetUserInfo(ctx context.Context, info *schema.GetUserInfoRequest) (*model.UserInfo, error) {

	user_info := new(model.UserInfo)
	err := r.DB.WithContext(ctx).Find(user_info, "user_id = ?", info.UserID).Error

	if err != nil {
		return nil, err
	}

	return user_info, err

}
