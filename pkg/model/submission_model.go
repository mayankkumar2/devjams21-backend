package model

type Submission struct {
	BaseModel
	Freeze bool `json:"-" gorm:"default:false"`
	Meta   JSON `json:"meta" gorm:"type:json"`
}
