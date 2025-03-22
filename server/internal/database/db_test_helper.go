package database

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
)

type TestDB struct {
	DB *sql.DB
}

func findProjectRoot(start string) (string, error) {
    current := start
    for {
        goModPath := filepath.Join(current, "go.mod")
        if info, err := os.Stat(goModPath); err == nil && !info.IsDir() {
            return current, nil
        }

        parent := filepath.Dir(current)
        if parent == current {
            // We are at the filesystem root and never found go.mod
            return "", errors.New("could not locate go.mod; are you sure you're in a Go project?")
        }
        current = parent
    }
}

func SetupTestDB(t *testing.T) *TestDB {
	_, thisFile, _, _ := runtime.Caller(0)
	dirOfThisFile := filepath.Dir(thisFile)

	root, err := findProjectRoot(dirOfThisFile)
	if err != nil {
		t.Fatalf("Error finding project root: %v", err)
	}

	envFilePath := filepath.Join(root, ".env.test")

	if err := godotenv.Load(envFilePath); err != nil {
		t.Fatalf("Error loading .env.test file: %v", err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	testDB, err := NewConnection(config)
	if err != nil {
		t.Fatalf("Could not connect to test database: %v", err)
	}


	return &TestDB{DB: testDB}
}

func (tdb *TestDB) Cleanup(t *testing.T) {
	_, err := tdb.DB.Exec("TRUNCATE TABLE submissions, users, review_schedules, review_logs RESTART IDENTITY CASCADE")
	if err != nil {
        t.Errorf("Could not truncate test database tables: %v", err)
    }
}