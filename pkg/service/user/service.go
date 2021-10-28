package user

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	FetchNetworkProfileByID(ctx context.Context, id *uuid.UUID) (*model.User, error)
	FindMessages(ctx context.Context, userID *uuid.UUID) ([]model.MessageBoard, error)
	MyParticipation(ctx context.Context, userId *uuid.UUID) ([]model.Participation, error)
	CreateUser(ctx context.Context, record *auth.UserRecord, req *schema.CreateUserRequest) (*model.User, error)
	FindByID(ctx context.Context, id *uuid.UUID) (*model.User, error)
	UpdateAttributes(ctx context.Context, id *uuid.UUID, p map[string]interface{}) error
	FindByUID(ctx context.Context, uid string) (*model.User, error)
	GetTeams(ctx context.Context, userID *uuid.UUID) ([]model.Team, error)
	IsLeader(ctx context.Context, userID *uuid.UUID, teamID *uuid.UUID) (bool, error)
	UpdateSocialAttributes(ctx context.Context, id *uuid.UUID, p map[string]interface{}) error
	NetworkWithPeers(ctx context.Context, id *uuid.UUID) ([]model.User,error)
}

type svc struct {
	repo Repository
}

func (s *svc) FetchNetworkProfileByID(ctx context.Context, id *uuid.UUID) (*model.User, error) {
	return s.repo.FetchNetworkProfileByID(ctx, id)
}

func (s *svc) NetworkWithPeers(ctx context.Context, id *uuid.UUID) ([]model.User, error) {
	return s.repo.NetworkWithPeers(ctx, id)
}

func (s *svc) UpdateSocialAttributes(ctx context.Context, id *uuid.UUID, p map[string]interface{}) error {
	return s.repo.UpdateSocialAttributes(ctx, id, p)
}

func (s *svc) FindMessages(ctx context.Context, userID *uuid.UUID) ([]model.MessageBoard, error) {
	return s.repo.FindMessages(ctx, userID)
}

func (s *svc) MyParticipation(ctx context.Context, userId *uuid.UUID) ([]model.Participation, error) {
	return s.repo.MyParticipation(ctx, userId)
}

func (s *svc) FindByUID(ctx context.Context, uid string) (*model.User, error) {
	return s.repo.FindByUID(ctx, uid)
}

func (s *svc) UpdateAttributes(ctx context.Context, id *uuid.UUID, p map[string]interface{}) error {
	return s.repo.UpdateAttributes(ctx, id, p)
}

func (s *svc) FindByID(ctx context.Context, id *uuid.UUID) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *svc) CreateUser(ctx context.Context, record *auth.UserRecord, req *schema.CreateUserRequest) (*model.User, error) {
	return s.repo.CreateUser(ctx, record, req)
}

func (s *svc) GetTeams(ctx context.Context, userID *uuid.UUID) ([]model.Team, error) {
	return s.repo.GetTeams(ctx, userID)
}

func (s *svc) IsLeader(ctx context.Context, userID *uuid.UUID, teamID *uuid.UUID) (bool, error) {
	return s.repo.IsLeader(ctx, userID, teamID)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
