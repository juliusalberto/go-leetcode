package handlers

import (
	"bytes"
	"encoding/json"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
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
	defer testDB.Cleanup()

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

	if rr.Code != http.StatusAccepted {
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