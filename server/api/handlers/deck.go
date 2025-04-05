package handlers

import (
	"encoding/json"
	"fmt"
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/open-spaced-repetition/go-fsrs/v3"
)

type DeckHandler struct {
	store                *models.DeckStore
	problemStore         *models.ProblemStore
	flashcardReviewStore *models.FlashcardReviewStore
}

func NewDeckHandler(deckStore *models.DeckStore, problemStore *models.ProblemStore, flashcardReviewStore *models.FlashcardReviewStore) *DeckHandler {
	return &DeckHandler{
		store:                deckStore,
		problemStore:         problemStore,
		flashcardReviewStore: flashcardReviewStore,
	}
}

func (h *DeckHandler) GetAllDecks(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	publicDecks, err := h.store.GetAllPublicDecks()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get public decks")
		return
	}

	userDecks, err := h.store.GetUserDecks(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get user decks")
		return
	}

	response.JSON(w, http.StatusOK, struct {
		PublicDecks []models.Deck `json:"public_decks"`
		UserDecks   []models.Deck `json:"user_decks"`
	}{
		PublicDecks: publicDecks,
		UserDecks:   userDecks,
	})
}

func (h *DeckHandler) CreateDeck(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	var deck models.Deck
	if err := json.NewDecoder(r.Body).Decode(&deck); err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid request body")
		return
	}

	deck.UserID = userID.String()
	if err := h.store.CreateDeck(&deck); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create deck")
		return
	}

	response.JSON(w, http.StatusOK, deck)
}

func (h *DeckHandler) UpdateDeck(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	var deck models.Deck
	deckIDStr := chi.URLParam(r, "id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err := json.NewDecoder(r.Body).Decode(&deck); err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid request body")
		return
	}

	existingDeck, err := h.store.GetDeckByID(deckID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get deck")
		return
	}

	if existingDeck.UserID != userID.String() {
		response.Error(w, http.StatusForbidden, "forbidden", "Forbidden")
		return
	}

	if err := h.store.UpdateDeck(&deck); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to update deck")
		return
	}

	response.JSON(w, http.StatusOK, deck)
}

func (h *DeckHandler) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	deckIDStr := chi.URLParam(r, "id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
		return
	}

	deck, err := h.store.GetDeckByID(deckID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get deck")
		return
	}

	if deck.UserID != userID.String() {
		response.Error(w, http.StatusForbidden, "forbidden", "Forbidden")
		return
	}

	if err := h.store.DeleteDeck(deckID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to delete deck")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeckHandler) GetDeckProblems(w http.ResponseWriter, r *http.Request) {
	deckIDStr := chi.URLParam(r, "id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
		return
	}

	// Get limit and offset from query parameters, with defaults
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20 // Default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // Default offset
	}

	problems, err := h.store.GetProblemsInDeck(deckID, limit, offset)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get deck problems (deckID: %d, limit: %d, offset: %d): %v", deckID, limit, offset, err)
		response.Error(w, http.StatusInternalServerError, "server_error", errorMessage)
		return
	}

	response.JSON(w, http.StatusOK, problems)
}

func (h *DeckHandler) AddProblemToDeckAndCreateFlashcard(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	deckIDStr := chi.URLParam(r, "id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
		return
	}

	deck, err := h.store.GetDeckByID(deckID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get deck")
		return
	}

	if deck.UserID != userID.String() {
		response.Error(w, http.StatusForbidden, "forbidden", "Forbidden")
		return
	}

	var req struct {
		ProblemID int `json:"problem_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid request body")
		return
	}

	// Add problem to the deck
	if err := h.store.AddProblemToDeck(deckID, req.ProblemID); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to add problem to deck")
		return
	}

	// Automatically create a flashcard review for the added problem
	now := time.Now().UTC()
	defaultCard := fsrs.NewCard()
	review := &models.FlashcardReview{
		ProblemID: req.ProblemID,
		UserID:    userID.String(),
		DeckID:    deckID,
		FsrsCard:  defaultCard,
	}
	// Set initial due date to now for immediate review
	review.FsrsCard.Due = now
	review.FsrsCard.LastReview = now // Set last review to now as well

	if err := h.flashcardReviewStore.CreateFlashcardReview(review); err != nil {
		// Log the error, but don't fail the entire request? Or should we?
		// For now, let's return an error if flashcard creation fails.
		fmt.Printf("Error creating flashcard review for user %s, deck %d, problem %d: %v\n", userID.String(), deckID, req.ProblemID, err)
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create flashcard review for the added problem")
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func (h *DeckHandler) RemoveProblemFromDeck(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	deckIDStr := chi.URLParam(r, "id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
		return
	}

	problemIDStr := chi.URLParam(r, "problem_id")
	problemID, err := strconv.Atoi(problemIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid problem ID")
		return
	}

	deck, err := h.store.GetDeckByID(deckID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get deck")
		return
	}

	if deck.UserID != userID.String() {
		response.Error(w, http.StatusForbidden, "forbidden", "Forbidden")
		return
	}

	// Pass context and userID to the updated store method
	if err := h.store.RemoveProblemFromDeck(r.Context(), deckID, problemID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to remove problem from deck")
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

// StartPracticePublicDeck handles the logic for when a user starts practicing a public deck.
// It ensures that flashcard review entries are created for the user for all problems
// in the deck, but only if they don't already exist.
func (h *DeckHandler) StartPracticePublicDeck(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserUUIDFromContext(r.Context())
	if ok != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized", "Unauthorized")
		return
	}

	deckIDStr := chi.URLParam(r, "id")
	deckID, err := strconv.Atoi(deckIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad_request", "Invalid deck ID")
		return
	}

	// Fetch the deck to ensure it exists and is public
	deck, err := h.store.GetDeckByID(deckID)
	if err != nil {
		// Use errors.Is for better error checking if using Go 1.13+
		// For simplicity here, checking directly. Adapt as needed.
		if err.Error() == "sql: no rows in result set" { // Or use specific error type if your store returns one
			response.Error(w, http.StatusNotFound, "not_found", "Deck not found")
		} else {
			response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get deck details")
		}
		return
	}

	// Ensure the deck is public
	if !deck.IsPublic {
		response.Error(w, http.StatusForbidden, "forbidden", "Cannot start practice on a private deck you don't own via this endpoint")
		return
	}

	// Add the deck's problems to the user's flashcards if they aren't already there.
	// The underlying store function handles the logic to avoid duplicates.
	if err := h.flashcardReviewStore.AddDeckToUserFlashcards(userID, deckID); err != nil {
		fmt.Printf("Error adding deck %d to user %s flashcards: %v\n", deckID, userID.String(), err)
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to prepare deck for practice")
		return
	}

	// Respond with success
	response.JSON(w, http.StatusOK, map[string]string{"message": "Deck prepared for practice successfully"})
}
