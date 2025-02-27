package handlers

import (
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
)

type SubmissionHandler struct {
	store *models.SubmissionStore
}


func (s *SubmissionHandler) CreateSubmission(rr *httptest.ResponseRecorder, req *http.Request) {
	panic("unimplemented")
}

func NewSubmissionHandler(store *models.SubmissionStore) *SubmissionHandler {
	return &SubmissionHandler{store: store}
}
