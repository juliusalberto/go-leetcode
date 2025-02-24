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
	}

	if req.UserID < 0 {
		http.Error(w, "UserID should not be below 0", http.StatusBadRequest)
	}

	reviews, err := h.store.GetUpcomingReviews(req.UserID)
	if err != nil {
		http.Error(w, "Failed to get upcoming reviews", http.StatusInternalServerError)
	}
	
	// write it down as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}

func (h *ReviewHandler) UpdateReviewSchedule(w http.ResponseWriter, r *http.Request) {

}