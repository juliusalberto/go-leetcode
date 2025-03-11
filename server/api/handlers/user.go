package handlers

import (
	"encoding/json"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
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
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
		return 
	}

	if req.Username == "" || req.LeetcodeUsername == "" {
		response.ValidationError(w, "username_or_leetcode_username", "Missing username or leetcode username")
		return
	}

	// check if username already exists
    exists, err := h.store.CheckUserExistsByUsername(req.Username)
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "server_error", "Internal server error")
        return
    }
    if exists {
        response.Error(w, http.StatusConflict, "conflict", "Username already exists")
        return
    }

	newUser := models.User{
		Username: req.Username,
		LeetcodeUsername: req.LeetcodeUsername,
		CreatedAt: time.Now(),
	}

	if err := h.store.CreateUser(&newUser); err != nil {
		response.Error(w, http.StatusInternalServerError, "server_error", "Failed to create user")
		return 
	}

	response.JSON(w, http.StatusCreated, newUser)
}