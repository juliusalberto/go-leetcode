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

    // Finally create review with FSRS fields
    now := time.Now()
    store := NewReviewScheduleStore(testDB.DB)
    testReview := ReviewSchedule{
        SubmissionID:  testSubmission.ID,
        NextReviewAt:  now.Add(24 * time.Hour),
        CreatedAt:     now,
        // FSRS specific fields
        Stability:     2.5,
        Difficulty:    0.3,
        ElapsedDays:   0,
        ScheduledDays: 1,
        Reps:          0,
        Lapses:        0,
        State:         0, // New state
        LastReview:    time.Time{}, // Zero value for new card
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
   
   // Verify FSRS fields
   if reviews[0].Stability < 0.1 {
       t.Errorf("Expected positive stability, got %f", reviews[0].Stability)
   }
   
   if reviews[0].Difficulty < 0 || reviews[0].Difficulty > 1 {
       t.Errorf("Expected difficulty between 0-1, got %f", reviews[0].Difficulty)
   }
}

func TestUpdateReviewSchedule(t *testing.T) {
   store, testDB, review := setupTestReview(t)
   defer testDB.Cleanup(t)

   // Update review with FSRS values
   now := time.Now().UTC()
   review.Stability = 5.0
   review.Difficulty = 0.4
   review.ScheduledDays = 3
   review.Reps = 1
   review.State = 2 // Review state
   review.LastReview = now
   review.NextReviewAt = now.AddDate(0, 0, 3) // 3 days later
   
   err := store.UpdateReviewSchedule(&review)
   testutils.CheckErr(t, err, "Failed to update review")

   // Verify update
   updatedReview, err := store.GetReviewByID(review.ID)
   testutils.CheckErr(t, err, "Failed to get updated review")
   
   if updatedReview.Stability != 5.0 {
       t.Errorf("Expected stability 5.0, got %f", updatedReview.Stability)
   }
   
   if updatedReview.Difficulty != 0.4 {
       t.Errorf("Expected difficulty 0.4, got %f", updatedReview.Difficulty)
   }
   
   if updatedReview.ScheduledDays != 3 {
       t.Errorf("Expected scheduled days 3, got %d", updatedReview.ScheduledDays)
   }
   
   if updatedReview.Reps != 1 {
       t.Errorf("Expected reps 1, got %d", updatedReview.Reps)
   }
   
   if updatedReview.State != 2 {
       t.Errorf("Expected state 2 (review), got %d", updatedReview.State)
   }
   
   // Check that LastReview was recorded
   if updatedReview.LastReview.IsZero() {
       t.Error("Expected LastReview to be set, got zero time")
   }
}

func TestFSRSWorkflow(t *testing.T) {
   store, testDB, review := setupTestReview(t)
   defer testDB.Cleanup(t)
   
   // Simulate a complete FSRS review cycle
   
   // First review (Good response)
   now := time.Now().UTC()
   review.State = 1 // Learning -> Review
   review.Stability = 3.0
   review.Difficulty = 0.3
   review.ScheduledDays = 3
   review.Reps = 1
   review.Lapses = 0
   review.LastReview = now
   review.NextReviewAt = now.AddDate(0, 0, 3)
   
   err := store.UpdateReviewSchedule(&review)
   testutils.CheckErr(t, err, "Failed to update review after first response")
   
   // Second review (Hard response - shorter interval)
   secondReview, err := store.GetReviewByID(review.ID)
   testutils.CheckErr(t, err, "Failed to get review for second update")
   
   laterTime := now.AddDate(0, 0, 3).Add(1 * time.Hour) // 3 days + 1 hour later
   
   secondReview.State = 2 // Stays in Review state
   secondReview.Stability = 2.0 // Reduced stability due to Hard rating
   secondReview.Difficulty = 0.4 // Increased difficulty
   secondReview.ScheduledDays = 2 // Shorter interval
   secondReview.Reps = 2
   secondReview.Lapses = 0
   secondReview.LastReview = laterTime
   secondReview.NextReviewAt = laterTime.AddDate(0, 0, 2)
   
   err = store.UpdateReviewSchedule(&secondReview)
   testutils.CheckErr(t, err, "Failed to update review after second response")
   
   // Verify the final state
   finalReview, err := store.GetReviewByID(review.ID)
   testutils.CheckErr(t, err, "Failed to get final review state")
   
   if finalReview.Difficulty <= review.Difficulty {
       t.Errorf("Expected increased difficulty after Hard rating, got %f -> %f", 
           review.Difficulty, finalReview.Difficulty)
   }
   
   if finalReview.Reps != 2 {
       t.Errorf("Expected 2 repetitions, got %d", finalReview.Reps)
   }
}