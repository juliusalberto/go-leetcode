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

	// Create test review with FSRS fields
	testReview := models.ReviewSchedule{
		SubmissionID:  "1",
		NextReviewAt:  time.Now().Add(24 * time.Hour),
		CreatedAt:     time.Now(),
		Stability:     3.0,
		Difficulty:    5.0,
		ElapsedDays:   0,
		ScheduledDays: 1,
		Reps:          1,
		Lapses:        0,
		State:         2, // Review state
		LastReview:    time.Now(),
	}

	err := handler.store.CreateReviewSchedule(&testReview)
	testutils.CheckErr(t, err, "Failed to create test review")

	url := fmt.Sprintf("/reviews/upcoming?user_id=%d", userID)

	req := httptest.NewRequest("GET", url, nil)
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

	// Create test review with FSRS fields
	testReview := models.ReviewSchedule{
		SubmissionID:  "1",
		NextReviewAt:  time.Now().UTC(),
		CreatedAt:     time.Now().UTC(),
		Stability:     3.0,
		Difficulty:    5.0,
		ElapsedDays:   0,
		ScheduledDays: 1,
		Reps:          1,
		Lapses:        0,
		State:         2, // Review state
		LastReview:    time.Now().UTC(),
	}

	err := handler.store.CreateReviewSchedule(&testReview)
	testutils.CheckErr(t, err, "Failed to create test review")

	// Test updating with a "Good" rating (3)
	testData := map[string]interface{}{
		"review_id": testReview.ID,
		"rating":    3, // Good
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
	
	// Parse the response
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	testutils.CheckErr(t, err, "Failed to unmarshal response")
	
	// Verify response contains expected fields
	if _, ok := response["success"]; !ok || response["success"] != true {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}
	
	if _, ok := response["next_review_at"]; !ok {
		t.Errorf("Response missing next_review_at field")
	}
	
	if _, ok := response["days_until_review"]; !ok {
		t.Errorf("Response missing days_until_review field")
	}
	
	// Check if the review was updated in the database
	updatedReview, err := handler.store.GetReviewByID(testReview.ID)
	testutils.CheckErr(t, err, "Failed to get review")

	// Verify FSRS fields were updated
	if updatedReview.Reps <= testReview.Reps {
		t.Errorf("Expected reps to increase, got %d", updatedReview.Reps)
	}
	
	if updatedReview.LastReview.Before(testReview.LastReview) {
		t.Errorf("Expected last_review to be updated")
	}
	
	// Check that next review date is in the future
	if !updatedReview.NextReviewAt.After(time.Now()) {
		t.Errorf("Expected next_review_at to be in the future")
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
	
	// Verify FSRS fields are set
	if review.Stability <= 0 {
		t.Errorf("Expected stability > 0, got %f", review.Stability)
	}
	
	if review.Reps != 1 {
		t.Errorf("Expected reps to be 1, got %d", review.Reps)
	}
	
	if review.State < 0 || review.State > 3 {
		t.Errorf("Expected valid state (0-3), got %d", review.State)
	}
	
	// Verify next review date is in the future
	if !review.NextReviewAt.After(time.Now().UTC()) {
		t.Errorf("Next review date should be in the future")
	}
}
