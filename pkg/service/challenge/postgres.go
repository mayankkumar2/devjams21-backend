package challenge

import "gorm.io/gorm"

type repo struct {
	DB *gorm.DB
}


func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}