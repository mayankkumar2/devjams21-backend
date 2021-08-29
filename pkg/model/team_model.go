package model

type Team struct {
	BaseModel
	TeamName string `json:"team_name" gorm:"type:varchar(50);index"`
	JoinCode string `json:"-" gorm:"type:varchar(32);uniqueIndex"`
	TeamXUser []TeamXUser `json:"members,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
