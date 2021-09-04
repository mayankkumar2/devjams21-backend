package user_info

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
)

type Service interface {
	CreateUserInfo(ctx context.Context, info *schema.CreateUserInfoRequest) (*model.UserInfo, error)
	DeleteUserInfo(ctx context.Context, info *schema.DeleteUserInfoRequest) error
	UpdateUserInfo(ctx context.Context, info *schema.UpdateUserInfoRequest) error
	GetUserInfo(ctx context.Context, info *schema.GetUserInfoRequest) (*model.UserInfo, error)
}

type svc struct {
	repo Repository
}

func (s *svc) CreateUserInfo(ctx context.Context, info *schema.CreateUserInfoRequest) (*model.UserInfo, error) {
	return s.repo.CreateUserInfo(ctx, info)
}

func (s *svc) DeleteUserInfo(ctx context.Context, info *schema.DeleteUserInfoRequest) error {
	return s.repo.DeleteUserInfo(ctx, info)
}

func (s *svc) UpdateUserInfo(ctx context.Context, info *schema.UpdateUserInfoRequest) error {
	return s.repo.UpdateUserInfo(ctx, info)
}

func (s *svc) GetUserInfo(ctx context.Context, info *schema.GetUserInfoRequest) (*model.UserInfo, error) {
	return s.repo.GetUserInfo(ctx, info)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
