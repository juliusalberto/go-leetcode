package handlers

import (
	"encoding/json"
	"fmt"
	"go-leetcode/backend/api/middleware"
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

// PLANNED FOR REMOVAL
// SINCE WE ALR USE SUPABASE
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

// GetUser handles GET requests to fetch a user by username
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Get username from query parameters
	username := r.URL.Query().Get("username")
	
	if username == "" {
		response.ValidationError(w, "username", "Username parameter is required")
		return
	}
	
	// Get user by username
	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		response.Error(w, http.StatusNotFound, "not_found", "User not found")
		return
	}
	
	// Return user information
	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) CompleteProfile(w http.ResponseWriter, r *http.Request) {
	// get user details from context
	userID, err := middleware.GetUserUUIDFromContext(r.Context())
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid_context", "Could not read user ID from the context")
		return
	}

	email, _ := middleware.GetStringFromContext(r.Context(), middleware.UserEmailKey)
	// decode request body

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid_request", "Invalid request body")
	}

	// validate input
	// check if the username in the request is empty
	if user.Username == "" {
		response.ValidationError(w, "username", "username cannot be empty")
		return
	}

	if user.LeetcodeUsername == "" {
		response.ValidationError(w, "leetcode_username", "leetcode username cannot be empty")
		return
	}

	// check if uuid is already there

	if exists, _ := h.store.CheckUserExistsByID(userID); exists {
		fmt.Printf("ERROR: Attempted to complete profile for an existing user with the UID %s", userID)
		response.Error(w, http.StatusConflict, "profile_exists", "User profile already exists")
		return
	}

	// check if username already taken
	exists, err := h.store.CheckUserExistsByUsername(user.Username)
	if exists {
		response.Error(w, http.StatusConflict, "username_exists", "Username already exists")
		return
	}

	// Create user profile in the DB with the UUID from auth
	user.ID = userID
	user.Email = email
	
	// Use CreateUserByAuth which properly includes the ID field
	if err := h.store.CreateUserByAuth(&user); err != nil {
		fmt.Printf("ERROR: Failed to create user profile: %v\n", err)
		response.Error(w, http.StatusInternalServerError, "db_error", "Failed to create user profile")
		return
	}
	
	fmt.Printf("SUCCESS: Created user profile for %s with username %s\n", userID, user.Username)

	// return success response + user profile
	response.JSON(w, http.StatusOK, user)
}