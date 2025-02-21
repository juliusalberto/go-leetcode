package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Submission struct {
	ID 			string
	UserID		int
	Title		string 
	TitleSlug 	string
	SubmittedAt time.Time
	CreatedAt	time.Time
}

func CreateSubmission(db *sql.DB, sub Submission) error {
	query := `
		INSERT INTO submissions
		(id, user_id, title, title_slug, submitted_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := db.Exec(query, sub.ID, sub.UserID, sub.Title, sub.TitleSlug, sub.SubmittedAt, sub.CreatedAt)

    if err != nil {
        return fmt.Errorf("error creating submission: %v", err)
    }

	return nil
}