package course_request_handler_test

import (
	coursemodel "backend/course_model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCourseRequestHandler(t *testing.T) {
	handler := &coursemodel.CourseRequestHandler{}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.HandleRequestDependingOnMethod(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	handler.HandleRequestDependingOnMethod(w, req)

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	expected := "Hello, World!"
	if body["message"] != expected {
		t.Errorf("expected message %q, got %q", expected, body["message"])
	}
}
