package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	writer := httptest.NewRecorder()
	GetHealth(writer, request, nil)
	response := writer.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected http status 200, status received, %v", response.StatusCode)
	}
}
