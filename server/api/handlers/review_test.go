package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
)

func setupReviewTest(t *testing.T) (*ReviewHandler, *database.TestDB, uuid.UUID) {
	testDB := database.SetupTestDB(t)
	reviewStore := models.NewReviewScheduleStore(testDB.DB)
	submissionStore := models.NewSubmissionStore(testDB.DB)
	userStore := models.NewUserStore(testDB.DB)
	handler := NewReviewHandler(reviewStore, submissionStore)

	// create user first
	testUser := models.User{
		Username:         "testuser",
		LeetcodeUsername: "test_leetcode",
	}

	userStore.CreateUser(&testUser)

	for i := 0; i < 5; i++ {
		testSub := models.Submission{
			ID:          strconv.Itoa(i + 1),
			UserID:      testUser.ID,
			Title:       fmt.Sprintf("%d Sum", i+1),
			TitleSlug:   fmt.Sprintf("%d-sum", i+1),
			SubmittedAt: time.Now(),
			CreatedAt:   time.Now(),
		}

		if err := submissionStore.CreateSubmission(testSub); err != nil {
			t.Errorf("Failed to create submission")
		}
	}

	return handler, testDB, testUser.ID
}

func TestProcessSubmission(t *testing.T) {
	handler, testDB, userID := setupReviewTest(t)
	defer testDB.Cleanup(t)
	userUUIDKey := middleware.UserUUIDKey
	
	// Note: We're not using userID directly in the request anymore as it should be
	// retrieved from the authentication context in the handler

	// Test case 1: Create a new submission and review with struct
	requestStruct := struct{
		IsInternal          bool   `json:"is_internal"`
		LeetcodeSubmissionID string `json:"leetcode_submission_id"`
		Title               string `json:"title"`
		TitleSlug           string `json:"title_slug"`
		SubmittedAt         string `json:"submitted_at"`
	}{
		IsInternal:          true,
		LeetcodeSubmissionID: "",
		Title:               "Process Test Problem",
		TitleSlug:           "two-sum",
		SubmittedAt:         time.Now().UTC().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(requestStruct)
	testutils.CheckErr(t, err, "Failed to marshal test submission")

	req := httptest.NewRequest("POST", "/api/reviews/process-submission", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), userUUIDKey, userID)
	req = req.WithContext(ctx)

	// Call the handler
	handler.ProcessSubmission(rr, req)

	// Check response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Get response data
	respData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal response data")

	var respObj map[string]interface{}
	err = json.Unmarshal(respData, &respObj)
	testutils.CheckErr(t, err, "Failed to unmarshal response data")

	// Verify response contains expected fields
	if success, ok := respObj["success"].(bool); !ok || !success {
		t.Errorf("Expected success: true, got %v", respObj["success"])
	}

	if _, ok := respObj["submission_id"].(string); !ok {
		t.Errorf("Expected submission_id to be present and a string")
	}

	if _, ok := respObj["next_review_at"]; !ok {
		t.Errorf("Response missing next_review_at field")
	}

	if _, ok := respObj["days_until_review"]; !ok {
		t.Errorf("Response missing days_until_review field")
	}

	if _, ok := respObj["is_due"]; !ok {
		t.Errorf("Response missing is_due field")
	}

	// Test case 2: Process same problem with a new submission ID
	requestStruct2 := struct{
		IsInternal          bool   `json:"is_internal"`
		LeetcodeSubmissionID string `json:"leetcode_submission_id"`
		Title               string `json:"title"`
		TitleSlug           string `json:"title_slug"`
		SubmittedAt         string `json:"submitted_at"`
	}{
		IsInternal:          true,
		LeetcodeSubmissionID: "",
		Title:               "Process Test Problem",
		TitleSlug:           "two-sum",
		SubmittedAt:         time.Now().UTC().Add(24 * time.Hour).Format(time.RFC3339),
	}

	jsonData2, err := json.Marshal(requestStruct2)
	testutils.CheckErr(t, err, "Failed to marshal second test submission")

	req2 := httptest.NewRequest("POST", "/api/reviews/process-submission", bytes.NewBuffer(jsonData2))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()
	req2 = req2.WithContext(ctx)

	handler.ProcessSubmission(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("Expected status 200 for second request, got %d", rr2.Code)
	}

	// Extract the data to verify it processed correctly
	var resp2 response.Response
	err = json.Unmarshal(rr2.Body.Bytes(), &resp2)
	testutils.CheckErr(t, err, "Failed to unmarshal second response")

	respData2, err := json.Marshal(resp2.Data)
	testutils.CheckErr(t, err, "Failed to marshal second response data")

	var respObj2 map[string]interface{}
	err = json.Unmarshal(respData2, &respObj2)
	testutils.CheckErr(t, err, "Failed to unmarshal second response data")

	// Verify the response contains expected fields
	if _, ok := respObj2["submission_id"].(string); !ok {
		t.Errorf("Expected submission_id to be present and a string")
	}

	// Test case 3: Missing required fields
	badRequestStruct := struct{
		// Missing TitleSlug
		IsInternal          bool   `json:"is_internal"`
		LeetcodeSubmissionID string `json:"leetcode_submission_id"`
		Title               string `json:"title"`
		TitleSlug           string `json:"title_slug"`
		SubmittedAt         string `json:"submitted_at"`
	}{
		IsInternal:          true,
		LeetcodeSubmissionID: "",
		Title:               "Bad Test Submission",
		SubmittedAt:         time.Now().UTC().Format(time.RFC3339),
	}

	badJsonData, err := json.Marshal(badRequestStruct)
	testutils.CheckErr(t, err, "Failed to marshal bad test submission")

	badReq := httptest.NewRequest("POST", "/api/reviews/process-submission", bytes.NewBuffer(badJsonData))
	badReq.Header.Set("Content-Type", "application/json")
	badRr := httptest.NewRecorder()
	badReq = badReq.WithContext(ctx)

	handler.ProcessSubmission(badRr, badReq)

	// Should get a validation error
	if badRr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for bad request, got %d", badRr.Code)
	}
}

func TestUpdateOrCreateReview(t *testing.T) {
	handler, testDB, userID := setupReviewTest(t)
	defer testDB.Cleanup(t)

	// Create a submission to test with
	testSubmission := models.Submission{
		ID:         "test_submission_999",
		UserID:     userID,
		Title:      "Test Problem for Update",
		TitleSlug:  "test-problem-update",
		SubmittedAt: time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}
	
	// Save submission to database
	submissionStore := models.NewSubmissionStore(testDB.DB)
	err := submissionStore.CreateSubmission(testSubmission)
	testutils.CheckErr(t, err, "Failed to create test submission for handler test")

	// Create the request with the submission
	jsonData, err := json.Marshal(testSubmission)
	testutils.CheckErr(t, err, "Failed to marshal test submission")

	req := httptest.NewRequest("POST", "/reviews/update-or-create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call the handler
	handler.UpdateOrCreateReview(rr, req)

	// Check response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Get response data
	respData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal response data")

	var respObj map[string]interface{}
	err = json.Unmarshal(respData, &respObj)
	testutils.CheckErr(t, err, "Failed to unmarshal response data")

	// Verify response contains expected fields
	if success, ok := respObj["success"].(bool); !ok || !success {
		t.Errorf("Expected success: true, got %v", respObj["success"])
	}

	if _, ok := respObj["next_review_at"]; !ok {
		t.Errorf("Response missing next_review_at field")
	}

	if _, ok := respObj["days_until_review"]; !ok {
		t.Errorf("Response missing days_until_review field")
	}

	// Test the handler again with the same title_slug to check the update case
	testSubmission2 := models.Submission{
		ID:         "test_submission_1000", // Different ID
		UserID:     userID,
		Title:      "Test Problem for Update",
		TitleSlug:  "test-problem-update", // Same title_slug
		SubmittedAt: time.Now().UTC().Add(24 * time.Hour), // Later submission
		CreatedAt:   time.Now().UTC(),
	}
	
	// Save second submission to database
	err = submissionStore.CreateSubmission(testSubmission2)
	testutils.CheckErr(t, err, "Failed to create second test submission for handler test")

	jsonData2, err := json.Marshal(testSubmission2)
	testutils.CheckErr(t, err, "Failed to marshal second test submission")

	req2 := httptest.NewRequest("POST", "/reviews/update-or-create", bytes.NewBuffer(jsonData2))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()

	handler.UpdateOrCreateReview(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("Expected status 200 for second request, got %d", rr2.Code)
	}

	// Verify in the database that we have a review for the second submission ID
	reviews, err := handler.store.GetReviewsBySubmissionID(testSubmission2.ID)
	testutils.CheckErr(t, err, "Failed to get reviews for second submission")

	if len(reviews) != 1 {
		t.Errorf("Expected 1 review for the second submission, got %d", len(reviews))
	}
}

func TestGetUpcomingReviewHandler(t *testing.T) {
	handler, testDB, userID := setupReviewTest(t)
	userUUIDKey := middleware.UserUUIDKey
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

	url := "/api/reviews?status=upcoming&page=1&per_page=10"

	req := httptest.NewRequest("GET", url, nil)
	ctx := context.WithValue(req.Context(), userUUIDKey, userID)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	handler.GetReviews(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Decode standardized response
	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to parse response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Check pagination data
	if resp.Meta.Pagination == nil {
		t.Errorf("Expected pagination data in response")
	} else {
		if resp.Meta.Pagination.Total != 1 {
			t.Errorf("Expected total count of 1, got %d", resp.Meta.Pagination.Total)
		}
		if resp.Meta.Pagination.Page != 1 {
			t.Errorf("Expected page 1, got %d", resp.Meta.Pagination.Page)
		}
		if resp.Meta.Pagination.PerPage != 10 {
			t.Errorf("Expected per_page 10, got %d", resp.Meta.Pagination.PerPage)
		}
	}

	// Extract review data from response
	reviewData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal review data")

	var reviews []models.ReviewSchedule
	err = json.Unmarshal(reviewData, &reviews)
	testutils.CheckErr(t, err, "Failed to unmarshal review data")

	if len(reviews) != 1 || reviews[0].SubmissionID != testReview.SubmissionID {
		t.Errorf("unexpected response: %+v", reviews)
	}
}

func TestGetDueReviewHandler(t *testing.T) {
	handler, testDB, userID := setupReviewTest(t)
	userUUIDKey := middleware.UserUUIDKey
	defer testDB.Cleanup(t)

	// Create a due test review with FSRS fields (nextReviewAt is in the past)
	dueReview := models.ReviewSchedule{
		SubmissionID:  "2",
		NextReviewAt:  time.Now().Add(-24 * time.Hour), // Due 1 day ago
		CreatedAt:     time.Now().Add(-48 * time.Hour), // Created 2 days ago
		Stability:     3.0,
		Difficulty:    5.0,
		ElapsedDays:   0,
		ScheduledDays: 1,
		Reps:          1,
		Lapses:        0,
		State:         2, // Review state
		LastReview:    time.Now().Add(-48 * time.Hour),
	}

	err := handler.store.CreateReviewSchedule(&dueReview)
	testutils.CheckErr(t, err, "Failed to create due test review")

	// Create a future test review with FSRS fields
	upcomingReview := models.ReviewSchedule{
		SubmissionID:  "3",
		NextReviewAt:  time.Now().Add(24 * time.Hour), // Due in 1 day
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

	err = handler.store.CreateReviewSchedule(&upcomingReview)
	testutils.CheckErr(t, err, "Failed to create upcoming test review")

	// Test 1: Get only due reviews
	url := fmt.Sprintf("/api/reviews?user_id=%s&status=due&page=1&per_page=10", userID.String())
	req := httptest.NewRequest("GET", url, nil)
	rr := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), userUUIDKey, userID)
	req = req.WithContext(ctx)

	handler.GetReviews(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Decode standardized response
	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to parse due reviews response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Check pagination data
	if resp.Meta.Pagination == nil {
		t.Errorf("Expected pagination data in response")
	} else {
		if resp.Meta.Pagination.Total != 1 {
			t.Errorf("Expected total count of 1, got %d", resp.Meta.Pagination.Total)
		}
	}

	// Extract review data from response
	reviewData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal due reviews data")

	var dueReviews []models.ReviewSchedule
	err = json.Unmarshal(reviewData, &dueReviews)
	testutils.CheckErr(t, err, "Failed to unmarshal due reviews data")

	// Should only contain the due review
	if len(dueReviews) != 1 || dueReviews[0].SubmissionID != dueReview.SubmissionID {
		t.Errorf("Expected 1 due review with ID %s, got %d reviews: %+v",
			dueReview.SubmissionID, len(dueReviews), dueReviews)
	}

	// Test 2: Get all reviews (both due and upcoming)
	url = "/api/reviews?page=1&per_page=10"
	req = httptest.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.GetReviews(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Decode standardized response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to parse all reviews response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Check pagination data
	if resp.Meta.Pagination == nil {
		t.Errorf("Expected pagination data in response for all reviews")
	} else {
		if resp.Meta.Pagination.Total != 2 {
			t.Errorf("Expected total count of 2, got %d", resp.Meta.Pagination.Total)
		}
	}

	// Extract review data from response
	reviewData, err = json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal all reviews data")

	var allReviews []models.ReviewSchedule
	err = json.Unmarshal(reviewData, &allReviews)
	testutils.CheckErr(t, err, "Failed to unmarshal all reviews data")

	// Should contain both reviews
	if len(allReviews) != 2 {
		t.Errorf("Expected 2 reviews, got %d", len(allReviews))
	}

	// Verify that both due and upcoming reviews are included
	var foundDue, foundUpcoming bool
	for _, review := range allReviews {
		if review.SubmissionID == dueReview.SubmissionID {
			foundDue = true
		} else if review.SubmissionID == upcomingReview.SubmissionID {
			foundUpcoming = true
		}
	}

	if !foundDue {
		t.Errorf("Due review not found in the combined result")
	}
	if !foundUpcoming {
		t.Errorf("Upcoming review not found in the combined result")
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

	// Decode standardized response
	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract update data from response
	updateData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal update data")

	var updateResponse map[string]interface{}
	err = json.Unmarshal(updateData, &updateResponse)
	testutils.CheckErr(t, err, "Failed to unmarshal update data")

	// Verify response contains expected fields
	if _, ok := updateResponse["success"]; !ok || updateResponse["success"] != true {
		t.Errorf("Expected success to be true, got %v", updateResponse["success"])
	}

	if _, ok := updateResponse["next_review_at"]; !ok {
		t.Errorf("Response missing next_review_at field")
	}

	if _, ok := updateResponse["days_until_review"]; !ok {
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

	// Decode standardized response
	var resp response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Extract data from response
	responseData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal response data")

	var responseObj struct {
		ID int `json:"id"`
	}
	err = json.Unmarshal(responseData, &responseObj)
	testutils.CheckErr(t, err, "Failed to unmarshal response data")

	review, err := handler.store.GetReviewByID(responseObj.ID)
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