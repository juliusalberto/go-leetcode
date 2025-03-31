package handlers

import (
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
	"net/http"
)

type AuthStatusHandler struct {
	store *models.UserStore
}

type AuthStatusResponse struct {
	UserID			string	`json:"user_id"`
	Email			string 	`json:"email"`
	ProfileExists	bool 	`json:"profile_exists"`
	ProfileComplete	bool	`json:"profile_complete"`
}

func NewAuthStatusHandler(s *models.UserStore)(*AuthStatusHandler) {
	return &AuthStatusHandler{store: s}
}

func (h *AuthStatusHandler) GetUserAuthStatus(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserUUIDFromContext(r.Context())
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid_context", "Could not identify authenticated user")
	}

	email, _ := middleware.GetStringFromContext(r.Context(), middleware.UserEmailKey)
	// check if the userID is in the DB
	exists, err := h.store.CheckUserExistsByID(userID);
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "db_error", "failed to check user status")
		return
	}

	auth_response := AuthStatusResponse{
		UserID: userID.String(),
		Email: email,
		ProfileExists: exists,
		ProfileComplete: exists, // initially complete means it exists - can be enhanced later
	}

	response.JSON(w, http.StatusOK, auth_response)
}