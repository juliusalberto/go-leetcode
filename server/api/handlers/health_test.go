package handlers

import (
	"encoding/json"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	// Test the response
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp response.Response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	// Check that we have data and no errors
	if resp.Data == nil {
		t.Errorf("response data should not be nil")
	}
	if len(resp.Errors) > 0 {
		t.Errorf("response should not contain errors, got %v", resp.Errors)
	}
}