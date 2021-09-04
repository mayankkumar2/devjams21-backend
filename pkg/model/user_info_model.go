package model

import "github.com/google/uuid"

type UserInfo struct {
	UserID          *uuid.UUID `json:"user_id" gorm:"unique"`
	GithubURL       string     `json:"github_url"`
	LinkedinURL     string     `json:"linkedin_url"`
	DiscordUsername string     `json:"discord_username"`
}
