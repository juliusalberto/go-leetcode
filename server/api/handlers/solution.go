package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
)

type SolutionHandler struct {
	store *models.SolutionStore
}

func NewSolutionHandler(s *models.SolutionStore) *SolutionHandler {
	return &SolutionHandler{store: s}
}

func (h *SolutionHandler) GetSolutions(w http.ResponseWriter, r *http.Request) {
	// Extract the problem ID from query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response.Error(w, http.StatusBadRequest, "missing_id", "Problem ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_id", "Invalid problem ID")
		return
	}

	// Check if language filter is specified
	language := r.URL.Query().Get("language")
	
	// If language is specified, get solution for that specific language
	if language != "" {
		solution, err := h.store.GetSolutionByProblemAndLanguage(id, language)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get solution")
			return
		}

		// Return a single solution as a language-code map
		solutionMap := map[string]string{
			solution.Language: solution.SolutionCode,
		}

		jsonData, err := json.Marshal(solutionMap)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "server_error", "Failed to marshal solution")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	
	// Otherwise, get all solutions for this problem and structure them as a map
	solutions, err := h.store.GetSolutionsByProblemID(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to get solutions")
		return
	}

	// Convert the solutions array to a map of language -> code
	solutionsMap := make(map[string]string)
	for _, solution := range solutions {
		solutionsMap[solution.Language] = solution.SolutionCode
	}

	jsonData, err := json.Marshal(solutionsMap)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to marshal solutions")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *SolutionHandler) CreateSolution(w http.ResponseWriter, r *http.Request) {
	var solution models.Solution
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&solution); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
		return
	}

	// Validate that problem ID is provided in the request body
	if solution.ProblemID == 0 {
		response.Error(w, http.StatusBadRequest, "missing_problem_id", "Problem ID is required")
		return
	}

	createdSolution, err := h.store.CreateSolution(solution)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create solution")
		return
	}

	jsonData, err := json.Marshal(createdSolution)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to marshal created solution")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (h *SolutionHandler) UpdateSolution(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response.Error(w, http.StatusBadRequest, "missing_id", "Solution ID is required")
		return
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_id", "Invalid solution ID")
		return
	}

	var solution models.Solution
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&solution); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
		return
	}

	solution.ID = id

	err = h.store.UpdateSolution(solution)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to update solution")
		return
	}

	jsonData, err := json.Marshal(solution)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to marshal updated solution")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *SolutionHandler) DeleteSolution(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response.Error(w, http.StatusBadRequest, "missing_id", "Solution ID is required")
		return
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_id", "Invalid solution ID")
		return
	}

	err = h.store.DeleteSolution(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to delete solution")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
