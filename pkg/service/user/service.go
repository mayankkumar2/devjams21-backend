package user

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
)

type Service interface {
	CreateUser(ctx context.Context, record *auth.UserRecord, req *schema.CreateUserRequest) (*model.User, error)
}

type svc struct {
	repo Repository
}

func (s *svc) CreateUser(ctx context.Context, record *auth.UserRecord, req *schema.CreateUserRequest) (*model.User, error) {
	return s.repo.CreateUser(ctx, record, req)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
