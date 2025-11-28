package test

import (
	coursemodel "backend/course_model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// ---------------- GET /course?Group=A&TeacherID=10 ----------------
func TestGetAllCourses_GroupAndTeacher(t *testing.T) {
	db, mock := setupMockDB(t)
	defer closeDB(db)

	// Mock query
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "teacher_id", "classroom_code", "group"}).
		AddRow(1, time.Now(), time.Now(), nil, 10, "101", "A")

	mock.ExpectQuery(`SELECT .* FROM "courses".*WHERE.*`).
		WithArgs("A", 10).
		WillReturnRows(rows)

	req := httptest.NewRequest("GET", "/course?group=A&teacher_id=10", nil)
	rr := httptest.NewRecorder()

	handler := &coursemodel.CourseRequestHandler{}
	handler.GetAllCourses(req, rr, db)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	var courses []coursemodel.Course
	if err := json.NewDecoder(rr.Body).Decode(&courses); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(courses) != 1 {
		t.Errorf("expected 1 course, got %d", len(courses))
	}
}

// ---------------- GET /course?Group=A ----------------
func TestGetAllCourses_GroupOnly(t *testing.T) {
	db, mock := setupMockDB(t)
	defer closeDB(db)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "teacher_id", "classroom_code", "group"}).
		AddRow(1, time.Now(), time.Now(), nil, 10, "101", "A")

	mock.ExpectQuery(`SELECT .* FROM "courses".*WHERE.*group.*`).
		WithArgs("A").
		WillReturnRows(rows)

	req := httptest.NewRequest("GET", "/course?group=A", nil)
	rr := httptest.NewRecorder()

	handler := &coursemodel.CourseRequestHandler{}
	handler.GetAllCourses(req, rr, db)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	var courses []coursemodel.Course
	if err := json.NewDecoder(rr.Body).Decode(&courses); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(courses) != 1 {
		t.Errorf("expected 1 course, got %d", len(courses))
	}
}

// ---------------- GET /course?TeacherID=10 ----------------
func TestGetAllCourses_TeacherOnly(t *testing.T) {
	db, mock := setupMockDB(t)
	defer closeDB(db)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "teacher_id", "classroom_code", "group"}).
		AddRow(1, time.Now(), time.Now(), nil, 10, "101", "A")

	mock.ExpectQuery(`SELECT .* FROM "courses".*WHERE.*teacher_id.*`).
		WithArgs(10).
		WillReturnRows(rows)

	req := httptest.NewRequest("GET", "/course?teacher_id=10", nil)
	rr := httptest.NewRecorder()

	handler := &coursemodel.CourseRequestHandler{}
	handler.GetAllCourses(req, rr, db)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	var courses []coursemodel.Course
	if err := json.NewDecoder(rr.Body).Decode(&courses); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(courses) != 1 {
		t.Errorf("expected 1 course, got %d", len(courses))
	}
}

// ---------------- GET /course (no params) ----------------
func TestGetAllCourses_NoParams(t *testing.T) {
	db, mock := setupMockDB(t)
	defer closeDB(db)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "teacher_id", "classroom_code", "group"}).
		AddRow(1, time.Now(), time.Now(), nil, 10, "101", "A").
		AddRow(2, time.Now(), time.Now(), nil, 11, "102", "B")

	mock.ExpectQuery(`SELECT .* FROM "courses"`).
		WillReturnRows(rows)

	req := httptest.NewRequest("GET", "/course", nil)
	rr := httptest.NewRecorder()

	handler := &coursemodel.CourseRequestHandler{}
	handler.GetAllCourses(req, rr, db)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	var courses []coursemodel.Course
	if err := json.NewDecoder(rr.Body).Decode(&courses); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(courses) != 2 {
		t.Errorf("expected 2 courses, got %d", len(courses))
	}
}

func RunAllCoursesTests(t *testing.T) {
	t.Run("GetAllCourses_GroupAndTeacher", TestGetAllCourses_GroupAndTeacher)
	t.Run("GetAllCourses_GroupOnly", TestGetAllCourses_GroupOnly)
	t.Run("GetAllCourses_TeacherOnly", TestGetAllCourses_TeacherOnly)
	t.Run("GetAllCourses_NoParams", TestGetAllCourses_NoParams)
}
