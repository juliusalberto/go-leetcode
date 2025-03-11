package response

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response is the standard API response format
type Response struct {
	Data   interface{}   `json:"data"`
	Meta   *MetaData     `json:"meta"`
	Errors []ErrorDetail `json:"errors"`
}

// MetaData contains metadata information
type MetaData struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Timestamp  string      `json:"timestamp"`
}

// Pagination information
type Pagination struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// ErrorDetail contains error information
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

// JSON sends a standardized JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := Response{
		Data: data,
		Meta: &MetaData{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Errors: []ErrorDetail{},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// JSONWithPagination sends a JSON response with pagination data
func JSONWithPagination(w http.ResponseWriter, statusCode int, data interface{}, total, page, perPage int) {
	resp := Response{
		Data: data,
		Meta: &MetaData{
			Pagination: &Pagination{
				Total:   total,
				Page:    page,
				PerPage: perPage,
			},
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Errors: []ErrorDetail{},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// Error sends a standardized error response
func Error(w http.ResponseWriter, statusCode int, code, message string) {
	resp := Response{
		Data: nil,
		Meta: &MetaData{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Errors: []ErrorDetail{
			{
				Code:    code,
				Message: message,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// ValidationError sends a validation error response
func ValidationError(w http.ResponseWriter, field, message string) {
	resp := Response{
		Data: nil,
		Meta: &MetaData{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Errors: []ErrorDetail{
			{
				Code:    "validation_error",
				Message: message,
				Field:   field,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}