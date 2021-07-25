package model

type Team struct {
	BaseModel
	TeamName string `json:"team_name" gorm:"type:varchar(50);index"`
	JoinCode string `json:"join_code" gorm:"type:varchar(32);uniqueIndex"`
}
