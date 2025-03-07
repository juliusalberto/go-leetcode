package handlers

import (
	"encoding/json"
	"go-leetcode/backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/open-spaced-repetition/go-fsrs/v3"
)

type ReviewHandler struct {
	store *models.ReviewScheduleStore
}


func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SubmissionID string `json:"submission_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Initialize FSRS parameters
	f := fsrs.NewFSRS(fsrs.DefaultParam())
	card := fsrs.NewCard()
	
	// Create initial schedule with "Good" rating
	now := time.Now().UTC()
	result := f.Next(card, now, fsrs.Good)

	reviewToAdd := models.ReviewSchedule{
		SubmissionID:  req.SubmissionID,
		NextReviewAt:  result.Card.Due,
		CreatedAt:     now,
		Stability:     result.Card.Stability,
		Difficulty:    result.Card.Difficulty,
		ElapsedDays:   result.Card.ElapsedDays,
		ScheduledDays: result.Card.ScheduledDays,
		Reps:          result.Card.Reps,
		Lapses:        result.Card.Lapses,
		State:         int(result.Card.State),
		LastReview:    result.Card.LastReview,
	}

	err := h.store.CreateReviewSchedule(&reviewToAdd)
	if err != nil {
		http.Error(w, "Failed to create new review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": reviewToAdd.ID})
}

func NewReviewHandler(store *models.ReviewScheduleStore) *ReviewHandler {
	return &ReviewHandler{store: store}
}


func (h *ReviewHandler) GetUpcomingReviews(w http.ResponseWriter, r *http.Request) {
	// we want to write the review that is in the db
	reqUserID, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if reqUserID < 0 {
		http.Error(w, "UserID should not be below 0", http.StatusBadRequest)
		return
	}

	reviews, err := h.store.GetUpcomingReviews(reqUserID)
	if err != nil {
		http.Error(w, "Failed to get upcoming reviews", http.StatusInternalServerError)
		return
	}

	// write it down as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}

func (h *ReviewHandler) UpdateReviewSchedule(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     int `json:"review_id"`
		Rating int `json:"rating"` // 1=Again, 2=Hard, 3=Good, 4=Easy
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.ID < 0 {
		http.Error(w, "Review ID should not be below 0", http.StatusBadRequest)
		return
	}

	if req.Rating < 1 || req.Rating > 4 {
		http.Error(w, "Rating must be between 1 and 4", http.StatusBadRequest)
		return
	}

	// Get the previous review
	currReview, err := h.store.GetReviewByID(req.ID)
	if err != nil {
		http.Error(w, "Failed to find review", http.StatusNotFound)
		return
	}

	// Convert to FSRS Card
	fsrsCard := fsrs.Card{
		Due:           currReview.NextReviewAt,
		Stability:     currReview.Stability,
		Difficulty:    currReview.Difficulty,
		ElapsedDays:   currReview.ElapsedDays,
		ScheduledDays: currReview.ScheduledDays,
		Reps:          currReview.Reps,
		Lapses:        currReview.Lapses,
		State:         fsrs.State(currReview.State),
		LastReview:    currReview.LastReview,
	}

	// Process the rating
	fsrsScheduler := fsrs.NewFSRS(fsrs.DefaultParam())
	now := time.Now()
	result := fsrsScheduler.Next(fsrsCard, now, fsrs.Rating(req.Rating))

	// Update the review
	updatedReview := currReview
	updatedReview.NextReviewAt = result.Card.Due
	updatedReview.Stability = result.Card.Stability
	updatedReview.Difficulty = result.Card.Difficulty
	updatedReview.ElapsedDays = result.Card.ElapsedDays
	updatedReview.ScheduledDays = result.Card.ScheduledDays
	updatedReview.Reps = result.Card.Reps
	updatedReview.Lapses = result.Card.Lapses
	updatedReview.State = int(result.Card.State)
	updatedReview.LastReview = now

	if err := h.store.UpdateReviewSchedule(&updatedReview); err != nil {
		http.Error(w, "Failed to update review schedule", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"next_review_at": updatedReview.NextReviewAt,
		"days_until_review": int(updatedReview.ScheduledDays),
	})
}
