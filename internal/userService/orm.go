package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);"`
	Pass  string `gorm:"type:varchar(100);"`
}
