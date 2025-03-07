package handlers

import (
	"encoding/json"
	"fmt"
	"go-leetcode/backend/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SubmissionHandler struct {
	store *models.SubmissionStore
}

func (s *SubmissionHandler) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from query parameters
	userIDStr := r.URL.Query().Get("user_id")
	
	if userIDStr == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}
	
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id parameter", http.StatusBadRequest)
		return
	}

	submissions, err := s.store.GetSubmissionsByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to get submissions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(submissions)
}

func (s *SubmissionHandler) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	var subReq struct{
		UserID int`json:"user_id"`
		Title string `json:"title"`
		TitleSlug string `json:"title-slug"`
		SubmittedAt string `json:"submitted_at"`
	}

	if err := json.NewDecoder(r.Body).Decode(&subReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	id := uuid.New().String()
    shortID := strings.Replace(id, "-", "", -1)[:12]
	submitted_time, err := time.Parse(time.RFC3339, subReq.SubmittedAt)

	if err != nil {
		http.Error(w, "Invalid time format", http.StatusBadRequest)
	}

	submissionToAdd := models.Submission{
		ID: fmt.Sprintf("internal-user-%d", &shortID),
		UserID: subReq.UserID,
		Title: subReq.Title,
		TitleSlug: subReq.TitleSlug,
		CreatedAt: time.Now().UTC(),
		SubmittedAt: submitted_time,
	}

	err = s.store.CreateSubmission(submissionToAdd)
	if err != nil {
		http.Error(w, "Failed to create new submission", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(submissionToAdd)
	return
}

func NewSubmissionHandler(store *models.SubmissionStore) *SubmissionHandler {
	return &SubmissionHandler{store: store}
}
