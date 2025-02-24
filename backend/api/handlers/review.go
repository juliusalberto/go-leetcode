package handlers

import (
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

func NewReviewHandler(store *models.ReviewScheduleStore)(*ReviewHandler) {
	return &ReviewHandler{store: store}
}

func (h *ReviewHandler) GetUpcomingReviews(w http.ResponseWriter, r *http.Request) {

}

func (h *ReviewHandler) UpdateReviewSchedule(w http.ResponseWriter, r *http.Request) {

}