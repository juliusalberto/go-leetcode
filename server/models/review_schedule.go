package models

import (
    "database/sql"
    "fmt"
    "time"
)

type ReviewSchedule struct {
    ID             int       `json:"id"`
    SubmissionID   string    `json:"submission_id"`
    NextReviewAt   time.Time `json:"next_review_at"`
    CreatedAt      time.Time `json:"created_at"`
    
    // FSRS fields
    Stability      float64   `json:"stability"`
    Difficulty     float64   `json:"difficulty"`
    ElapsedDays    uint64    `json:"elapsed_days"`
    ScheduledDays  uint64    `json:"scheduled_days"`
    Reps           uint64    `json:"reps"`
    Lapses         uint64    `json:"lapses"`
    State          int       `json:"state"` // 0=New, 1=Learning, 2=Review, 3=Relearning
    LastReview     time.Time `json:"last_review"`
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
        (submission_id, next_review_at, created_at, stability, difficulty, 
         elapsed_days, scheduled_days, reps, lapses, state, last_review)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id
    `
    
    err := s.db.QueryRow(
        query,
        review.SubmissionID,
        review.NextReviewAt,
        review.CreatedAt,
        review.Stability,
        review.Difficulty,
        review.ElapsedDays,
        review.ScheduledDays,
        review.Reps,
        review.Lapses,
        review.State,
        review.LastReview,
    ).Scan(&review.ID)

    if err != nil {
        return fmt.Errorf("error creating review schedule: %v", err)
    }

    return nil
}

func (s *ReviewScheduleStore) UpdateReviewSchedule(review *ReviewSchedule) error {
    query := `
        UPDATE review_schedules
        SET next_review_at = $1, stability = $2, difficulty = $3,
            elapsed_days = $4, scheduled_days = $5, reps = $6, 
            lapses = $7, state = $8, last_review = $9
        WHERE id = $10
    `

    result, err := s.db.Exec(
        query,
        review.NextReviewAt,
        review.Stability,
        review.Difficulty,
        review.ElapsedDays,
        review.ScheduledDays,
        review.Reps,
        review.Lapses,
        review.State,
        review.LastReview,
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

func (s *ReviewScheduleStore) GetReviewsBySubmissionID(submissionID string) ([]ReviewSchedule, error) {
    query := `
        SELECT id, submission_id, next_review_at, created_at, 
               stability, difficulty, elapsed_days, scheduled_days, 
               reps, lapses, state, last_review
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
        var lastReview sql.NullTime
        
        if err := rows.Scan(
            &review.ID,
            &review.SubmissionID,
            &review.NextReviewAt,
            &review.CreatedAt,
            &review.Stability,
            &review.Difficulty,
            &review.ElapsedDays,
            &review.ScheduledDays,
            &review.Reps,
            &review.Lapses,
            &review.State,
            &lastReview,
        ); err != nil {
            return nil, fmt.Errorf("error scanning review: %v", err)
        }
        
        if lastReview.Valid {
            review.LastReview = lastReview.Time
        }
        
        reviews = append(reviews, review)
    }

    return reviews, nil
}

func (s *ReviewScheduleStore) GetUpcomingReviews(userID int) ([]ReviewSchedule, error) {
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at,
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND r.next_review_at >= NOW()
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
        var lastReview sql.NullTime
        
        if err := rows.Scan(
            &review.ID,
            &review.SubmissionID,
            &review.NextReviewAt,
            &review.CreatedAt,
            &review.Stability,
            &review.Difficulty,
            &review.ElapsedDays,
            &review.ScheduledDays,
            &review.Reps,
            &review.Lapses,
            &review.State,
            &lastReview,
        ); err != nil {
            return nil, fmt.Errorf("error scanning review: %v", err)
        }
        
        if lastReview.Valid {
            review.LastReview = lastReview.Time
        }
        
        reviews = append(reviews, review)
    }

    return reviews, nil
}

func (s *ReviewScheduleStore) GetReviewsByUserID(userID int) ([]ReviewSchedule, error) {
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at,
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1
        ORDER BY r.created_at
    `

    rows, err := s.db.Query(query, userID)
    if err != nil {
        return nil, fmt.Errorf("error fetching reviews by user: %v", err)
    }
    defer rows.Close()

    var reviews []ReviewSchedule
    for rows.Next() {
        var review ReviewSchedule
        var lastReview sql.NullTime
        
        if err := rows.Scan(
            &review.ID,
            &review.SubmissionID,
            &review.NextReviewAt,
            &review.CreatedAt,
            &review.Stability,
            &review.Difficulty,
            &review.ElapsedDays,
            &review.ScheduledDays,
            &review.Reps,
            &review.Lapses,
            &review.State,
            &lastReview,
        ); err != nil {
            return nil, fmt.Errorf("error scanning review: %v", err)
        }
        
        if lastReview.Valid {
            review.LastReview = lastReview.Time
        }
        
        reviews = append(reviews, review)
    }

    return reviews, nil
}

func (s *ReviewScheduleStore) GetReviewByID(reviewID int) (ReviewSchedule, error) {
    query := `
        SELECT id, submission_id, next_review_at, created_at,
               stability, difficulty, elapsed_days, scheduled_days,
               reps, lapses, state, last_review
        FROM review_schedules
        WHERE id = $1
    `

    var review ReviewSchedule
    var lastReview sql.NullTime
    
    err := s.db.QueryRow(query, reviewID).Scan(
        &review.ID,
        &review.SubmissionID,
        &review.NextReviewAt,
        &review.CreatedAt,
        &review.Stability,
        &review.Difficulty,
        &review.ElapsedDays,
        &review.ScheduledDays,
        &review.Reps,
        &review.Lapses,
        &review.State,
        &lastReview,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return ReviewSchedule{}, fmt.Errorf("review with ID %d not found", reviewID)
        }
        return ReviewSchedule{}, fmt.Errorf("error fetching review: %v", err)
    }
    
    if lastReview.Valid {
        review.LastReview = lastReview.Time
    }

    return review, nil
}