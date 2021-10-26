package model

type Submission struct {
	BaseModel
	Freeze bool `json:"freeze" gorm:"default:false"`
	Meta   JSON `json:"meta" gorm:"type:json"`
}
