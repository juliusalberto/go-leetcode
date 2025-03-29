package handlers

import (
	"encoding/json"
	"fmt"
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/open-spaced-repetition/go-fsrs/v3"
)

type ReviewHandler struct {
	store           *models.ReviewScheduleStore
	submissionStore *models.SubmissionStore
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
		SubmissionID: req.SubmissionID,
		CreatedAt:    now,
	}
	
	// Set FSRS fields
	models.ConvertFSRSToReviewSchedule(result.Card, &reviewToAdd)
	reviewToAdd.LastReview = result.Card.LastReview

	err := h.store.CreateReviewSchedule(&reviewToAdd)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create new review")
		return
	}

	response.JSON(w, http.StatusCreated, map[string]int{"id": reviewToAdd.ID})
}

func NewReviewHandler(store *models.ReviewScheduleStore, submissionStore *models.SubmissionStore) *ReviewHandler {
	return &ReviewHandler{
		store:           store,
		submissionStore: submissionStore,
	}
}

func (h *ReviewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	// Parse user_id from query params
	userID, err := middleware.GetUserUUIDFromContext(r.Context())
	if err != nil {
		fmt.Printf("Internal Server Error: Failed to get user UUID from context in GetReviews: %v\n", err) // Log detailed error
		response.Error(w, http.StatusInternalServerError, "server_error", "Could not identify authenticated user")
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
		reviews, total, err = h.store.GetDueReviews(userID, perPage, offset)
	case "upcoming":
		// Get only upcoming reviews with pagination
		reviews, total, err = h.store.GetUpcomingReviews(userID, perPage, offset)
	default:
		// Get all reviews (both due and upcoming)
		// For combined results, we need to handle pagination specially
		dueReviews, dueTotal, dueErr := h.store.GetDueReviews(userID, perPage, offset)
		
		if dueErr != nil {
			err = dueErr
		} else {
			reviews = dueReviews
			total = dueTotal
			
			// If we haven't filled the page with due reviews, get some upcoming reviews
			if len(dueReviews) < perPage {
				remainingItems := perPage - len(dueReviews)
				upcomingReviews, upcomingTotal, upcomingErr := h.store.GetUpcomingReviews(userID, remainingItems, 0)
				
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
	fsrsCard := models.ConvertReviewScheduleToFSRS(&currReview)

	// Process the rating
	fsrsScheduler := fsrs.NewFSRS(fsrs.DefaultParam())
	now := time.Now()
	result := fsrsScheduler.Next(fsrsCard, now, fsrs.Rating(req.Rating))

	// Update the review
	updatedReview := currReview
	models.ConvertFSRSToReviewSchedule(result.Card, &updatedReview)
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

func (h *ReviewHandler) UpdateOrCreateReview(w http.ResponseWriter, r *http.Request) {
	// we expect that the serverless function will send us a submission data
	// this is mainly used to read data from the leetcode graphql api
	// thus, 

	var req models.Submission
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
		return
	}

	review, err := h.store.UpdateOrCreateReviewForSubmission(&req)
	
	if err != nil{
		err_string := fmt.Sprintf("Failed to update or create review: %v", err)
		response.Error(w, http.StatusInternalServerError, "server_error", err_string)
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"success":           true,
		"next_review_at":    review.NextReviewAt,
		"days_until_review": int(review.ScheduledDays),
	})
}

// ProcessSubmission handles creating/updating both submissions and reviews in one endpoint
func (h *ReviewHandler) ProcessSubmission(w http.ResponseWriter, r *http.Request) {
    // 1. Parse the submission data from request using the same structure as SubmissionHandler.CreateSubmission
    var subReq struct{
        IsInternal          bool   `json:"is_internal"`
        LeetcodeSubmissionID string `json:"leetcode_submission_id,omitempty"`
        Title               string `json:"title"`
        TitleSlug           string `json:"title_slug"`
        SubmittedAt         string `json:"submitted_at"`
    }

    if err := json.NewDecoder(r.Body).Decode(&subReq); err != nil {
        response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
        return
    }
    
    userID, err := middleware.GetUserUUIDFromContext(r.Context())
	if err != nil {
		fmt.Printf("Internal Server Error: Failed to get user UUID from context %v\n", err)
		response.Error(w, http.StatusInternalServerError, "server_error", "Could not identify authenticated user")
		return
	}

    
    if subReq.TitleSlug == "" {
        response.ValidationError(w, "title_slug", "Problem title slug is required")
        return
    }

    // Generate submission ID based on source
    var submissionID string
    
    if subReq.IsInternal {
        // Generate an internal ID
        id := uuid.New().String()
        shortID := strings.Replace(id, "-", "", -1)[:12]
        submissionID = fmt.Sprintf("internal-user-%s", shortID)
    } else {
        // Use the LeetCode ID with prefix
        if subReq.LeetcodeSubmissionID == "" {
            response.ValidationError(w, "leetcode_submission_id", "LeetcodeSubmissionID is required when IsInternal is false")
            return
        }
        submissionID = fmt.Sprintf("leetcode-%s", subReq.LeetcodeSubmissionID)
    }


    // Parse the submission time
    submittedTime, err := time.Parse(time.RFC3339, subReq.SubmittedAt)
    if err != nil {
        response.ValidationError(w, "submitted_at", "Invalid time format (expected RFC3339)")
        return
    }

    // Create the submission object
    sub := models.Submission{
        ID:          submissionID,
        UserID:      userID,
        Title:       subReq.Title,
        TitleSlug:   subReq.TitleSlug,
        CreatedAt:   time.Now().UTC(),
        SubmittedAt: submittedTime,
    }

    // 2. Check if submission exists - create if not
    exists, err := h.submissionStore.CheckSubmissionExists(sub.ID)
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "server_error", 
            fmt.Sprintf("Error checking submission: %v", err))
        return
    }
    
    if !exists {
        if err := h.submissionStore.CreateSubmission(sub); err != nil {
            response.Error(w, http.StatusInternalServerError, "server_error", 
                fmt.Sprintf("Failed to create submission: %v", err))
            return
        }
    } else {
		response.Error(w, http.StatusConflict, "insertion_error", 
                fmt.Sprintf("Submission with ID: %s already exists", sub.ID))
		return
	}
    
    // 3. Process the review using the existing method
    review, err := h.store.UpdateOrCreateReviewForSubmission(&sub)
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "server_error", 
            fmt.Sprintf("Failed to process review: %v", err))
        return
    }
    
    // 4. Return comprehensive response
    now := time.Now().UTC()
    isDue := now.After(review.NextReviewAt)
    
    response.JSON(w, http.StatusOK, map[string]interface{}{
        "success":           true,
        "submission_id":     sub.ID,
        "next_review_at":    review.NextReviewAt,
        "days_until_review": int(review.ScheduledDays),
        "is_due":            isDue,
        "title":             sub.Title,
        "title_slug":        sub.TitleSlug,
    })
}
