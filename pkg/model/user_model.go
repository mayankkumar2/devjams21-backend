package model

type User struct {
	BaseModel
	Name    string `json:"Name" gorm:"type:varchar(100)"`
	UID     string `json:"-" gorm:"type:varchar(50);uniqueIndex"`
	Email   string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	RegNo   string `json:"reg_no" gorm:"type:varchar(20)"`
	College string `json:"college" gorm:"type:varchar(100)"`
}
