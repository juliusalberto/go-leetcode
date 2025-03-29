package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Submission struct {
	ID 			string			`json:"id"`
	UserID		uuid.UUID		`json:"user_id"`
	Title		string 			`json:"title"`
	TitleSlug 	string			`json:"title_slug"`
	SubmittedAt time.Time		`json:"submitted_at"`
	CreatedAt	time.Time		`json:"created_at"`
}

type SubmissionStore struct {
	db *sql.DB
}

func NewSubmissionStore(db *sql.DB) *SubmissionStore {
	return &SubmissionStore{db: db}
}

func (s *SubmissionStore) CreateSubmission(sub Submission) error {
	query := `
		INSERT INTO submissions
		(id, user_id, title, title_slug, submitted_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := s.db.Exec(query, sub.ID, sub.UserID, sub.Title, sub.TitleSlug, sub.SubmittedAt, sub.CreatedAt)

    if err != nil {
        return fmt.Errorf("error creating submission: %v", err)
    }

	return nil
}

func (s *SubmissionStore) GetSubmissionByID(id string)(Submission, error) {
	var sub Submission

	query := `
		SELECT id, user_id, title, title_slug, submitted_at, created_at 
		FROM submissions WHERE ID = $1
	`

	err := s.db.QueryRow(query, id).Scan(
		&sub.ID,
		&sub.UserID,
		&sub.Title,
		&sub.TitleSlug,
		&sub.SubmittedAt,
		&sub.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Submission{}, fmt.Errorf("submission with ID %s not found", id)
		}

		return Submission{}, fmt.Errorf("error fetching submission: %v", err)
	}

	return sub, nil 
}

func (s *SubmissionStore) GetSubmissionsByUserID(userID uuid.UUID)([]Submission, error) {
	query := `
		SELECT id, user_id, title, title_slug, submitted_at, created_at
		FROM submissions WHERE user_id = $1
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	var submissions []Submission

	for rows.Next() {
		var sub Submission
		if err := rows.Scan(&sub.ID, &sub.UserID, &sub.Title, &sub.TitleSlug, &sub.SubmittedAt, &sub.CreatedAt); err != nil {
			return submissions, err
		}

		submissions = append(submissions, sub)
	}

	if err = rows.Err(); err != nil {
		return submissions, err
	}


	return submissions, nil
}

func (s *SubmissionStore) CheckSubmissionExists(submissionID string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM submissions
			WHERE id = $1
		)
	`

	var exists bool 
	err := s.db.QueryRow(query, submissionID).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error checking submission existence : %v", err)
	}

	return exists, nil
}