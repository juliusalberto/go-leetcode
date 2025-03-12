package handlers

import (
	"bytes"
	"encoding/json"
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"go-leetcode/backend/models"
	"go-leetcode/backend/pkg/response"
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

func setupUserTestWithUser(t *testing.T) (*UserHandler, *database.TestDB, *models.User) {
	testDB := database.SetupTestDB(t)
	userStore := models.NewUserStore(testDB.DB)
	handler := NewUserHandler(userStore)

	// Create a test user
	testUser := &models.User{
		Username:         "testuser",
		LeetcodeUsername: "leetcode_testuser",
	}

	err := userStore.CreateUser(testUser)
	testutils.CheckErr(t, err, "Failed to create test user")

	return handler, testDB, testUser
}

func TestRegisterHandler(t* testing.T) {
	handler, testDB := setupUserTest(t)
	defer testDB.Cleanup(t)

	tests := []struct {
		name 		string 
		body		RegisterRequest
		wantStatus	int
		checkErrors bool
	}{
		{
			name: "valid registration",
			body: RegisterRequest{
				Username: "testuser",
				LeetcodeUsername: "leetcode_testuser",
			},
			wantStatus: http.StatusCreated,
			checkErrors: false,
		},
		{
			name: "missing username",
			body: RegisterRequest{
				LeetcodeUsername: "leetcode_testuser",
			},
			wantStatus: http.StatusBadRequest,
			checkErrors: true,
		},
		{
			name: "missing leetcode username",
			body: RegisterRequest{
				Username: "testuser",
			},
			wantStatus: http.StatusBadRequest,
			checkErrors: true,
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

			// Decode standardized response
			var resp response.Response
			err = json.Unmarshal(rr.Body.Bytes(), &resp)
			testutils.CheckErr(t, err, "Failed to decode response")

			// For success cases, check if data exists
			if tt.wantStatus == http.StatusCreated {
				if resp.Data == nil {
					t.Error("Response data should not be nil for successful creation")
				}
				if len(resp.Errors) > 0 {
					t.Errorf("Response should not contain errors, got %v", resp.Errors)
				}

				exists, err := handler.store.CheckUserExistsByUsername(tt.body.Username)
				testutils.CheckErr(t, err, "Failed to check user existence")

				if !exists {
					t.Error("User not created in database")
				}
			}

			// For error cases, check for errors array
			if tt.checkErrors {
				if len(resp.Errors) == 0 {
					t.Error("Expected errors in response, but got none")
				}
			}
		})
	}
}

func TestGetUserHandler(t *testing.T) {
	handler, testDB, testUser := setupUserTestWithUser(t)
	defer testDB.Cleanup(t)

	// Test case: Valid username
	req := httptest.NewRequest("GET", "/api/users?username="+testUser.Username, nil)
	rr := httptest.NewRecorder()

	handler.GetUser(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode standardized response
	var resp response.Response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	testutils.CheckErr(t, err, "Failed to unmarshal response")

	// Check for errors
	if len(resp.Errors) > 0 {
		t.Errorf("Response contains errors: %v", resp.Errors)
	}

	// Convert data to User structure
	userData, err := json.Marshal(resp.Data)
	testutils.CheckErr(t, err, "Failed to marshal user data")

	var user models.User
	err = json.Unmarshal(userData, &user)
	testutils.CheckErr(t, err, "Failed to unmarshal user data")

	// Verify returned user data
	if user.Username != testUser.Username {
		t.Errorf("Expected username %s, got %s", testUser.Username, user.Username)
	}

	if user.LeetcodeUsername != testUser.LeetcodeUsername {
		t.Errorf("Expected leetcode_username %s, got %s", testUser.LeetcodeUsername, user.LeetcodeUsername)
	}

	// Test case: Non-existent username
	req = httptest.NewRequest("GET", "/api/users?username=nonexistent", nil)
	rr = httptest.NewRecorder()

	handler.GetUser(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status code %d for non-existent user, got %d", 
			http.StatusNotFound, status)
	}

	// Test case: Missing username parameter
	req = httptest.NewRequest("GET", "/api/users", nil)
	rr = httptest.NewRecorder()

	handler.GetUser(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %d for missing username, got %d", 
			http.StatusBadRequest, status)
	}
}