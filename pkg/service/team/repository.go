package team

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Repository interface {
	UpdateTeamName(ctx context.Context, teamId *uuid.UUID, teamName string) error
	FindByID(ctx context.Context, id *uuid.UUID) (*model.Team, error)
	CreateTeam(ctx context.Context, usr *model.User, teamName string) (*model.Team, error)
	GetMembers(ctx context.Context, id *uuid.UUID) ([]model.TeamXUser, error)
	GetTeamMember(ctx context.Context, teamId, userId *uuid.UUID) (*model.TeamXUser, error)
	UpdateTeamCode(ctx context.Context, team *model.Team) error
	FindByJoinCode(ctx context.Context, code string) (*model.Team, error)
	JoinTeam(ctx context.Context, team *model.Team, usr *model.User) error
	RemoveFromTeam(ctx context.Context, team *model.Team, usr *model.User) error
	FetchTeamMembers(ctx context.Context, teamId *uuid.UUID) ([]model.TeamXUser, error)
	AcceptJoinRequest(ctx context.Context, team *model.Team, userID *uuid.UUID) error
}
