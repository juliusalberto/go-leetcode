package leetcode

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecentSubmission(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
            "data": {
                "recentAcSubmissionList": [
                    {
                        "id": "123",
                        "title": "Two Sum",
                        "titleSlug": "two-sum",
                        "timestamp": 1624512345
                    }
                ]
            }
		}`))
	}))

	defer server.Close()

	client := &Client{endpoint: server.URL}

	submissions, err := client.GetRecentSubmission("testuser", 1)

	if (err != nil) {
		t.Errorf("expected no error, got %v", err)
	}

	if len(submissions) != 1 {
		t.Errorf("expected 1 submission, got %d", len(submissions))
	}

	expected := Submission{
		ID: "123",
		Title: "Two Sum",
		TitleSlug: "two-sum",
		Timestamp: 1624512345,
	}

	if submissions[0] != expected {
		t.Errorf("expected %+v, got %+v", expected, submissions[0])
	}
}