package models

import (
    "database/sql"
    "fmt"
    "time"
    
    "github.com/open-spaced-repetition/go-fsrs/v3"
)

type ReviewSchedule struct {
    ID             int       `json:"id"`
    SubmissionID   string    `json:"submission_id"`
    Title          string    `json:"title,omitempty"`
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
        SET submission_id = $1, next_review_at = $2, stability = $3, difficulty = $4,
            elapsed_days = $5, scheduled_days = $6, reps = $7, 
            lapses = $8, state = $9, last_review = $10
        WHERE id = $11
    `

    result, err := s.db.Exec(
        query,
        review.SubmissionID,
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

func (s *ReviewScheduleStore) GetUpcomingReviews(userID int, limit, offset int) ([]ReviewSchedule, int, error) {
    // First, count total records for pagination
    countQuery := `
        SELECT COUNT(*)
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND r.next_review_at >= NOW()
    `
    var total int
    err := s.db.QueryRow(countQuery, userID).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("error counting upcoming reviews: %v", err)
    }

    // Then get paginated results
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at,
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review, s.title
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND r.next_review_at >= NOW()
        ORDER BY r.next_review_at
        LIMIT $2 OFFSET $3
    `

    rows, err := s.db.Query(query, userID, limit, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("error fetching upcoming reviews: %v", err)
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
            &review.Title,
        ); err != nil {
            return nil, 0, fmt.Errorf("error scanning review: %v", err)
        }
        
        if lastReview.Valid {
            review.LastReview = lastReview.Time
        }
        
        reviews = append(reviews, review)
    }

    return reviews, total, nil
}

func (s *ReviewScheduleStore) GetDueReviews(userID int, limit, offset int) ([]ReviewSchedule, int, error) {
    // First, count total records for pagination
    countQuery := `
        SELECT COUNT(*)
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND r.next_review_at <= NOW()
    `
    var total int
    err := s.db.QueryRow(countQuery, userID).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("error counting due reviews: %v", err)
    }

    // Then get paginated results
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at,
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review, s.title
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND r.next_review_at <= NOW()
        ORDER BY r.next_review_at
        LIMIT $2 OFFSET $3
    `

    rows, err := s.db.Query(query, userID, limit, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("error fetching due reviews: %v", err)
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
            &review.Title,
        ); err != nil {
            return nil, 0, fmt.Errorf("error scanning review: %v", err)
        }
        
        if lastReview.Valid {
            review.LastReview = lastReview.Time
        }
        
        reviews = append(reviews, review)
    }

    return reviews, total, nil
}

func (s *ReviewScheduleStore) GetReviewsByUserID(userID int) ([]ReviewSchedule, error) {
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at,
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review, s.title
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
            &review.Title,
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
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at,
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review, s.title
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
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
        &review.Title,
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

func (s *ReviewScheduleStore) GetReviewByTitleSlug(userID int, titleSlug string) (ReviewSchedule, error) {
    query := `
        SELECT r.id, r.submission_id, r.next_review_at, r.created_at, 
               r.stability, r.difficulty, r.elapsed_days, r.scheduled_days,
               r.reps, r.lapses, r.state, r.last_review, s.title
        FROM review_schedules r
        JOIN submissions s ON r.submission_id = s.id
        WHERE s.user_id = $1 AND s.title_slug = $2
        ORDER BY s.submitted_at DESC
        LIMIT 1
    `

    var review ReviewSchedule
    var lastReview sql.NullTime
    
    err := s.db.QueryRow(query, userID, titleSlug).Scan(
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
        &review.Title,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return ReviewSchedule{}, fmt.Errorf("no review found for title slug %s and user ID %d", titleSlug, userID)
        }
        return ReviewSchedule{}, fmt.Errorf("error fetching review by title slug: %v", err)
    }
    
    if lastReview.Valid {
        review.LastReview = lastReview.Time
    }

    return review, nil
}

// ConvertReviewScheduleToFSRS converts a ReviewSchedule to an FSRS Card
func ConvertReviewScheduleToFSRS(review *ReviewSchedule) fsrs.Card {
    return fsrs.Card{
        Due:           review.NextReviewAt,
        Stability:     review.Stability,
        Difficulty:    review.Difficulty,
        ElapsedDays:   review.ElapsedDays,
        ScheduledDays: review.ScheduledDays,
        Reps:          review.Reps,
        Lapses:        review.Lapses,
        State:         fsrs.State(review.State),
        LastReview:    review.LastReview,
    }
}

// ConvertFSRSToReviewSchedule updates a ReviewSchedule with values from an FSRS Card
func ConvertFSRSToReviewSchedule(card fsrs.Card, review *ReviewSchedule) {
    review.NextReviewAt = card.Due
    review.Stability = card.Stability
    review.Difficulty = card.Difficulty
    review.ElapsedDays = card.ElapsedDays
    review.ScheduledDays = card.ScheduledDays
    review.Reps = card.Reps
    review.Lapses = card.Lapses
    review.State = int(card.State)
    review.LastReview = card.LastReview
}

func (s *ReviewScheduleStore) UpdateOrCreateReviewForSubmission(submission *Submission) (ReviewSchedule, error) {
    // Check if we already have a review for this problem
    existingReview, err := s.GetReviewByTitleSlug(submission.UserID, submission.TitleSlug)
    
    if err == nil {
        fsrsCard := ConvertReviewScheduleToFSRS(&existingReview)
        
        // Process with "Good" rating as a baseline for solving the problem again
        fsrsScheduler := fsrs.NewFSRS(fsrs.DefaultParam())
        now := time.Now().UTC()
        result := fsrsScheduler.Next(fsrsCard, now, fsrs.Good)
        
        // Update the review with new FSRS values
        existingReview.SubmissionID = submission.ID
        ConvertFSRSToReviewSchedule(result.Card, &existingReview)
        existingReview.LastReview = now
        
        if err := s.UpdateReviewSchedule(&existingReview); err != nil {
            return ReviewSchedule{}, fmt.Errorf("error updating existing review: %v", err)
        }
        return existingReview, nil
    }
    
    // No review exists, create a new one
    // Initialize FSRS parameters
    fsrsScheduler := fsrs.NewFSRS(fsrs.DefaultParam())
    card := fsrs.NewCard()
    
    // Create initial schedule with "Good" rating
    now := time.Now().UTC()
    result := fsrsScheduler.Next(card, now, fsrs.Good)
    
    newReview := ReviewSchedule{
        SubmissionID: submission.ID,
        CreatedAt:    now,
    }
    
    // Set FSRS fields
    ConvertFSRSToReviewSchedule(result.Card, &newReview)
    
    if err := s.CreateReviewSchedule(&newReview); err != nil {
        return ReviewSchedule{}, fmt.Errorf("error creating new review: %v", err)
    }
    return newReview, nil
}