package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
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

	// check if the submission exist
	var response models.Submission

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	if response.ID == "" {
		t.Error("Expected submission ID in response")
	}

	exists, err := handler.store.CheckSubmissionExists(response.ID)
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

	testData := map[string]int {
		"user_id": userID,
	}

	// send the testdata towards the endpoint
	jsonData, err := json.Marshal(testData)
	testutils.CheckErr(t, err, "Failed to marshal json")

	req := httptest.NewRequest("GET", "/submissions", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.GetSubmissions(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var submissions []models.Submission
	err = json.Unmarshal(rr.Body.Bytes(), &submissions)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	if len(submissions) != 3 {
		t.Errorf("Expected 3 submissions, got %d", len(submissions))
	}

	for i := 0; i < 3; i++ {
		if !strings.HasPrefix(submissions[i].TitleSlug, "test-problem") {
			t.Errorf("Expected to have 'test-problem' as the prefix for test slug, got %s", submissions[i].TitleSlug)
		}
	}
}