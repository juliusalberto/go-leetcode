package handlers

import (
	"net/http"
	"go-leetcode/backend/pkg/response"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}