package handlers

import (
	"encoding/json"
	"go-leetcode/backend/models"
	"net/http"
	"time"
)

type UserHandler struct {
	store *models.UserStore
}

type RegisterRequest struct{
	Username string `json:"username"`
	LeetcodeUsername string `json:"leetcode_username"`
}

func NewUserHandler(store *models.UserStore)(*UserHandler) {
	return &UserHandler{store: store}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// first we decode the json body
	
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	if req.Username == "" || req.LeetcodeUsername == "" {
		http.Error(w, "Missing username or leetcode username", http.StatusBadRequest)
		return
	}

	// check if username already exists
    exists, err := h.store.CheckUserExistsByUsername(req.Username)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    if exists {
        http.Error(w, "Username already exists", http.StatusConflict)
        return
    }

	newUser := models.User{
		Username: req.Username,
		LeetcodeUsername: req.LeetcodeUsername,
		CreatedAt: time.Now(),
	}

	if err := h.store.CreateUser(&newUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}