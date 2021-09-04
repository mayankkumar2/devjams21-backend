package schema

import "github.com/google/uuid"

type CreateUserInfoRequest struct {
	UserID          *uuid.UUID `json:"user_id"`
	GithubURL       string     `json:"github_url"`
	LinkedinURL     string     `json:"linkedin_url"`
	DiscordUsername string     `json:"discord_username"`
}

type DeleteUserInfoRequest struct {
	UserID *uuid.UUID `json:"user_id"`
}

type UpdateUserInfoRequest struct {
	UserID          *uuid.UUID `json:"user_id"`
	GithubURL       string     `json:"github_url"`
	LinkedinURL     string     `json:"linkedin_url"`
	DiscordUsername string     `json:"discord_username"`
}

type GetUserInfoRequest struct {
	UserID *uuid.UUID `json:"user_id"`
}
