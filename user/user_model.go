package userModel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Phone string
	FB_ID string
	Email string
}
