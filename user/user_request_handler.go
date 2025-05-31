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

	if req.Name == "" || req.Phone == "" {
		http.Error(w, "name and phone cannot be empty", http.StatusBadRequest)
		return
	}

	user := req.createUser()

	dbhandler.GetDBPointer().Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (uRH *UserRequestHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		http.Error(w, "user_id cannot be empty", http.StatusBadRequest)
	}

	var user User
	err := dbhandler.GetDBPointer().First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Fprintf(w, "User not found")
		} else {
			fmt.Fprintf(w, "DB error:%+v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
