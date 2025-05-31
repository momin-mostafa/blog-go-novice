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

type CreateUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	FB_ID string `json:"fb_id"`
	Email string `json:"email"`
}

func (req *CreateUserRequest) createUser() User {
	user := User{
		Name:  req.Name,
		Phone: req.Phone,
		FB_ID: req.FB_ID,
		Email: req.Email,
	}

	return user
}
