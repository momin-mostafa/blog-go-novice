package coursemodel

import (
	dbhandler "backend/db_handler"
	"encoding/json"
	"net/http"
)

type CourseRequestHandler struct{}

func (cRH *CourseRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cRH.HandleRequestDependingOnMethod(w, r)
}

func (cRH *CourseRequestHandler) HandleRequestDependingOnMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World! from course"})
		cRH.GetAllCourses(w)

	case http.MethodPost:
		json.NewEncoder(w).Encode(map[string]string{"message": "course created"})

	case http.MethodDelete:
		json.NewEncoder(w).Encode(map[string]string{"message": "id course deleted"})

	case http.MethodPatch:
		json.NewEncoder(w).Encode(map[string]string{"message": "course updated"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"message": "Method not handled"})
	}

}

func (cRH *CourseRequestHandler) GetAllCourses(w http.ResponseWriter) {
	var course []Course
	if err := dbhandler.GetDBPointer().Find(&course).Error; err != nil {
		http.Error(w, "Error fetching courses", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(course)
}
