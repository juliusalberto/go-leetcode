package handlers

import (
	// "fmt" // Removed unused import
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
	"time" // Import time package
)

type AuthStatusHandler struct {
	store *models.UserStore
}

type AuthStatusResponse struct {
	UserID           string `json:"user_id"`
	Email            string `json:"email"`
	ProfileExists    bool   `json:"profile_exists"`
	ProfileComplete  bool   `json:"profile_complete"`
	LeetcodeUsername string `json:"leetcode_username"`
}

func NewAuthStatusHandler(s *models.UserStore) *AuthStatusHandler {
	return &AuthStatusHandler{store: s}
}

func (h *AuthStatusHandler) GetUserAuthStatus(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserUUIDFromContext(r.Context())
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid_context", "Could not identify authenticated user")
		return
	}

	email, _ := middleware.GetStringFromContext(r.Context(), middleware.UserEmailKey)
	// check if the userID is in the DB
	println("AuthStatus: Checking if user exists with ID:", userID.String())
	exists, err := h.store.CheckUserExistsByID(userID)
	// Handle initial check error - Check if it's a real error, not just "not found"
	// Note: CheckUserExistsByID currently doesn't return a specific "not found" error, only true/false or other DB errors.
	if err != nil {
		println("AuthStatus: Error checking user existence (initial):", err.Error())
		response.Error(w, http.StatusInternalServerError, "db_error", "failed to check user status")
		return
	}
	println("AuthStatus: Initial check - User exists:", exists)

	// If user doesn't exist on first check, retry after a short delay to handle potential visibility delay
	if !exists {
		time.Sleep(250 * time.Millisecond) // Adjust delay if needed (e.g., 100-500ms)
		println("AuthStatus: Retrying user existence check for:", userID.String())
		exists, err = h.store.CheckUserExistsByID(userID)
		// Handle retry check error
		if err != nil {
			println("AuthStatus: Error checking user existence (retry):", err.Error())
			// Decide how to handle retry error: proceed with exists=false or return server error?
			exists = false
		}
		println("AuthStatus: Retry check - User exists:", exists)
	}

	leetcodeUsername := ""
	if exists {
		// Attempt to get user details only if exists is true after potential retry
		user, getUserErr := h.store.GetUserByID(userID)
		if getUserErr == nil {
			leetcodeUsername = user.LeetcodeUsername
		} else {
			// Log if getting user details failed even if exists=true (shouldn't happen often)
			println("AuthStatus: Warning - User reported as existing, but GetUserByID failed:", getUserErr.Error())
		}
	}

	auth_response := AuthStatusResponse{
		UserID:           userID.String(),
		Email:            email,
		ProfileExists:    exists,
		ProfileComplete:  exists, // initially complete means it exists - can be enhanced later
		LeetcodeUsername: leetcodeUsername,
	}

	response.JSON(w, http.StatusOK, auth_response)
}
