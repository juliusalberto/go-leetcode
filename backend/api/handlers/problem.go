package handlers

import (
	"encoding/json"
	"go-leetcode/backend/models"
	"net/http"
	"strconv"
	"strings"
)

type ProblemHandler struct {
	store *models.ProblemStore
}

func (h *ProblemHandler) GetProblemByID(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	res, err := h.store.GetProblemByID(req.ID)

	if err != nil {
		http.Error(w, "Failed to get problem", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *ProblemHandler) GetProblemByFrontendID(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FrontendID int `json:"frontend_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	res, err := h.store.GetProblemByFrontendID(req.FrontendID)

	if err != nil {
		http.Error(w, "Failed to get problem", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *ProblemHandler) GetProblemBySlug(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Slug string `json:"slug"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	res, err := h.store.GetProblemBySlug(req.Slug)

	if err != nil {
		http.Error(w, "Failed to get problem", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
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
		tags = strings.Split(tagsParam, "")
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

	res, err := h.store.ListProblems(options)

	if err != nil {
		http.Error(w, "Failed to get problems list", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}