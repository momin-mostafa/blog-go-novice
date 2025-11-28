package userModel

import (
	dbhandler "backend/db_handler"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type UserRequestHandler struct{}

func (uRH *UserRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		uRH.getUser(w, r)
	case http.MethodPost:
		uRH.createUser(w, r)
	}
}

func (uRH *UserRequestHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.FullName == "" || req.StudentID == "" {
		http.Error(w, "full_name & student_id cannot be empty to create user", http.StatusBadRequest)
		return
	}

	user := req.createUser()

	dbhandler.GetDBPointer().Create(&user)

	json.NewEncoder(w).Encode(user)
}

func (uRH *UserRequestHandler) getUser(w http.ResponseWriter, r *http.Request) {
	student_id := r.URL.Query().Get("student_id")
	if student_id == "" {
		http.Error(w, "student_id cannot be empty", http.StatusBadRequest)
	}

	var user User
	err := dbhandler.GetDBPointer().Where("student_id = ?", student_id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Fprintf(w, "User not found")
		} else {
			fmt.Fprintf(w, "DB error:%+v", err)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}
