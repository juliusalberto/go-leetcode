package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupSubmissionTest(t *testing.T)(*SubmissionHandler, *database.TestDB, int) {
	testDB := database.SetupTestDB(t)
	submissionStore := models.NewSubmissionStore(testDB.DB)
	userStore := models.NewUserStore(testDB.DB)
	handler := NewSubmissionHandler(submissionStore)

	testUser := models.User {
		Username:		"testuser",
		LeetcodeUsername: "leetcode_testuser",
	}

	userStore.CreateUser(&testUser)

	return handler, testDB, testUser.ID
}

func TestCreateSubmission(t *testing.T) {
	handler, testDB, userID := setupSubmissionTest(t)
	defer testDB.Cleanup(t)

	testData := map[string]interface{}{
		"user_id":		userID,
		"title":		"Two Sum",
		"title-slug":	"two-sum",
		"submitted_at": 	time.Now().UTC().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(testData)
	testutils.CheckErr(t, err, "Failed to marshal json data")

	req := httptest.NewRequest("POST", "/submissions/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateSubmission(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rr.Code)
	}

	// Decode standardized response
	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract data from response
	submissionData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal submission data")

	var submission models.Submission
	err = json.Unmarshal(submissionData, &submission)
	testutils.CheckErr(t, err, "Failed to unmarshal submission data")

	if submission.ID == "" {
		t.Error("Expected submission ID in response")
	}

	exists, err := handler.store.CheckSubmissionExists(submission.ID)
	testutils.CheckErr(t, err, "Failed to check if submission exists")

	if !exists {
		t.Error("Submission not created in database")
	}
}

func TestGetSubmission(t *testing.T) {
	handler, testDB, userID := setupSubmissionTest(t)
	defer testDB.Cleanup(t)

	for i := 0; i < 3; i++ {
		sub := models.Submission{
			ID:			fmt.Sprintf("internal-test-%d", i),
			UserID: 	userID,
			Title: 		fmt.Sprintf("Test Problem %d", i),
			TitleSlug: 	fmt.Sprintf("test-problem-%d", i),
			SubmittedAt: time.Now().UTC(),
			CreatedAt: 	time.Now().UTC(),
		}

		err := handler.store.CreateSubmission(sub)
		testutils.CheckErr(t, err, "Failed to create submission")
	}

	// Create URL with query parameter instead of using JSON body
	url := fmt.Sprintf("/submissions?user_id=%d", userID)
	
	// Create request with query parameters and no body
	req := httptest.NewRequest("GET", url, nil)
	rr := httptest.NewRecorder()

	handler.GetSubmissions(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Decode standardized response
	var resp response.Response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract data from response
	submissionsData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal submissions data")

	var submissions []models.Submission
	err = json.Unmarshal(submissionsData, &submissions)
	testutils.CheckErr(t, err, "Failed to unmarshal submissions data")

	if len(submissions) != 3 {
		t.Errorf("Expected 3 submissions, got %d", len(submissions))
	}

	for i := 0; i < 3; i++ {
		if !strings.HasPrefix(submissions[i].TitleSlug, "test-problem") {
			t.Errorf("Expected to have 'test-problem' as the prefix for test slug, got %s", submissions[i].TitleSlug)
		}
	}
}