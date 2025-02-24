package handlers

import (
	"encoding/json"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupReviewTest(t *testing.T)(*ReviewHandler, *database.TestDB) {
	testDB := database.SetupTestDB(t)
	reviewStore := models.NewReviewScheduleStore(testDB.DB)
	handler := NewReviewHandler(reviewStore)
	return handler, testDB
}

func TestGetUpcomingReviewHandler(t *testing.T) {
	handler, testDB := setupReviewTest(t)
	defer testDB.Cleanup(t)

	testReview := models.ReviewSchedule{
		SubmissionID: "test_submission_id",
		NextReviewAt: time.Now(),
		IntervalDays: 1,
		TimesReviewed: 0,
		CreatedAt: time.Now(),
	}

	err := handler.store.CreateReviewSchedule(&testReview)
	testutils.CheckErr(t, err, "Failed to create test review")

	req := httptest.NewRequest("GET", "/reviews/upcoming", nil)
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