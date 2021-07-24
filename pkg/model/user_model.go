package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	baseModel
	Name    string `json:"Name" gorm:"type:varchar(100)"`
	UID     string `json:"-" gorm:"type:varchar(50);uniqueIndex"`
	Email   string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	RegNo   string `json:"reg_no" gorm:"type:varchar(20)"`
	College string `json:"college" gorm:"type:varchar(100)"`
}


func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	id := uuid.New()
	u.ID = &id
	return nil
}