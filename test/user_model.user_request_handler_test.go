package test

import (
	userModel "backend/user"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	db, mock := setupMockDB(t)
	defer closeDB(db)

	handler := &userModel.UserRequestHandler{}
	userData := `{"full_name":"tamim mostafa","student_id":"12345"}`
	req := httptest.NewRequest("POST", "/user", bytes.NewBufferString(userData))
	rr := httptest.NewRecorder()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "tamim mostafa", "12345").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	handler.CreateUser(rr, req, db)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	var user userModel.User
	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if user.FullName != "tamim mostafa" || user.StudentID != "12345" {
		t.Errorf("unexpected user data: %+v", user)
	}
}
