package schema

type FindTeamRequest struct {
	ID string `json:"team_id" form:"team_id" binding:"uuid"`
}

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
}
