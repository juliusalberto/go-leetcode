package handlers

import "net/http"

func HealthCheck(w http.ResponseWriter, r* http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
}