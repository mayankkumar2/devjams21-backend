package user

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
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
