package schema

type LeaderboardResponse struct {
	FirstName     string `json:"first_name"`
	LastName string `json:"last_name"`
	Scr      uint   `json:"score"`
	PhotoUrl string `json:"photo_url"`
}
