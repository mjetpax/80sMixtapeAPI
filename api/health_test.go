package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	GetHealth(w, req, nil)
	response := w.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected http status 200, status received, %v", response.StatusCode)
	}
}
