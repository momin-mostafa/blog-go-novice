package userModel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string
	Phone          string
	StudentID      string
	PersonalEmail  string
	HashedPassword string
	UnivesityEmail string
}

type CreateUserRequest struct {
	FullName       string `json:"full_name"`
	Phone          string `json:"phone"`
	StudentID      string `json:"student_id"`
	PersonalEmail  string `json:"personal_email"`
	HashedPassword string `json:"password"`
	UnivesityEmail string `json:"university_email"`
}

func (req *CreateUserRequest) createUser() User {
	user := User{
		FullName:       req.FullName,
		Phone:          req.Phone,
		StudentID:      req.StudentID,
		PersonalEmail:  req.PersonalEmail,
		HashedPassword: req.HashedPassword,
		UnivesityEmail: req.UnivesityEmail,
	}

	return user
}
