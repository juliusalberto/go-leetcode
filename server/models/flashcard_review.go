package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/open-spaced-repetition/go-fsrs/v3"
)

func NewFlashcardReviewStore(db *sql.DB) *FlashcardReviewStore {
	return &FlashcardReviewStore{db: db}
}

type FlashcardReview struct {
	ID        int       `json:"id"`
	ProblemID int       `json:"problem_id"`
	UserID    string    `json:"user_id"`
	DeckID    int       `json:"deck_id"`
	FsrsCard  fsrs.Card `json:"fsrs_card"`
}

type FlashcardReviewLog struct {
	ID                int       `json:"id"`
	FlashcardReviewID int       `json:"flashcard_review_id"`
	Rating            int       `json:"rating"`
	ReviewDate        time.Time `json:"review_date"`
	ElapsedDays       int       `json:"elapsed_days"`
	ScheduledDays     int       `json:"scheduled_days"`
	State             int       `json:"state"`
}

type FlashcardReviewWithProblem struct {
	FlashcardReview
	Problem Problem `json:"problem"`
}

type FlashcardReviewStore struct {
	db *sql.DB
}

func (s *FlashcardReviewStore) GetDueFlashcardReviews(userID uuid.UUID, deckID int, limit, offset int) ([]FlashcardReviewWithProblem, int, error) {
	baseQuery := `
		SELECT
			fr.id, fr.problem_id, fr.user_id, fr.deck_id,
			fr.stability, fr.difficulty, fr.elapsed_days, fr.scheduled_days,
			fr.reps, fr.lapses, fr.state, fr.last_review, fr.next_review_at,
			p.id, p.frontend_id, p.title, p.title_slug, p.difficulty, p.is_paid_only, p.content, p.solution_approach
		FROM flashcard_reviews fr
		JOIN problems p ON fr.problem_id = p.id
		WHERE fr.user_id = $1 AND fr.next_review_at <= NOW()::timestamp
	`
	countQuery := `SELECT COUNT(*) FROM flashcard_reviews WHERE user_id = $1 AND next_review_at <= NOW()::timestamp`

	var params []interface{}
	params = append(params, userID.String())
	paramPos := 2

	if deckID > 0 {
		baseQuery += fmt.Sprintf(" AND fr.deck_id = $%d", paramPos)
		countQuery += fmt.Sprintf(" AND deck_id = $%d", paramPos)
		params = append(params, deckID)
		paramPos++
	}

	baseQuery += fmt.Sprintf(" ORDER BY fr.next_review_at LIMIT $%d OFFSET $%d", paramPos, paramPos+1)
	params = append(params, limit, offset)

	var total int
	err := s.db.QueryRow(countQuery, params[:paramPos-1]...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := s.db.Query(baseQuery, params...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var reviews []FlashcardReviewWithProblem
	for rows.Next() {
		var review FlashcardReviewWithProblem
		err := rows.Scan(
			&review.ID,
			&review.ProblemID,
			&review.UserID,
			&review.DeckID,
			&review.FsrsCard.Stability,
			&review.FsrsCard.Difficulty,
			&review.FsrsCard.ElapsedDays,
			&review.FsrsCard.ScheduledDays,
			&review.FsrsCard.Reps,
			&review.FsrsCard.Lapses,
			&review.FsrsCard.State,
			&review.FsrsCard.LastReview,
			&review.FsrsCard.Due,
			&review.Problem.ID,
			&review.Problem.FrontendID,
			&review.Problem.Title,
			&review.Problem.TitleSlug,
			&review.Problem.Difficulty,
			&review.Problem.IsPaidOnly,
			&review.Problem.Content,
			&review.Problem.SolutionApproach,
		)
		if err != nil {
			return nil, 0, err
		}

		reviews = append(reviews, review)
	}

	return reviews, total, nil
}

func (s *FlashcardReviewStore) CreateFlashcardReview(review *FlashcardReview) error {
	query := `
		INSERT INTO flashcard_reviews (
			problem_id, user_id, deck_id,
			stability, difficulty, elapsed_days, scheduled_days,
			reps, lapses, state, last_review, next_review_at
		) VALUES (
			$1, $2, $3,
			$4, $5, $6, $7,
			$8, $9, $10, $11, $12
		)
		RETURNING id
	`
	return s.db.QueryRow(query,
		review.ProblemID,
		review.UserID,
		review.DeckID,
		review.FsrsCard.Stability,
		review.FsrsCard.Difficulty,
		review.FsrsCard.ElapsedDays,
		review.FsrsCard.ScheduledDays,
		review.FsrsCard.Reps,
		review.FsrsCard.Lapses,
		review.FsrsCard.State,
		review.FsrsCard.LastReview,
		review.FsrsCard.Due,
	).Scan(&review.ID)
}

func (s *FlashcardReviewStore) UpdateFlashcardReview(review *FlashcardReview) error {
	query := `
		UPDATE flashcard_reviews
		SET
			stability = $1,
			difficulty = $2,
			elapsed_days = $3,
			scheduled_days = $4,
			reps = $5,
			lapses = $6,
			state = $7,
			last_review = $8,
			next_review_at = $9
		WHERE id = $10
	`
	_, err := s.db.Exec(query,
		review.FsrsCard.Stability,
		review.FsrsCard.Difficulty,
		review.FsrsCard.ElapsedDays,
		review.FsrsCard.ScheduledDays,
		review.FsrsCard.Reps,
		review.FsrsCard.Lapses,
		review.FsrsCard.State,
		review.FsrsCard.LastReview,
		review.FsrsCard.Due,
		review.ID,
	)
	return err
}

func (s *FlashcardReviewStore) CreateFlashcardReviewLog(log *FlashcardReviewLog) error {
	query := `
		INSERT INTO flashcard_review_logs 
		(flashcard_review_id, rating, review_date, elapsed_days, scheduled_days, state)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	return s.db.QueryRow(query,
		log.FlashcardReviewID,
		log.Rating,
		log.ReviewDate,
		log.ElapsedDays,
		log.ScheduledDays,
		log.State,
	).Scan(&log.ID)
}

func (s *FlashcardReviewStore) GetReviewByID(reviewID int) (FlashcardReview, error) {
	query := `
		SELECT
			id, problem_id, user_id, deck_id,
			stability, difficulty, elapsed_days, scheduled_days,
			reps, lapses, state, last_review, next_review_at
		FROM flashcard_reviews
		WHERE id = $1
	`
	var review FlashcardReview
	err := s.db.QueryRow(query, reviewID).Scan(
		&review.ID,
		&review.ProblemID,
		&review.UserID,
		&review.DeckID,
		&review.FsrsCard.Stability,
		&review.FsrsCard.Difficulty,
		&review.FsrsCard.ElapsedDays,
		&review.FsrsCard.ScheduledDays,
		&review.FsrsCard.Reps,
		&review.FsrsCard.Lapses,
		&review.FsrsCard.State,
		&review.FsrsCard.LastReview,
		&review.FsrsCard.Due,
	)
	if err != nil {
		return FlashcardReview{}, err
	}
	return review, nil
}

func (s *FlashcardReviewStore) AddDeckToUserFlashcards(userID uuid.UUID, deckID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO flashcard_reviews (
			problem_id, user_id, deck_id,
			stability, difficulty, elapsed_days, scheduled_days,
			reps, lapses, state, last_review, next_review_at
		)
		SELECT dp.problem_id, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 -- Params: userID, deckID, fsrs_params...
		FROM deck_problems dp
		-- Check if review exists for this user AND this specific deck_id
		LEFT JOIN flashcard_reviews fr ON dp.problem_id = fr.problem_id AND fr.user_id = $1 AND fr.deck_id = $2
		WHERE dp.deck_id = $2 AND fr.id IS NULL -- Only insert if no review exists for this specific user/problem/deck combo
	`

	// Initialize new FSRS card with default state
	defaultCard := fsrs.NewCard()
	now := time.Now().UTC()

	_, err = tx.Exec(query,
		userID.String(),
		deckID,
		defaultCard.Stability,
		defaultCard.Difficulty,
		defaultCard.ElapsedDays,
		defaultCard.ScheduledDays,
		defaultCard.Reps,
		defaultCard.Lapses,
		defaultCard.State,
		now,    // last_review
		now,    // due (initial review is due immediately)
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
