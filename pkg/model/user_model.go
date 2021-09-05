package model

type User struct {
	BaseModel
	Name           string `json:"Name" gorm:"type:varchar(100)"`
	UID            string `json:"-" gorm:"type:varchar(50);uniqueIndex"`
	Email          string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	RegNo          string `json:"-" gorm:"type:varchar(20)"`
	College        string `json:"college" gorm:"type:varchar(100)"`
	PhotoUrl       string `json:"-" gorm:"type:varchar(200)"`
	PhoneNumber    string `json:"-" gorm:"type:varchar(15)"`
	Gender         string `json:"-" gorm:"type:char(1)"`
	Degree         string `json:"-" gorm:"type:varchar(100)"`
	Stream         string `json:"-" gorm:"type:varchar(100)"`
	GraduationYear string `json:"-" gorm:"type:char(5)"`
	Age            uint   `json:"-"`
	Address        string `json:"-" gorm:"type:varchar(500)"`
	TShirtSize     string `json:"-" gorm:"type:varchar(20)"`
}

//Degree, Stream, city,  graduation year, age, address, T-shirt size
