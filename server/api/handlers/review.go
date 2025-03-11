package handlers

import (
	"encoding/json"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
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
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid Request Body")
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
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create new review")
		return
	}

	response.JSON(w, http.StatusCreated, map[string]int{"id": reviewToAdd.ID})
}

func NewReviewHandler(store *models.ReviewScheduleStore) *ReviewHandler {
	return &ReviewHandler{store: store}
}

func (h *ReviewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	// Parse user_id from query params
	reqUserID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		response.ValidationError(w, "user_id", "Invalid user_id format")
		return
	}

	if reqUserID < 0 {
		response.ValidationError(w, "user_id", "UserID should not be below 0")
		return
	}

	// Parse pagination parameters
	page := 1
	perPage := 10
	
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if pageNum, err := strconv.Atoi(pageStr); err == nil && pageNum > 0 {
			page = pageNum
		}
	}
	
	if perPageStr := r.URL.Query().Get("per_page"); perPageStr != "" {
		if perPageNum, err := strconv.Atoi(perPageStr); err == nil && perPageNum > 0 && perPageNum <= 100 {
			perPage = perPageNum
		}
	}
	
	// Calculate offset
	offset := (page - 1) * perPage

	// Get reviews based on status parameter
	status := r.URL.Query().Get("status")
	var reviews []models.ReviewSchedule
	var total int

	switch status {
	case "due":
		// Get only due reviews with pagination
		reviews, total, err = h.store.GetDueReviews(reqUserID, perPage, offset)
	case "upcoming":
		// Get only upcoming reviews with pagination
		reviews, total, err = h.store.GetUpcomingReviews(reqUserID, perPage, offset)
	default:
		// Get all reviews (both due and upcoming)
		// For combined results, we need to handle pagination specially
		dueReviews, dueTotal, dueErr := h.store.GetDueReviews(reqUserID, perPage, offset)
		
		if dueErr != nil {
			err = dueErr
		} else {
			reviews = dueReviews
			total = dueTotal
			
			// If we haven't filled the page with due reviews, get some upcoming reviews
			if len(dueReviews) < perPage {
				remainingItems := perPage - len(dueReviews)
				upcomingReviews, upcomingTotal, upcomingErr := h.store.GetUpcomingReviews(reqUserID, remainingItems, 0)
				
				if upcomingErr != nil {
					err = upcomingErr
				} else {
					// Combine both lists and totals
					reviews = append(reviews, upcomingReviews...)
					total += upcomingTotal
				}
			}
		}
	}

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get reviews")
		return
	}

	response.JSONWithPagination(w, http.StatusOK, reviews, total, page, perPage)
}

func (h *ReviewHandler) UpdateReviewSchedule(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     int `json:"review_id"`
		Rating int `json:"rating"` // 1=Again, 2=Hard, 3=Good, 4=Easy
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
		return
	}

	if req.ID < 0 {
		response.ValidationError(w, "review_id", "Review ID should not be below 0")
		return
	}

	if req.Rating < 1 || req.Rating > 4 {
		response.ValidationError(w, "rating", "Rating must be between 1 and 4")
		return
	}

	// Get the previous review
	currReview, err := h.store.GetReviewByID(req.ID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "not_found", "Failed to find review")
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
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to update review schedule")
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"success":           true,
		"next_review_at":    updatedReview.NextReviewAt,
		"days_until_review": int(updatedReview.ScheduledDays),
	})
}
