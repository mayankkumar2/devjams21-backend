package team

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
)

type Service interface {
	// FindByID finds a team by its id
	FindByID(ctx context.Context, id *uuid.UUID) (*model.Team, error)
	// CreateTeam creates a team with usr and teamName supplied
	CreateTeam(ctx context.Context, usr *model.User, teamName string) (*model.Team, error)
	// GetMembers gets the members of a team
	GetMembers(ctx context.Context, id *uuid.UUID) ([]model.TeamXUser, error)
	GetTeamMember(ctx context.Context, teamId, userId *uuid.UUID) (*model.TeamXUser, error)
	UpdateTeamCode(ctx context.Context, team *model.Team) error
	FindByJoinCode(ctx context.Context, code string) (*model.Team, error)
	JoinTeam(ctx context.Context, team *model.Team, usr *model.User) error
	RemoveFromTeam(ctx context.Context, team *model.Team, usr *model.User) error
	AcceptJoinRequest(ctx context.Context, team *model.Team, userID *uuid.UUID) error
	FetchTeamMembers(ctx context.Context, teamId *uuid.UUID) ([]model.TeamXUser, error)
	CountTeamByTeamName(ctx context.Context, teamName string,  eventId *uuid.UUID) (int64,error)
	UpdateTeamName(ctx context.Context, teamId *uuid.UUID, teamName string) error
}

type svc struct {
	repo Repository
}

func (s *svc) CountTeamByTeamName(ctx context.Context, teamName string,  eventId *uuid.UUID) (int64,error) {
	return s.repo.CountTeamByTeamName(ctx, teamName, eventId)
}

func (s *svc) UpdateTeamName(ctx context.Context, teamId *uuid.UUID, teamName string) error {
	return s.repo.UpdateTeamName(ctx, teamId, teamName)
}

func (s *svc) FetchTeamMembers(ctx context.Context, teamId *uuid.UUID) ([]model.TeamXUser, error) {
	return s.repo.FetchTeamMembers(ctx, teamId)
}

func (s *svc) FindByJoinCode(ctx context.Context, code string) (*model.Team, error) {
	return s.repo.FindByJoinCode(ctx, code)
}

func (s *svc) JoinTeam(ctx context.Context, team *model.Team, usr *model.User) error {
	return s.repo.JoinTeam(ctx, team, usr)
}

func (s *svc) RemoveFromTeam(ctx context.Context, team *model.Team, usr *model.User) error {
	return s.repo.RemoveFromTeam(ctx, team, usr)
}

func (s *svc) AcceptJoinRequest(ctx context.Context, team *model.Team, userID *uuid.UUID) error {
	return s.repo.AcceptJoinRequest(ctx, team, userID)
}

func (s *svc) UpdateTeamCode(ctx context.Context, team *model.Team) error {
	return s.repo.UpdateTeamCode(ctx, team)
}

func (s *svc) GetMembers(ctx context.Context, id *uuid.UUID) ([]model.TeamXUser, error) {
	return s.repo.GetMembers(ctx, id)
}

func (s *svc) GetTeamMember(ctx context.Context, teamId, userId *uuid.UUID) (*model.TeamXUser, error) {
	return s.repo.GetTeamMember(ctx, teamId, userId)
}

func (s *svc) CreateTeam(ctx context.Context, usr *model.User, teamName string) (*model.Team, error) {
	return s.repo.CreateTeam(ctx, usr, teamName)
}

func (s *svc) FindByID(ctx context.Context, id *uuid.UUID) (*model.Team, error) {
	return s.repo.FindByID(ctx, id)
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}
