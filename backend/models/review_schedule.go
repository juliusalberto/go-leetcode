package models

import (
    "database/sql"
    "fmt"
    "time"
)

type ReviewSchedule struct {
    ID             int
    SubmissionID   string
    NextReviewAt   time.Time
    IntervalDays   int
    TimesReviewed  int
    CreatedAt      time.Time
}

type ReviewScheduleStore struct {
    db *sql.DB
}

func NewReviewScheduleStore(db *sql.DB) *ReviewScheduleStore {
    return &ReviewScheduleStore{db: db}
}

func (s *ReviewScheduleStore) CreateReviewSchedule(review *ReviewSchedule) error {
    query := `
        INSERT INTO review_schedules
        (submission_id, next_review_at, interval_days, times_reviewed, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
    
    err := s.db.QueryRow(
        query,
        review.SubmissionID,
        review.NextReviewAt,
        review.IntervalDays,
        review.TimesReviewed,
        review.CreatedAt,
    ).Scan(&review.ID)

    if err != nil {
        return fmt.Errorf("error creating review schedule: %v", err)
    }

    return nil
}

func (s *ReviewScheduleStore) GetReviewsBySubmissionID(submissionID string) ([]ReviewSchedule, error) {
    query := `
        SELECT id, submission_id, next_review_at, interval_days, times_reviewed, created_at
        FROM review_schedules 
        WHERE submission_id = $1
        ORDER BY next_review_at
    `

    rows, err := s.db.Query(query, submissionID)
    if err != nil {
        return nil, fmt.Errorf("error fetching reviews: %v", err)
    }
    defer rows.Close()

    var reviews []ReviewSchedule
    for rows.Next() {
        var review ReviewSchedule
        if err := rows.Scan(
            &review.ID,
            &review.SubmissionID,
            &review.NextReviewAt,
            &review.IntervalDays,
            &review.TimesReviewed,
            &review.CreatedAt,
        ); err != nil {
            return nil, fmt.Errorf("error scanning review: %v", err)
        }
        reviews = append(reviews, review)
    }

    return reviews, nil
}

func (s *ReviewScheduleStore) GetUpcomingReviews(userID int) ([]ReviewSchedule, error) {
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.interval_days, r.times_reviewed, r.created_at
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND r.next_review_at <= NOW()
        ORDER BY r.next_review_at
    `

    rows, err := s.db.Query(query, userID)
    if err != nil {
        return nil, fmt.Errorf("error fetching upcoming reviews: %v", err)
    }
    defer rows.Close()

    var reviews []ReviewSchedule
    for rows.Next() {
        var review ReviewSchedule
        if err := rows.Scan(
            &review.ID,
            &review.SubmissionID,
            &review.NextReviewAt,
            &review.IntervalDays,
            &review.TimesReviewed,
            &review.CreatedAt,
        ); err != nil {
            return nil, fmt.Errorf("error scanning review: %v", err)
        }
        reviews = append(reviews, review)
    }

    return reviews, nil
}

func (s *ReviewScheduleStore) UpdateReviewSchedule(review *ReviewSchedule) error {
    query := `
        UPDATE review_schedules
        SET next_review_at = $1, interval_days = $2, times_reviewed = $3
        WHERE id = $4
    `

    result, err := s.db.Exec(
        query,
        review.NextReviewAt,
        review.IntervalDays,
        review.TimesReviewed,
        review.ID,
    )
    if err != nil {
        return fmt.Errorf("error updating review schedule: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking update result: %v", err)
    }
    if rows == 0 {
        return fmt.Errorf("review schedule with ID %d not found", review.ID)
    }

    return nil
}