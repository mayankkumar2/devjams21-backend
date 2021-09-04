package user_info

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
)

type Repository interface {
	CreateUserInfo(ctx context.Context, info *schema.CreateUserInfoRequest) (*model.UserInfo, error)
	DeleteUserInfo(ctx context.Context, info *schema.DeleteUserInfoRequest) error
	UpdateUserInfo(ctx context.Context, info *schema.UpdateUserInfoRequest) error
	GetUserInfo(ctx context.Context, info *schema.GetUserInfoRequest) (*model.UserInfo, error)
}
