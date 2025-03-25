package handlers

import (
	"encoding/json"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	var solutions map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &solutions)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should have at least one solution
	if len(solutions) == 0 {
		t.Errorf("Expected solutions for problem ID 1, got none")
	}

	// Check if the solution exists for the given language
	if _, ok := solutions["Python"]; !ok {
		t.Errorf("Expected a Python solution, but found none")
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
		t.Logf("Response body: %s", rr.Body.String())
		solutions = make(map[string]string) // Reset the solutions map
		err = json.Unmarshal(rr.Body.Bytes(), &solutions)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
	
		// Should only have the Python solution
		if len(solutions) != 1 {
			t.Errorf("Expected only 1 solution (Python), got %d", len(solutions))
		}
	
		// Check if the Python solution exists
		if _, ok := solutions["Python"]; !ok {
			t.Errorf("Expected a Python solution, but found none")
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
		var solutions map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &solutions)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
	
		// Should have at least one solution
		if len(solutions) == 0 {
			t.Errorf("Expected solutions for problem ID 9, got none")
		}
	
		// We can't verify the problem ID directly from the solution code
		// But we can check that we got some solution code for at least one language
		foundSolution := false
		for language, code := range solutions {
			if len(code) > 0 {
				foundSolution = true
				t.Logf("Found solution for language: %s", language)
			}
		}
		
		if !foundSolution {
			t.Errorf("Expected at least one valid solution for problem ID 9")
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
		var solutions map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &solutions)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
	
		// Should have no solutions
		if len(solutions) != 0 {
			t.Errorf("Expected no solutions for non-existent problem, got %d", len(solutions))
		}
}
