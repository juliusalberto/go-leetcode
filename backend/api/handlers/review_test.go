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

func TestUpdateReviewSchedule(t *testing.T) {
	handler, testDB, _ := setupReviewTest(t)
	defer testDB.Cleanup(t)

	testReview := models.ReviewSchedule{
		SubmissionID: "1",
		NextReviewAt: time.Now().UTC(),
		IntervalDays: 1,
		TimesReviewed: 0,
		CreatedAt: time.Now().UTC(),
	}

	err := handler.store.CreateReviewSchedule(&testReview)
	testutils.CheckErr(t, err, "Failed to create test review")
	nextReview := time.Now().UTC().Add(24 * time.Hour)

	testData := map[string]interface{}{
		"review_id": testReview.ID,
		"next_review_at": nextReview.Format(time.RFC3339),
		"times_reviewed": testReview.TimesReviewed + 1,
		"interval_days": 3,
	}

	jsonData, err := json.Marshal(testData)
	testutils.CheckErr(t, err, "Failed to marshal json data")

	req := httptest.NewRequest("PUT", "/reviews/update", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.UpdateReviewSchedule(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	
	// check if the review is updated
	updatedReview, err := handler.store.GetReviewByID(testReview.ID)
	testutils.CheckErr(t, err, "Failed to get review")

	if updatedReview.IntervalDays != 3 {
		t.Errorf("Expected interval_days to be 3, got %d", updatedReview.IntervalDays)
	}
	
	if updatedReview.TimesReviewed != 1 {
		t.Errorf("Expected times_reviewed to be 1, got %d", updatedReview.TimesReviewed)
	}
	
	// Check if next_review_at was updated to approximately the expected time
	// Using a small delta to account for processing time differences
	expectedTime := nextReview.Unix()
	actualTime := updatedReview.NextReviewAt.UTC().Unix()
	if abs(expectedTime - actualTime) > 60 { // within 1 mins
		t.Errorf("Expected next_review_at to be close to %v, got %v", 
			nextReview, updatedReview.NextReviewAt)
	}
}

func TestCreateNewReview(t *testing.T) {
	handler, testDB, userID := setupReviewTest(t)
	defer testDB.Cleanup(t)

	// First create a submission (this should exist for the review to reference)
	submissionStore := models.NewSubmissionStore(testDB.DB)
	testSubmission := models.Submission{
		ID:          "test_submission_123",
		UserID:      userID,
		Title:       "Two Sum",
		TitleSlug:   "two-sum",
		SubmittedAt: time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}
	
	err := submissionStore.CreateSubmission(testSubmission)
	testutils.CheckErr(t, err, "Failed to create test submission")
	
	newReviewData := map[string]interface{}{
		"submission_id": testSubmission.ID,
		"interval_days": 1,
	}
	
	jsonData, err := json.Marshal(newReviewData)
	testutils.CheckErr(t, err, "Failed to marshal json data")
	
	req := httptest.NewRequest("POST", "/reviews/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	
	handler.CreateReview(rr, req)
	
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", rr.Code)
	}
	

	var response struct {
		ID int `json:"id"`
	}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to unmarshal response")
	
	review, err := handler.store.GetReviewByID(response.ID)
	testutils.CheckErr(t, err, "Failed to get created review")
	
	if review.SubmissionID != testSubmission.ID {
		t.Errorf("Expected submission ID %s, got %s", testSubmission.ID, review.SubmissionID)
	}
	
	if review.IntervalDays != 1 {
		t.Errorf("Expected interval days 1, got %d", review.IntervalDays)
	}
	
	// Verify next review date is in the future
	if !review.NextReviewAt.After(time.Now().UTC()) {
		t.Errorf("Next review date should be in the future")
	}
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}