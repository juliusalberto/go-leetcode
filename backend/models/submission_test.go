package models

import (
	"go-leetcode/backend/internal/database"
	"testing"
	"time"
)

func TestGetSubmissionByID(t *testing.T) {
	testDB := database.SetupTestDB(t)
	defer testDB.Cleanup(t)

	submission_store := NewSubmissionStore(testDB.DB)

	sub := Submission{
		ID:          "123",
		UserID:      1,
		Title:       "Two Sum",
		TitleSlug:   "two-sum",
		SubmittedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	err := submission_store.CreateSubmission(sub)
	if err != nil {
		t.Errorf("Failed to create submission: %v", err)
	}

	received_sub, err := submission_store.GetSubmissionByID("123")

	if err != nil {
		t.Errorf("Failed to get submission: %v", err)
	}

	if received_sub.TitleSlug != sub.TitleSlug {
		t.Error("The submission received does not match the submission sent.")
	}

}
