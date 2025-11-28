package coursemodel

import (
	dbhandler "backend/db_handler"
	"encoding/json"
	"net/http"
	"strconv"

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
		cRH.CreateCourse(w, r, db)

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
	group := r.URL.Query().Get("group")
	teacherIDStr := r.URL.Query().Get("teacher_id")
	var teacherID int
	var err error

	if teacherIDStr != "" {
		teacherID, err = strconv.Atoi(teacherIDStr)
		if err != nil {
			http.Error(w, "TeacherID must be a number", http.StatusBadRequest)
			return
		}
	}

	// Combined query first
	if group != "" && teacherIDStr != "" {
		var courses []Course
		if err := db.Where(&Course{Group: group, TeacherID: uint(teacherID)}).Find(&courses).Error; err != nil {
			http.Error(w, "Error fetching courses", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
		return
	}

	// Single filters
	if group != "" {
		var courses []Course
		if err := db.Where(&Course{Group: group}).Find(&courses).Error; err != nil {
			http.Error(w, "Error fetching courses", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
		return
	}

	if teacherIDStr != "" {
		var courses []Course
		if err := db.Where(&Course{TeacherID: uint(teacherID)}).Find(&courses).Error; err != nil {
			http.Error(w, "Error fetching courses", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(courses)
		return
	}

	// No filters, return all
	var courses []Course
	if err := db.Find(&courses).Error; err != nil {
		http.Error(w, "Error fetching courses", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(courses)
}

func (cRH *CourseRequestHandler) CreateCourse(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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
