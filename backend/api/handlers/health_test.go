package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHealthEndpoint(t* testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
    rr := httptest.NewRecorder()

    // Test the response
    handler := http.HandlerFunc(HealthCheck)
    handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}