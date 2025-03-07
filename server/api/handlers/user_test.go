package handlers

import (
	"bytes"
	"encoding/json"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupUserTest(t *testing.T)(*UserHandler, *database.TestDB) {
	testDB := database.SetupTestDB(t)
	userStore := models.NewUserStore(testDB.DB)
	handler := NewUserHandler(userStore)
	return handler, testDB
}

func TestRegisterHandler(t* testing.T) {
	handler, testDB := setupUserTest(t)
	defer testDB.Cleanup(t)

	tests := []struct {
		name 		string 
		body		RegisterRequest
		wantStatus	int
	}{
		{
			name: "valid registration",
			body: RegisterRequest{
				Username: "testuser",
				LeetcodeUsername: "leetcode_testuser",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "missing username",
			body: RegisterRequest{
				LeetcodeUsername: "leetcode_testuser",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "missing leetcode username",
			body: RegisterRequest{
				Username: "testuser",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			body, err := json.Marshal(tt.body)
			testutils.CheckErr(t, err, "Failed to marshal request body")

			req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.Register(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusCreated {
				exists, err := handler.store.CheckUserExistsByUsername(tt.body.Username)
				testutils.CheckErr(t, err, "Failed to check user existence")

				if !exists {
					t.Error("User not created in database")
				}
			}
		})
	}
}