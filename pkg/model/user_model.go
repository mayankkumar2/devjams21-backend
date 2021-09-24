package model

type User struct {
	BaseModel
	FirstName      string         `json:"first_name" gorm:"type:varchar(100)"`
	LastName       string         `json:"last_name" gorm:"type:varchar(100)" `
	UID            string         `json:"-" gorm:"type:varchar(50);uniqueIndex"`
	Email          string         `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	RegNo          string         `json:"-" gorm:"type:varchar(20)"`
	College        string         `json:"college" gorm:"type:varchar(100)"`
	PhotoUrl       string         `json:"photo_url" gorm:"type:varchar(200)"`
	PhoneNumber    string         `json:"-" gorm:"type:varchar(15)"`
	Gender         string         `json:"-" gorm:"type:varchar(50)"`
	Degree         string         `json:"-" gorm:"type:varchar(100)"`
	Stream         string         `json:"-" gorm:"type:varchar(100)"`
	GraduationYear string         `json:"-" gorm:"type:char(5)"`
	Age            uint           `json:"-"`
	Address        string         `json:"-" gorm:"type:varchar(500)"`
	TShirtSize     string         `json:"-" gorm:"type:varchar(20)"`
	MessageBoard   []MessageBoard `json:"messages"`
	FCMToken       string         `json:"-" gorm:"type:varchar(1000)"`
	AgreeMLHCodeOfConduct bool `json:"-"`
	AgreeMLHPrivacyPolicy bool `json:"-"`
	AgreeMLHEventDetails bool `json:"-"`
}

//Degree, Stream, city,  graduation year, age, address, T-shirt size
