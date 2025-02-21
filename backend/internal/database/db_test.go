package database

import (
	"testing"
	"database/sql"
)

type TestDB struct {
	DB *sql.DB
}

func SetupTestDB(t *testing.T) *TestDB {
	config := &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "yourpassword",
		DBName:   "leetcode_app",
	}

	db, err := sql.Open("postgres", "postgres://postgres:yourpassword@localhost:5432?sslmode=disable")
	if (err != nil) {
		t.Errorf("Failed to connect to database: %v", err)
	}

	_, err = db.Exec("CREATE DATABASE leetcode_app_test")
	if err != nil {
		t.Fatalf("Could not create a test database: %v", err)
	}

	testDB, err := NewConnection(config)
	if err != nil {
		t.Fatalf("Could not connect to test database: %v", err)
	}

	_, err = testDB.Exec(`
        CREATE TABLE submissions (
            id VARCHAR PRIMARY KEY,
            title VARCHAR NOT NULL,
            title_slug VARCHAR NOT NULL,
            submitted_at TIMESTAMP NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)

	if err != nil {
		t.Fatalf("Could not create tables: %v", err)
	}

	return &TestDB{DB: testDB}
}

func (tdb *TestDB) Cleanup(t *testing.T) {
	tdb.DB.Close()

	db, err := sql.Open("postgres", "postgres://postgres:yourpassword@localhost:5432?sslmode=disable")
	if err != nil {
		t.Errorf("Error in opening database: %v", err)
		return
	}

	defer db.Close()

	_, err = db.Exec("DROP DATABASE leetcode_app_test")
	if err != nil {
        t.Errorf("Could not drop test database: %v", err)
    }
}

func TestDBConnection(t *testing.T) {
	testDB := SetupTestDB(t)
	defer testDB.Cleanup(t)
}
