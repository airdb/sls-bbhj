package po

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Nickname   string `gorm:"type:varchar(64)"`
	Headimgurl string `gorm:"type:varchar(128)"`
	Token      string `gorm:"type:varchar(128)"`
}

func List(voUser string) *User {
	var user User

	return &user
}
