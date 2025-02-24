package handlers

import (
	"go-leetcode/backend/models"
	"net/http"
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
	
}