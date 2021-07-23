package model

type User struct {
	baseModel
	Name    string `json:"Name" gorm:"type:varchar(100)"`
	UID     string `json:"-" gorm:"type:varchar(50);uniqueIndex"`
	Email   string `json:"email" gorm:"type:varchar(100)"`
	RegNo   string `json:"reg_no" gorm:"type:varchar(20)"`
	College string `json:"college" gorm:"type:varchar(100)"`
}
