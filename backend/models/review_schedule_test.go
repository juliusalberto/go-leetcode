package models

import (
   "testing"
   "time"
   "go-leetcode/backend/internal/database"
   "go-leetcode/backend/internal/testutils"
)

func setupTestReview(t *testing.T) (*ReviewScheduleStore, *database.TestDB, ReviewSchedule) {
    testDB := database.SetupTestDB(t)
    
    // Create user first
    userStore := NewUserStore(testDB.DB)
    testUser := User{
        Username: "testuser",
        LeetcodeUsername: "leetcode_testuser",
        CreatedAt: time.Now(),
    }
    err := userStore.CreateUser(&testUser)
    testutils.CheckErr(t, err, "Failed to create test user")

    // Create submission
    submissionStore := NewSubmissionStore(testDB.DB)
    testSubmission := Submission{
        ID: "test_submission_id",
        UserID: testUser.ID,
        Title: "Test Problem",
        TitleSlug: "test-problem",
        SubmittedAt: time.Now(),
        CreatedAt: time.Now(),
    }
    err = submissionStore.CreateSubmission(testSubmission)
    testutils.CheckErr(t, err, "Failed to create test submission")

    // Finally create review
    store := NewReviewScheduleStore(testDB.DB)
    testReview := ReviewSchedule{
        SubmissionID:  testSubmission.ID,
        NextReviewAt:  time.Now().Add(24 * time.Hour),
        IntervalDays:  1,
        TimesReviewed: 0,
        CreatedAt:     time.Now(),
    }
    err = store.CreateReviewSchedule(&testReview)
    testutils.CheckErr(t, err, "Failed to create test review")

    return store, testDB, testReview
}

func TestCreateAndGetReviewSchedule(t *testing.T) {
   store, testDB, review := setupTestReview(t)
   defer testDB.Cleanup(t)

   // Test Get
   reviews, err := store.GetReviewsBySubmissionID(review.SubmissionID)
   testutils.CheckErr(t, err, "Failed to get reviews")
   if len(reviews) != 1 {
       t.Errorf("Expected 1 review, got %d", len(reviews))
   }
}

func TestUpdateReviewSchedule(t *testing.T) {
   store, testDB, review := setupTestReview(t)
   defer testDB.Cleanup(t)

   // Update review
   review.IntervalDays = 3
   review.TimesReviewed = 1
   err := store.UpdateReviewSchedule(&review)
   testutils.CheckErr(t, err, "Failed to update review")

   // Verify update
   reviews, err := store.GetReviewsBySubmissionID(review.SubmissionID)
   testutils.CheckErr(t, err, "Failed to get updated review")
   if reviews[0].IntervalDays != 3 || reviews[0].TimesReviewed != 1 {
       t.Error("Review not updated correctly")
   }
}