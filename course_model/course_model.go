package coursemodel

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	TeacherID     uint   `json:"teacher_id"`
	ClassroomCode string `json:"classroom_code"`
	Group         string `json:"group"`
}
