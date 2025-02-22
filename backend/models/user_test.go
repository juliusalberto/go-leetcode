package models

import (
    "testing"
    "time"
    "go-leetcode/backend/internal/database"
    "go-leetcode/backend/internal/testutils"
)

func setupTestUser(t *testing.T) (*UserStore, *database.TestDB, User) {
    testDB := database.SetupTestDB(t)
    store := NewUserStore(testDB.DB)
    
    testUser := User{
        Username:         "testuser",
        LeetcodeUsername: "leetcode_testuser",
        CreatedAt:       time.Now(),
    }
    
    err := store.CreateUser(&testUser)
    testutils.CheckErr(t, err, "Failed to create test user in setup")

    return store, testDB, testUser
}

func TestCreateUser(t *testing.T) {
    store, testDB, _ := setupTestUser(t)
    defer testDB.Cleanup(t)

    newUser := User{
        Username:         "newuser",
        LeetcodeUsername: "leetcode_newuser",
        CreatedAt:       time.Now(),
    }

    err := store.CreateUser(&newUser)
    testutils.CheckErr(t, err, "Failed to create user")

    exists, err := store.CheckUserExistsByUsername(newUser.Username)
    testutils.CheckErr(t, err, "Error checking user existence")
    if !exists {
        t.Error("Created user not found")
    }
}

func TestGetUserByID(t *testing.T) {
    store, testDB, user := setupTestUser(t)
    defer testDB.Cleanup(t)

    fetchedUser, err := store.GetUserByID(user.ID)
    testutils.CheckErr(t, err, "Failed to get user by ID")
    if fetchedUser.Username != user.Username {
        t.Errorf("Got wrong user: expected %s, got %s", user.Username, fetchedUser.Username)
    }

    // Test non-existent user
    _, err = store.GetUserByID(9999)
    if err == nil {
        t.Error("Expected error when getting non-existent user, got nil")
    }
}

func TestGetUserByUsername(t *testing.T) {
    store, testDB, user := setupTestUser(t)
    defer testDB.Cleanup(t)

    fetchedUser, err := store.GetUserByUsername(user.Username)
    testutils.CheckErr(t, err, "Failed to get user by username")
    if fetchedUser.Username != user.Username {
        t.Errorf("Got wrong user: expected %s, got %s", user.Username, fetchedUser.Username)
    }

    // Test non-existent user
    _, err = store.GetUserByUsername("nonexistent")
    if err == nil {
        t.Error("Expected error when getting non-existent user, got nil")
    }
}

func TestGetUserByLeetcodeUsername(t *testing.T) {
    store, testDB, user := setupTestUser(t)
    defer testDB.Cleanup(t)

    fetchedUser, err := store.GetUserByLeetcodeUsername(user.LeetcodeUsername)
    testutils.CheckErr(t, err, "Failed to get user by leetcode username")
    if fetchedUser.LeetcodeUsername != user.LeetcodeUsername {
        t.Errorf("Got wrong user: expected leetcode username %s, got %s", 
            user.LeetcodeUsername, fetchedUser.LeetcodeUsername)
    }

    // Test non-existent user
    _, err = store.GetUserByLeetcodeUsername("nonexistent")
    if err == nil {
        t.Error("Expected error when getting non-existent user, got nil")
    }
}

func TestCheckUserExistsByID(t *testing.T) {
    store, testDB, user := setupTestUser(t)
    defer testDB.Cleanup(t)

    exists, err := store.CheckUserExistsByID(user.ID)
    testutils.CheckErr(t, err, "Error checking user existence")
    if !exists {
        t.Error("User should exist but doesn't")
    }

    exists, err = store.CheckUserExistsByID(9999)
    testutils.CheckErr(t, err, "Error checking non-existent user")
    if exists {
        t.Error("User shouldn't exist but does")
    }
}

func TestCheckUserExistsByUsername(t *testing.T) {
    store, testDB, user := setupTestUser(t)
    defer testDB.Cleanup(t)

    exists, err := store.CheckUserExistsByUsername(user.Username)
    testutils.CheckErr(t, err, "Error checking user existence")
    if !exists {
        t.Error("User should exist but doesn't")
    }

    exists, err = store.CheckUserExistsByUsername("nonexistent")
    testutils.CheckErr(t, err, "Error checking non-existent user")
    if exists {
        t.Error("User shouldn't exist but does")
    }
}

func TestCheckUserExistsByLeetcodeUsername(t *testing.T) {
    store, testDB, user := setupTestUser(t)
    defer testDB.Cleanup(t)

    exists, err := store.CheckUserExistsByLeetcodeUsername(user.LeetcodeUsername)
    testutils.CheckErr(t, err, "Error checking user existence by leetcode username")
    if !exists {
        t.Error("User should exist but doesn't")
    }

    exists, err = store.CheckUserExistsByLeetcodeUsername("nonexistent")
    testutils.CheckErr(t, err, "Error checking non-existent user")
    if exists {
        t.Error("User shouldn't exist but does")
    }
}