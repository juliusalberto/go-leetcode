package models

import (
    "testing"
    "time"
    "go-leetcode/backend/internal/database"
    "go-leetcode/backend/internal/testutils"
)

func setupTestUser(t *testing.T) (*UserStore, *database.TestDB) {
    testDB := database.SetupTestDB(t)
    store := NewUserStore(testDB.DB)
    return store, testDB
}

func TestCreateUser(t *testing.T) {
    store, testDB := setupTestUser(t)
    defer testDB.Cleanup(t)

    user := User{
        Username:  "testuser",
        CreatedAt: time.Now(),
    }

    err := store.CreateUser(user)
    testutils.CheckErr(t, err, "Failed to create user")

    exists, err := store.CheckUserExistsByUsername("testuser")
    testutils.CheckErr(t, err, "Error checking user existence")
    if !exists {
        t.Error("Created user not found")
    }
}

func TestGetUserByID(t *testing.T) {
    store, testDB := setupTestUser(t)
    defer testDB.Cleanup(t)

    user := User{
        Username:  "testuser",
        CreatedAt: time.Now(),
    }
    err := store.CreateUser(user)
    testutils.CheckErr(t, err, "Failed to create test user")

    fetchedUser, err := store.GetUserByID(user.ID)
    testutils.CheckErr(t, err, "Failed to get user by ID")
    if fetchedUser.Username != user.Username {
        t.Errorf("Got wrong user: expected %s, got %s", user.Username, fetchedUser.Username)
    }
}

func TestGetUserByUsername(t *testing.T) {
    store, testDB := setupTestUser(t)
    defer testDB.Cleanup(t)

    user := User{
        Username:  "testuser",
        CreatedAt: time.Now(),
    }
    err := store.CreateUser(user)
    testutils.CheckErr(t, err, "Failed to create test user")

    fetchedUser, err := store.GetUserByUsername("testuser")
    testutils.CheckErr(t, err, "Failed to get user by username")
    if fetchedUser.Username != user.Username {
        t.Errorf("Got wrong user: expected %s, got %s", user.Username, fetchedUser.Username)
    }
}

func TestCheckUserExistsByID(t *testing.T) {
    store, testDB := setupTestUser(t)
    defer testDB.Cleanup(t)

    user := User{
        Username:  "testuser",
        CreatedAt: time.Now(),
    }
    err := store.CreateUser(user)
    testutils.CheckErr(t, err, "Failed to create test user")

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
    store, testDB := setupTestUser(t)
    defer testDB.Cleanup(t)

    // Create test user
    user := User{
        Username:  "testuser",
        CreatedAt: time.Now(),
    }
    err := store.CreateUser(user)
    testutils.CheckErr(t, err, "Failed to create test user")

    exists, err := store.CheckUserExistsByUsername("testuser")
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