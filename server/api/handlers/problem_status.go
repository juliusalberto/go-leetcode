package handlers

import (
	"fmt" // <-- Add fmt package import
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"strconv"
	"strings"
)

type ProblemStatusHandler struct {
	problemStore    *models.ProblemStore
	submissionStore *models.SubmissionStore
}

func NewProblemStatusHandler(ps *models.ProblemStore, ss *models.SubmissionStore) *ProblemStatusHandler {
	return &ProblemStatusHandler{
		problemStore:    ps,
		submissionStore: ss,
	}
}

func (h *ProblemStatusHandler) GetProblemsWithStatus(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if offset < 0 {
		offset = 0
	}

	difficulty := r.URL.Query().Get("difficulty")
	orderBy := r.URL.Query().Get("order_by")
	orderDir := r.URL.Query().Get("order_dir")
	searchKeyword := r.URL.Query().Get("search")

	var tags []string
	if tagsParam := r.URL.Query().Get("tags"); tagsParam != "" {
		tags = strings.Split(tagsParam, ",")
	}

	var paidOnly *bool
	if paidParam := r.URL.Query().Get("paid_only"); paidParam != "" {
		boolVal := paidParam == "true"
		paidOnly = &boolVal
	}

	options := models.ListProblemOptions{
		Filter: models.ProblemFilter{
			Difficulty:    difficulty,
			Tags:          tags,
			SearchKeyword: searchKeyword,
			PaidOnly:      paidOnly,
		},
		Limit:    limit,
		Offset:   offset,
		OrderBy:  orderBy,
		OrderDir: orderDir,
	}

	// Calculate page number from offset and limit
	page := (offset / limit) + 1

	userID, err := middleware.GetUserUUIDFromContext(r.Context())

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to extract user UUID")
		return
	}
	fmt.Printf("DEBUG GetProblemsWithStatus: Processing request for userID: %s\n", userID.String()) // Log UserID

	// Get filtered problems
	problemsList, err := h.problemStore.ListProblems(options)
	if err != nil {
		// Log the specific error on the server for debugging
		fmt.Printf("ERROR: Failed to list problems: %v\n", err)
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get problems list")
		return
	}

	// Get user's submissions
	submissions, err := h.submissionStore.GetSubmissionsByUserID(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get user submissions")
		return
	}
	fmt.Printf("DEBUG GetProblemsWithStatus: Found %d submissions for userID: %s\n", len(submissions), userID.String()) // Log submission count

	// Create a lookup map of completed problems
	completedMap := make(map[string]bool)
	for _, sub := range submissions {
		completedMap[sub.TitleSlug] = true
	}
	fmt.Printf("DEBUG GetProblemsWithStatus: Created completedMap with %d entries for userID: %s\n", len(completedMap), userID.String()) // Log map size

	// Combine data
	var problemsWithStatus []models.ProblemWithStatus
	for _, problem := range problemsList.Problems {
		problemsWithStatus = append(problemsWithStatus, models.ProblemWithStatus{
			Problem:   problem,
			Completed: completedMap[problem.TitleSlug],
		})
	}

	result := models.ProblemListWithStatus{
		Problems: problemsWithStatus,
		Total:    problemsList.Total,
	}

	response.JSONWithPagination(w, http.StatusOK, result.Problems, result.Total, page, limit)
}
