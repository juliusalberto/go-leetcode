package handlers

import (
	"encoding/json"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func setupSolutionTest(t *testing.T) (*SolutionHandler, *database.TestDB) {
	testDB := database.SetupTestDB(t)
	
	// We're assuming solutions already exist in the database
	// If they don't, you would need to seed them here
	// This function just provides the handler and testDB for cleanup
	
	solutionStore := models.NewSolutionStore(testDB.DB)
	solutionHandler := NewSolutionHandler(solutionStore)
	
	return solutionHandler, testDB
}

func TestGetSolutions(t *testing.T) {
	// Initialize the handler using the setup function
	// This test assumes solutions are already seeded in the database
	// from seed_solutions.sql
	solutionHandler, testDB := setupSolutionTest(t)
	defer testDB.Cleanup(t)

	// Test 1: Get all solutions for problem ID 1 (Two Sum)
	req, err := http.NewRequest("GET", "/api/solutions?id=1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(solutionHandler.GetSolutions)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var solutions []models.Solution
	err = json.Unmarshal(rr.Body.Bytes(), &solutions)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should have at least one solution
	if len(solutions) == 0 {
		t.Errorf("Expected solutions for problem ID 1, got none")
	}

	// Test 2: Get solution for specific language
	req, err = http.NewRequest("GET", "/api/solutions?id=1&language=python", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var solution models.Solution
	err = json.Unmarshal(rr.Body.Bytes(), &solution)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if strings.ToLower(solution.Language) != "python" {
		t.Errorf("Expected language 'python', got '%s'", solution.Language)
	}

	if solution.ProblemID != 1 {
		t.Errorf("Expected problem ID 1, got %d", solution.ProblemID)
	}
}

func TestGetSolutionsForDifferentProblem(t *testing.T) {
	// This test checks fetching solutions for problem ID 9 (Palindrome Number)
	solutionHandler, testDB := setupSolutionTest(t)
	defer testDB.Cleanup(t)

	// Get solutions for problem ID 9
	req, err := http.NewRequest("GET", "/api/solutions?id=9", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(solutionHandler.GetSolutions)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var solutions []models.Solution
	err = json.Unmarshal(rr.Body.Bytes(), &solutions)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should have at least one solution
	if len(solutions) == 0 {
		t.Errorf("Expected solutions for problem ID 9, got none")
	}

	// Verify solutions are for the correct problem
	for _, sol := range solutions {
		if sol.ProblemID != 9 {
			t.Errorf("Expected solution for problem ID 9, got %d", sol.ProblemID)
		}
	}
}

func TestMissingIDParameter(t *testing.T) {
	// Test request without required ID parameter
	solutionHandler, testDB := setupSolutionTest(t)
	defer testDB.Cleanup(t)

	// Request without ID parameter
	req, err := http.NewRequest("GET", "/api/solutions", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(solutionHandler.GetSolutions)
	handler.ServeHTTP(rr, req)

	// Should return a 400 Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestInvalidIDParameter(t *testing.T) {
	// Test request with invalid ID parameter
	solutionHandler, testDB := setupSolutionTest(t)
	defer testDB.Cleanup(t)

	// Request with non-numeric ID
	req, err := http.NewRequest("GET", "/api/solutions?id=invalid", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(solutionHandler.GetSolutions)
	handler.ServeHTTP(rr, req)

	// Should return a 400 Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestNonExistentProblem(t *testing.T) {
	// Test request for a problem that doesn't exist
	solutionHandler, testDB := setupSolutionTest(t)
	defer testDB.Cleanup(t)

	// Use a very large ID that's unlikely to exist
	nonExistentID := 999999
	req, err := http.NewRequest("GET", "/api/solutions?id="+strconv.Itoa(nonExistentID), nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(solutionHandler.GetSolutions)
	handler.ServeHTTP(rr, req)

	// Should still return 200 OK, but with empty array
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var solutions []models.Solution
	err = json.Unmarshal(rr.Body.Bytes(), &solutions)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should have no solutions
	if len(solutions) != 0 {
		t.Errorf("Expected no solutions for non-existent problem, got %d", len(solutions))
	}
}