package model

import "github.com/google/uuid"

type Profile struct {
	BaseModel
	UserID          *uuid.UUID `json:"user_id"`
	GithubURL       string     `json:"github_url"`
	LinkedinURL     string     `json:"linkedin_url"`
	DiscordUsername string     `json:"discord_username"`
	TechStack string `json:"tech_stack"`
}
