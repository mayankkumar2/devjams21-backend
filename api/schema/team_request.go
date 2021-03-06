package schema

import "github.com/google/uuid"

type FindTeamRequest struct {
	ID string `json:"team_id" form:"team_id" type:"gorm:uuid"`
}
type UpdateTeamNameRequest struct {
	ID   *uuid.UUID `json:"team_id" form:"team_id" type:"gorm:uuid"`
	Name string     `json:"team_name"`
}

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
}

type JoinTeamRequest struct {
	Code string `json:"code"`
}

type TeamIDRequest struct {
	ID *uuid.UUID `json:"team_id"`
}
type TeamXMemberRequest struct {
	ID       *uuid.UUID `json:"team_id"`
	MemberID *uuid.UUID `json:"member_id"`
}
