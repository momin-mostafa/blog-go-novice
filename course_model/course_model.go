package coursemodel

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	TeacherID     uint
	ClassroomCode string
	Group         string
}
