package handlers

import (
	"encoding/json"
	"fmt"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
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

	url := fmt.Sprintf("/problems?id=%d", testProblemID)

	// Create test request
	req := httptest.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	handler.GetProblemByID(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode standardized response
	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to decode response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract problem data from response
	problemData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal problem data")

	var problem models.Problem
	err = json.Unmarshal(problemData, &problem)
	testutils.CheckErr(t, err, "Failed to unmarshal problem data")

	// Verify problem ID matches
	if problem.ID != testProblemID {
		t.Errorf("Expected problem ID %d, got %d", testProblemID, problem.ID)
	}
}

func TestGetProblemByFrontendID(t *testing.T) {
	handler, testDB := setupProblemTest(t)
	defer testDB.Cleanup(t)

	// Find a test problem from the database
	testFrontendID := 12

	// Create test request
	req := httptest.NewRequest("GET", "/problems?frontend_id=12", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	handler.GetProblemByFrontendID(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode standardized response
	var resp response.Response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to decode response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract problem data from response
	problemData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal problem data")

	var problem models.Problem
	err = json.Unmarshal(problemData, &problem)
	testutils.CheckErr(t, err, "Failed to unmarshal problem data")

	// Verify frontend ID matches
	if problem.FrontendID != testFrontendID {
		t.Errorf("Expected frontend ID %d, got %v", testFrontendID, problem.FrontendID)
	}
}

func TestGetProblemBySlug(t *testing.T) {
	handler, testDB := setupProblemTest(t)
	defer testDB.Cleanup(t)

	// Find a test problem slug from the database
	testSlug := "two-sum"

	// Create test request
	req := httptest.NewRequest("POST", "/problems?slug=two-sum", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	handler.GetProblemBySlug(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode standardized response
	var resp response.Response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to decode response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract problem data from response
	problemData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal problem data")

	var problem models.Problem
	err = json.Unmarshal(problemData, &problem)
	testutils.CheckErr(t, err, "Failed to unmarshal problem data")

	// Verify slug matches
	if problem.TitleSlug != testSlug {
		t.Errorf("Expected slug %s, got %s", testSlug, problem.TitleSlug)
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
		validateResult func(*testing.T, []models.Problem, int)
	}{
		{
			name:           "Default Parameters",
			queryParams:    map[string]string{},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, problems []models.Problem, total int) {
				if len(problems) > 20 {
					t.Errorf("Expected at most 20 problems, got %d", len(problems))
				}
				if total < len(problems) {
					t.Errorf("Total count %d should be at least as large as returned problems %d", 
						total, len(problems))
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
			validateResult: func(t *testing.T, problems []models.Problem, total int) {
				if len(problems) > 5 {
					t.Errorf("Expected at most 5 problems, got %d", len(problems))
				}
			},
		},
		{
			name: "Filter By Difficulty",
			queryParams: map[string]string{
				"difficulty": "Easy",
			},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, problems []models.Problem, total int) {
				if len(problems) > 0 {
					if problems[0].Difficulty != "Easy" {
						t.Errorf("Expected Easy difficulty, got %s", problems[0].Difficulty)
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
			validateResult: func(t *testing.T, problems []models.Problem, total int) {
				// This test might be less reliable depending on your test data
				// Just check if we got a response
				if total == 0 && testHasProblemWithKeyword(t, testDB, "sum") {
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
			validateResult: func(t *testing.T, problems []models.Problem, total int) {
				if len(problems) > 1 {
					// Check if first problem difficulty is >= second problem difficulty
					if difficultyValue(problems[0].Difficulty) < difficultyValue(problems[1].Difficulty) {
						t.Errorf("Expected descending difficulty order, got %s before %s", 
							problems[0].Difficulty, problems[1].Difficulty)
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

			// Decode standardized response first
			var resp response.Response
			err := json.Unmarshal(rr.Body.Bytes(), &resp)
			testutils.CheckErr(t, err, "Failed to decode standardized response")

			// Check for pagination metadata
			var pagination struct {
				Total   int `json:"total"`
				Page    int `json:"page"`
				PerPage int `json:"per_page"`
			}
			
			if resp.Meta != nil && resp.Meta.Pagination != nil {
				pagination.Total = resp.Meta.Pagination.Total
				pagination.Page = resp.Meta.Pagination.Page
				pagination.PerPage = resp.Meta.Pagination.PerPage
			}

			// Check for errors
			if len(resp.Errors) > 0 {
				t.Errorf("Response contains errors: %v", resp.Errors)
			}

			// Extract problem data from response
			problemsData, err := json.Marshal(resp.Data)
			testutils.CheckErr(t, err, "Failed to marshal problems data")

			var problems []models.Problem
			err = json.Unmarshal(problemsData, &problems)
			testutils.CheckErr(t, err, "Failed to unmarshal problems data")

			// Validate result
			tc.validateResult(t, problems, pagination.Total)
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