package handlers

import (
	"fmt"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"strconv"
	"strings"
)

type ProblemHandler struct {
	store *models.ProblemStore
}

func NewProblemHandler(s *models.ProblemStore)(*ProblemHandler) {
	return &ProblemHandler{store: s}
}

func (h *ProblemHandler) GetProblemByID(w http.ResponseWriter, r *http.Request) {
	reqID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	res, err := h.store.GetProblemByID(reqID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get problem")
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (h *ProblemHandler) GetProblemByFrontendID(w http.ResponseWriter, r *http.Request) {
	reqFrontendID, _ := strconv.Atoi(r.URL.Query().Get("frontend_id"))

	res, err := h.store.GetProblemByFrontendID(reqFrontendID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get problem")
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (h *ProblemHandler) GetProblemBySlug(w http.ResponseWriter, r *http.Request) {
	req_slug := r.URL.Query().Get("slug")

	res, err := h.store.GetProblemBySlug(req_slug)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get problem")
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (h *ProblemHandler) GetProblemList(w http.ResponseWriter, r *http.Request) {
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

	options := models.ListProblemOptions {
		Filter: models.ProblemFilter{
			Difficulty: difficulty,
			Tags: tags,
			SearchKeyword: searchKeyword,
			PaidOnly: paidOnly,
		},
		Limit: limit,
		Offset: offset,
		OrderBy: orderBy,
		OrderDir: orderDir,
	}

	// Calculate page number from offset and limit
	page := (offset / limit) + 1

	result, err := h.store.ListProblems(options)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get problems list")
		return
	}

	response.JSONWithPagination(w, http.StatusOK, result.Problems, result.Total, page, limit)
}