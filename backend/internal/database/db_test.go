package database

import (
	"database/sql"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

type TestDB struct {
	DB *sql.DB
}

func SetupTestDB(t *testing.T) *TestDB {
	if err := godotenv.Load("../../../.env.test"); err != nil {
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
	_, err := tdb.DB.Exec("TRUNCATE TABLE submissions, users, review_schedules CASCADE")
	if err != nil {
        t.Errorf("Could not drop test database: %v", err)
    }
	tdb.DB.Close()
}

func TestDBConnection(t *testing.T) {
	testDB := SetupTestDB(t)
	defer testDB.Cleanup(t)
}
