package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string // The username
	Password string // The user password
	Mobile   string // The mobile phone number
	Type     int64  // The user type, 0:normal,1:vip, for test golang keyword
}

func (User) TableName() string {
	return "user"
}
