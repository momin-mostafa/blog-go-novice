package coursemodel

import (
	dbhandler "backend/db_handler"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type CourseRequestHandler struct{}

func (cRH *CourseRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cRH.HandleRequestDependingOnMethod(w, r, dbhandler.GetDBPointer())
}

func (cRH *CourseRequestHandler) HandleRequestDependingOnMethod(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.Method {
	case http.MethodGet:
		cRH.GetAllCourses(r, w, db)

	case http.MethodPost:
		cRH.CreateCourse(w, db, r)

	case http.MethodDelete:
		json.NewEncoder(w).Encode(map[string]string{"message": "id course deleted"})

	case http.MethodPatch:
		json.NewEncoder(w).Encode(map[string]string{"message": "course updated"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"message": "Method not handled"})
	}

}

func (cRH *CourseRequestHandler) GetAllCourses(r *http.Request, w http.ResponseWriter, db *gorm.DB) {
	if r.URL.Query().Get("Group") != "" {
		group := r.URL.Query().Get("Group")
		var courses []Course
		if err := db.Where("Group = ?", group).Find(&courses).Error; err != nil {
			http.Error(w, "Error fetching courses", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
		return
	}
	if r.URL.Query().Get("TeacherID") != "" {
		teacherID := r.URL.Query().Get("TeacherID")
		var courses []Course
		if err := db.Where("teacher_id = ?", teacherID).Find(&courses).Error; err != nil {
			http.Error(w, "Error fetching courses", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
		return
	}
	// handle case where multiple query parameters are provided
	if r.URL.Query().Get("Group") != "" && r.URL.Query().Get("TeacherID") != "" {
		group := r.URL.Query().Get("Group")
		teacherID := r.URL.Query().Get("TeacherID")
		var courses []Course
		if err := db.Where("Group = ? AND TeacherID = ?", group, teacherID).Find(&courses).Error; err != nil {
			http.Error(w, "Error fetching courses", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
		return
	}

	var course []Course
	if err := dbhandler.GetDBPointer().Find(&course).Error; err != nil {
		http.Error(w, "Error fetching courses", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(course)
}

func (cRH *CourseRequestHandler) CreateCourse(w http.ResponseWriter, db *gorm.DB, r *http.Request) {
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid course data", http.StatusBadRequest)
		return
	}
	if err := db.Create(&course).Error; err != nil {
		http.Error(w, "Error creating course", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}
