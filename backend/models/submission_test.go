package models

import (
	"go-leetcode/backend/internal/database"
	"testing"
	"time"
)

func setUpData(t *testing.T, ss *SubmissionStore) {
	ss.db.Exec("INSERT INTO users (id, username, created_at) VALUES ($1, $2, $3)", 1, "test_user", time.Now())

	sub := Submission{
		ID:          "123",
		UserID:      1,
		Title:       "Two Sum",
		TitleSlug:   "two-sum",
		SubmittedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	err := ss.CreateSubmission(sub)
	if err != nil {
		t.Errorf("Failed to create submission: %v", err)
	}

} 

func TestGetSubmissionByID(t *testing.T) {
	testDB := database.SetupTestDB(t)
	defer testDB.Cleanup(t)

	submission_store := NewSubmissionStore(testDB.DB)
	setUpData(t, submission_store)

	received_sub, err := submission_store.GetSubmissionByID("123")

	if err != nil {
		t.Errorf("Failed to get submission: %v", err)
	}

	if received_sub.TitleSlug != "two-sum" {
		t.Error("The submission received does not match the submission sent.")
	}
}

func TestGetSubmissionByUserID(t *testing.T) {
	testDB := database.SetupTestDB(t)
	defer testDB.Cleanup(t)

	ss := NewSubmissionStore(testDB.DB)
	setUpData(t, ss)

	sub := Submission{
		ID:          "124",
		UserID:      1,
		Title:       "3Sum",
		TitleSlug:   "3-sum",
		SubmittedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	err := ss.CreateSubmission(sub)
	checkErr(t, err, "Failed to create submission")

	submissions, err := ss.GetSubmissionsByUserID(1)
	checkErr(t, err, "Failed to get submissions by user ID")

	if len(submissions) != 2 {
		t.Fatalf("Expected 2 submissions, got %d", len(submissions))
	}

	expectedTitleSlugs := map[string]bool {
		"two-sum": true,
		"3-sum": true,
	}

	for _, sub := range submissions {
		if _, exists := expectedTitleSlugs[sub.TitleSlug]; !exists {
			t.Errorf("Unexpected submission found: %v", sub.TitleSlug)
		}
	}
}

func TestCheckSubmissionExist(t *testing.T) {
	testDB := database.SetupTestDB(t)
	defer testDB.Cleanup(t)

	ss := NewSubmissionStore(testDB.DB)
	setUpData(t, ss)

	status, err := ss.CheckSubmissionExists("123", 1)
	checkErr(t, err, "There is an error in getting the submission")

	if status != true {
		t.Errorf("Submission not found")
	}
}

func checkErr(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}
