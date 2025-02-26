package handlers

import (
	"encoding/json"
	"go-leetcode/backend/models"
	"net/http"
	"time"
)

type ReviewHandler struct {
	store *models.ReviewScheduleStore
}

type UpdateRequest struct {
	ID 				int
	NextReviewAt 	time.Time
	TimesReviewed	int
}

type GetReviewRequest struct {
	UserID int `json:"user_id"`
}

func NewReviewHandler(store *models.ReviewScheduleStore)(*ReviewHandler) {
	return &ReviewHandler{store: store}
}

func (h *ReviewHandler) GetUpcomingReviews(w http.ResponseWriter, r *http.Request) {
	// we want to write the review that is in the db

	var req GetReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if req.UserID < 0 {
		http.Error(w, "UserID should not be below 0", http.StatusBadRequest)
		return
	}

	reviews, err := h.store.GetUpcomingReviews(req.UserID)
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
		ID int `json:"review_id"`
		NextReviewAt string `json:"next_review_at"`
		TimesReviewed int `json:"times_reviewed"`
		IntervalDays int `json:"interval_days"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.ID < 0 {
		http.Error(w, "Review ID should not be below 0", http.StatusBadRequest)
		return
	}

	// get the previous review first
	currReview, err := h.store.GetReviewByID(req.ID)
	if err != nil {
		http.Error(w, "Failed to find review", http.StatusNotFound)
		return
	}

	nextReviewTime, err := time.Parse(time.RFC3339, req.NextReviewAt)
	if err != nil {
		http.Error(w, "Failed to parse time", http.StatusNotFound)
		return
	}

	updatedReview := currReview
	updatedReview.NextReviewAt = nextReviewTime
	updatedReview.TimesReviewed = req.TimesReviewed
	updatedReview.IntervalDays = req.IntervalDays

	if err := h.store.UpdateReviewSchedule(&updatedReview); err != nil {
		http.Error(w, "failed to update review schedule", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}