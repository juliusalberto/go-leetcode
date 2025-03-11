package handlers

import (
	"encoding/json"
	"fmt"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
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
		response.ValidationError(w, "user_id", "Missing user_id parameter")
		return
	}
	
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.ValidationError(w, "user_id", "Invalid user_id parameter")
		return
	}

	submissions, err := s.store.GetSubmissionsByUserID(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get submissions")
		return
	}

	response.JSON(w, http.StatusOK, submissions)
}

func (s *SubmissionHandler) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	var subReq struct{
		UserID int`json:"user_id"`
		Title string `json:"title"`
		TitleSlug string `json:"title-slug"`
		SubmittedAt string `json:"submitted_at"`
	}

	if err := json.NewDecoder(r.Body).Decode(&subReq); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
		return
	}

	id := uuid.New().String()
    shortID := strings.Replace(id, "-", "", -1)[:12]
	submitted_time, err := time.Parse(time.RFC3339, subReq.SubmittedAt)

	if err != nil {
		response.ValidationError(w, "submitted_at", "Invalid time format")
		return
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
		response_str := "Failed to create new submission"
		response.Error(w, http.StatusInternalServerError, "server_error", response_str)
		return
	}

	response.JSON(w, http.StatusCreated, submissionToAdd)
}

func NewSubmissionHandler(store *models.SubmissionStore) *SubmissionHandler {
	return &SubmissionHandler{store: store}
}
