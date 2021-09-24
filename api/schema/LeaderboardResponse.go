package schema

type LeaderboardResponse struct {
	Name     string `json:"name"`
	Scr      uint   `json:"score"`
	PhotoUrl string `json:"photo_url"`
}
