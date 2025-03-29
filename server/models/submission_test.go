package models

import (
	"go-leetcode/backend/internal/database"
	"testing"
	"go-leetcode/backend/internal/testutils"
	"time"
	
	"github.com/google/uuid"
)

func setUpData(t *testing.T, ss *SubmissionStore) uuid.UUID {
	testUserID := uuid.New()
	ss.db.Exec("INSERT INTO users (id, username, leetcode_username, created_at) VALUES ($1, $2, $3, $4)", testUserID, "test_user", "test_user", time.Now())

	sub := Submission{
		ID:          "123",
		UserID:      testUserID,
		Title:       "Two Sum",
		TitleSlug:   "two-sum",
		SubmittedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	err := ss.CreateSubmission(sub)
	if err != nil {
		t.Errorf("Failed to create submission: %v", err)
	}

	return testUserID
}

func TestGetSubmissionByID(t *testing.T) {
	testDB := database.SetupTestDB(t)
	defer testDB.Cleanup(t)

	submission_store := NewSubmissionStore(testDB.DB)
	_ = setUpData(t, submission_store) // We don't need the UUID for this test

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
	testUserID := setUpData(t, ss)

	sub := Submission{
		ID:          "124",
		UserID:      testUserID,
		Title:       "3Sum",
		TitleSlug:   "3-sum",
		SubmittedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	err := ss.CreateSubmission(sub)
	testutils.CheckErr(t, err, "Failed to create submission")

	submissions, err := ss.GetSubmissionsByUserID(testUserID)
	testutils.CheckErr(t, err, "Failed to get submissions by user ID")

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
	_ = setUpData(t, ss) // We don't need the UUID for this test

	status, err := ss.CheckSubmissionExists("123")
	testutils.CheckErr(t, err, "There is an error in getting the submission")

	if status != true {
		t.Errorf("Submission not found")
	}
}

