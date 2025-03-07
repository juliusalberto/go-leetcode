package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupProblemTest(t *testing.T) (*ProblemHandler, *database.TestDB) {
	testDB := database.SetupTestDB(t)
	problemStore := models.NewProblemStore(testDB.DB)
	handler := &ProblemHandler{store: problemStore}
	return handler, testDB
}

func TestGetProblemByID(t *testing.T) {
	handler, testDB := setupProblemTest(t)
	defer testDB.Cleanup(t)

	// Find a test problem from the database
	var testProblemID int
	err := testDB.DB.QueryRow("SELECT id FROM problems LIMIT 1").Scan(&testProblemID)
	testutils.CheckErr(t, err, "Failed to find test problem")

	// Create test request
	reqBody := map[string]int{"id": testProblemID}
	jsonData, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/problems/id", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	handler.GetProblemByID(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode response
	var response models.Problem
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to decode response")

	// Verify problem ID matches
	if response.ID != testProblemID {
		t.Errorf("Expected problem ID %d, got %d", testProblemID, response.ID)
	}
}

func TestGetProblemByFrontendID(t *testing.T) {
	handler, testDB := setupProblemTest(t)
	defer testDB.Cleanup(t)

	// Find a test problem from the database
	testFrontendID := 1966

	// Create test request
	reqBody := map[string]int{"frontend_id": testFrontendID}
	jsonData, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("GET", "/problems/frontend_id", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	handler.GetProblemByFrontendID(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode response
	var response models.Problem
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to decode response")

	// Verify frontend ID matches
	if response.FrontendID != testFrontendID {
		t.Errorf("Expected frontend ID %d, got %v", testFrontendID, response.FrontendID)
	}
}

func TestGetProblemBySlug(t *testing.T) {
	handler, testDB := setupProblemTest(t)
	defer testDB.Cleanup(t)

	// Find a test problem slug from the database
	testSlug := "two-sum"

	// Create test request
	reqBody := map[string]string{"slug": testSlug}
	jsonData, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/problems/slug", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	handler.GetProblemBySlug(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode response
	var response models.Problem
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to decode response")

	// Verify slug matches
	if response.TitleSlug != testSlug {
		t.Errorf("Expected slug %s, got %s", testSlug, response.TitleSlug)
	}
}

func TestGetProblemList(t *testing.T) {
	handler, testDB := setupProblemTest(t)
	defer testDB.Cleanup(t)

	// Create test cases
	testCases := []struct {
		name           string
		queryParams    map[string]string
		expectedStatus int
		validateResult func(*testing.T, models.ProblemList)
	}{
		{
			name:           "Default Parameters",
			queryParams:    map[string]string{},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, result models.ProblemList) {
				if len(result.Problems) > 20 {
					t.Errorf("Expected at most 20 problems, got %d", len(result.Problems))
				}
				if result.Total < len(result.Problems) {
					t.Errorf("Total count %d should be at least as large as returned problems %d", 
						result.Total, len(result.Problems))
				}
			},
		},
		{
			name: "Pagination",
			queryParams: map[string]string{
				"limit":  "5",
				"offset": "5",
			},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, result models.ProblemList) {
				if len(result.Problems) > 5 {
					t.Errorf("Expected at most 5 problems, got %d", len(result.Problems))
				}
			},
		},
		{
			name: "Filter By Difficulty",
			queryParams: map[string]string{
				"difficulty": "Easy",
			},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, result models.ProblemList) {
				if len(result.Problems) > 0 {
					if result.Problems[0].Difficulty != "Easy" {
						t.Errorf("Expected Easy difficulty, got %s", result.Problems[0].Difficulty)
					}
				}
			},
		},
		{
			name: "Search Keyword",
			queryParams: map[string]string{
				"search": "sum",
			},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, result models.ProblemList) {
				// This test might be less reliable depending on your test data
				// Just check if we got a response
				if result.Total == 0 && testHasProblemWithKeyword(t, testDB, "sum") {
					t.Errorf("Expected to find problems with 'sum' but got none")
				}
			},
		},
		{
			name: "Sorting",
			queryParams: map[string]string{
				"order_by":  "difficulty",
				"order_dir": "desc",
			},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, result models.ProblemList) {
				if len(result.Problems) > 1 {
					// Check if first problem difficulty is >= second problem difficulty
					if difficultyValue(result.Problems[0].Difficulty) < difficultyValue(result.Problems[1].Difficulty) {
						t.Errorf("Expected descending difficulty order, got %s before %s", 
							result.Problems[0].Difficulty, result.Problems[1].Difficulty)
					}
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Build URL with query parameters
			url := "/problems"
			if len(tc.queryParams) > 0 {
				url += "?"
				for key, value := range tc.queryParams {
					url += key + "=" + value + "&"
				}
				// Remove trailing &
				url = url[:len(url)-1]
			}

			req := httptest.NewRequest("GET", url, nil)
			rr := httptest.NewRecorder()

			// Create router to parse query parameters
			r := chi.NewRouter()
			r.Get("/problems", handler.GetProblemList)
			r.ServeHTTP(rr, req)

			// Check response status
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			// Skip validation if not expecting OK
			if tc.expectedStatus != http.StatusOK {
				return
			}

			// Decode response
			var response models.ProblemList
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			testutils.CheckErr(t, err, "Failed to decode response")

			// Validate result
			tc.validateResult(t, response)
		})
	}
}

// Helper functions

func difficultyValue(difficulty string) int {
	switch difficulty {
	case "Easy":
		return 1
	case "Medium":
		return 2
	case "Hard":
		return 3
	default:
		return 0
	}
}

func testHasProblemWithKeyword(t *testing.T, testDB *database.TestDB, keyword string) bool {
	var count int
	err := testDB.DB.QueryRow("SELECT COUNT(*) FROM problems WHERE title ILIKE $1", 
		"%"+keyword+"%").Scan(&count)
	testutils.CheckErr(t, err, "Failed to query problems with keyword")
	return count > 0
}