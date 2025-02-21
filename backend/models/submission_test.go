package models

import (
	"testing"
	"time"
	"go-leetcode/backend/internal/database"
)


func TestCreateSubmission(t *testing.T) {
	testDB := database.SetupTestDB(t)
	defer testDB.Cleanup(t)

	sub := Submission{
        ID: "123",
        UserID: 1,
        Title: "Two Sum",
        TitleSlug: "two-sum",
        SubmittedAt: time.Now(),
        CreatedAt: time.Now(),
    }
    
    err := CreateSubmission(testDB.DB, sub)
    if err != nil {
        t.Errorf("Failed to create submission: %v", err)
    }
}