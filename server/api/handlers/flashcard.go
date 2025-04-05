package handlers

import (
	"encoding/json"
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"fmt" // Add fmt import
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/open-spaced-repetition/go-fsrs/v3"
)

type FlashcardHandler struct {
	store        *models.FlashcardReviewStore
	problemStore *models.ProblemStore
	deckStore    *models.DeckStore
}

func NewFlashcardHandler(
	store *models.FlashcardReviewStore,
	problemStore *models.ProblemStore,
	deckStore *models.DeckStore,
) *FlashcardHandler {
	return &FlashcardHandler{
		store:        store,
		problemStore: problemStore,
		deckStore:    deckStore,
	}
}

func (h *FlashcardHandler) GetFlashcardReviews(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	deckIDStr := r.URL.Query().Get("deck_id")
	var deckID int
	if deckIDStr != "" {
		var err error
		deckID, err = strconv.Atoi(deckIDStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
			return
		}
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "bad_request", "Invalid limit")
			return
		}
	}

	offsetStr := r.URL.Query().Get("offset")
	offset := 0
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "bad_request", "Invalid offset")
			return
		}
	}

	reviews, total, err := h.store.GetDueFlashcardReviews(userID, deckID, limit, offset)
	if err != nil {
		// Include detailed error in the response message
		errorMessage := fmt.Sprintf("Failed to get flashcard reviews: %v", err)
		response.Error(w, http.StatusInternalServerError, "server_error", errorMessage)
		return
	}

	response.JSON(w, http.StatusOK, struct {
		Reviews []models.FlashcardReviewWithProblem `json:"reviews"`
		Total   int                                 `json:"total"`
	}{
		Reviews: reviews,
		Total:   total,
	})
}

func (h *FlashcardHandler) SubmitFlashcardReview(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	var req struct {
		ReviewID int `json:"review_id"`
		Rating   int `json:"rating"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid request body")
		return
	}

	// Get the current review
	review, err := h.store.GetReviewByID(req.ReviewID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get review")
		return
	}

	// Verify review belongs to user
	if review.UserID != userID.String() {
		response.Error(w, http.StatusForbidden, "forbidden", "Forbidden")
		return
	}

	if req.Rating < 1 || req.Rating > 4 {
		response.Error(w, http.StatusBadRequest, "bad_request", "Rating must be between 1 and 4")
		return
	}

	// Process with FSRS
	fsrsScheduler := fsrs.NewFSRS(fsrs.DefaultParam())
	now := time.Now().UTC()
	result := fsrsScheduler.Next(review.FsrsCard, now, fsrs.Rating(req.Rating))

	// Update review
	review.FsrsCard = result.Card
	review.FsrsCard.LastReview = now
	if err := h.store.UpdateFlashcardReview(&review); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to update review")
		return
	}

	// Create review log
	log := models.FlashcardReviewLog{
		FlashcardReviewID: review.ID,
		Rating:            req.Rating,
		ReviewDate:        now,
		ElapsedDays:       int(result.Card.ElapsedDays),
		ScheduledDays:     int(result.Card.ScheduledDays),
		State:             int(result.Card.State),
	}
	if err := h.store.CreateFlashcardReviewLog(&log); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create review log")
		return
	}

	response.JSON(w, http.StatusOK, struct {
		Success         bool      `json:"success"`
		NextReviewAt    time.Time `json:"next_review_at"`
		DaysUntilReview int       `json:"days_until_review"`
	}{
		Success:         true,
		NextReviewAt:    result.Card.Due,
		DaysUntilReview: int(result.Card.ScheduledDays),
	})
}

func (h *FlashcardHandler) AddDeckToFlashcards(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	deckIDStr := chi.URLParam(r, "deck_id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
		return
	}

	// Verify deck exists and is accessible to user
	deck, err := h.deckStore.GetDeckByID(deckID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get deck")
		return
	}

	if !deck.IsPublic && deck.UserID != userID.String() {
		response.Error(w, http.StatusForbidden, "forbidden", "Forbidden")
		return
	}

	if err := h.store.AddDeckToUserFlashcards(userID, deckID); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to add deck to flashcards")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
