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
	"strconv"
	"testing"
	"time"
)

func setupReviewTest(t *testing.T)(*ReviewHandler, *database.TestDB, int) {
	testDB := database.SetupTestDB(t)
	reviewStore := models.NewReviewScheduleStore(testDB.DB)
	submissionStore := models.NewSubmissionStore(testDB.DB)
	userStore := models.NewUserStore(testDB.DB)
	handler := NewReviewHandler(reviewStore)

	// create user first
	testUser := models.User{
		Username: "testuser",
		LeetcodeUsername: "test_leetcode",
	}

	userStore.CreateUser(&testUser)

	for i := 0; i < 5; i++ {
		testSub := models.Submission{
			ID: strconv.Itoa(i + 1),
			UserID: testUser.ID,
			Title: fmt.Sprintf("%d Sum", i + 1),
			TitleSlug: fmt.Sprintf("%d-sum", i + 1),
			SubmittedAt: time.Now(),
			CreatedAt: time.Now(),
		}

		if err := submissionStore.CreateSubmission(testSub); err != nil {
			t.Errorf("Failed to create submission")
		}
	}

	return handler, testDB, testUser.ID
}

func TestGetUpcomingReviewHandler(t *testing.T) {
	handler, testDB, userID := setupReviewTest(t)
	defer testDB.Cleanup(t)

	testReview := models.ReviewSchedule{
		SubmissionID: "1",
		NextReviewAt: time.Now(),
		IntervalDays: 1,
		TimesReviewed: 0,
		CreatedAt: time.Now(),
	}

	err := handler.store.CreateReviewSchedule(&testReview)
	testutils.CheckErr(t, err, "Failed to create test review")

	testData := map[string]int{
		"user_id": userID,
	}

	jsonData, err := json.Marshal(testData)
	if err != nil {
    	t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest("GET", "/reviews/upcoming", bytes.NewBuffer(jsonData))
	rr := httptest.NewRecorder()

	handler.GetUpcomingReviews(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var response []models.ReviewSchedule

	// we want to unmarshal the json bytes
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to parse response")

	if len(response) != 1 || response[0].SubmissionID != testReview.SubmissionID {
		t.Errorf("unexpected response: %+v", response)
	}
}