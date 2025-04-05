package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID
	Username         string `json:"username"`
	Email            string
	LeetcodeUsername string `json:"leetcode_username"`
	CreatedAt        time.Time
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) CreateUser(user *User) error {
	query := `
        INSERT INTO users
        (username, leetcode_username, email)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	// Use QueryRow with RETURNING to get the generated ID
	err := s.db.QueryRow(query, user.Username, user.LeetcodeUsername, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

// useful to prevent race condition
// imagine that the CreateUser is called twice because of network hiccups
// the second request will return an error (UUID alr used)
func (s *UserStore) CreateUserByAuth(user *User) error {
    if user.ID == uuid.Nil {
		return fmt.Errorf("cannot create user from auth without a valid ID")
	}

    fmt.Printf("DEBUG CreateUserByAuth: Creating user with UUID: %s, Username: %s, Email: %s\n",
               user.ID.String(), user.Username, user.Email)

	query := `
        INSERT INTO users
        (id, email, username, leetcode_username, created_at)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO NOTHING
    `
	result, err := s.db.Exec(query, user.ID, user.Email, user.Username, user.LeetcodeUsername, user.CreatedAt)
	if err != nil {
		fmt.Printf("ERROR CreateUserByAuth: Database error: %v\n", err)
		return fmt.Errorf("error creating user from auth: %w", err)
	}

    // Check if the insert actually happened or was ignored due to ON CONFLICT DO NOTHING
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        fmt.Printf("WARN CreateUserByAuth: No rows affected - user might already exist or insert failed\n")
    } else {
        fmt.Printf("SUCCESS CreateUserByAuth: User created with UUID: %s\n", user.ID.String())
    }

	return nil
}

func (s *UserStore) GetUserByID(id uuid.UUID) (User, error) {
	var user User

	query := `
		SELECT id, username, leetcode_username, created_at
		FROM users WHERE ID = $1
	`

	err := s.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.LeetcodeUsername, &user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with ID %s not found", id)
		}

		return User{}, fmt.Errorf("error fetching user: %v", err)
	}

	return user, nil
}

func (s *UserStore) GetUserByUsername(username string) (User, error) {
	var user User

	query := `
		SELECT id, username, leetcode_username, created_at
		FROM users WHERE username = $1
	`

	err := s.db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.LeetcodeUsername, &user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with username %s not found", username)
		}

		return User{}, fmt.Errorf("error fetching user: %v", err)
	}

	return user, nil
}

func (s *UserStore) CheckUserExistsByID(id uuid.UUID) (bool, error) {
	fmt.Printf("CheckUserExistsByID: Checking if user exists with UUID: %s\n", id.String())

	query := `
        SELECT EXISTS (
            SELECT 1
            FROM users
            WHERE id = $1
        )
    `

	var exists bool
	err := s.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		// Add specific debug log for the error
		fmt.Printf("DEBUG CheckUserExistsByID: Scan error for UUID %s: %v\n", id.String(), err)
		// Still return the wrapped error
		return false, fmt.Errorf("error checking user existence: %w", err) // Use %w for error wrapping
	}

	fmt.Printf("CheckUserExistsByID: User exists: %v\n", exists)

	// If user doesn't exist, let's see what users do exist in the database
	if !exists {
		debugQuery := `SELECT id, username FROM users LIMIT 5`
		rows, err := s.db.Query(debugQuery)
		if err != nil {
			fmt.Printf("DEBUG: Failed to query existing users: %v\n", err)
		} else {
			defer rows.Close()
			fmt.Println("DEBUG: First 5 users in database:")
			for rows.Next() {
				var dbID uuid.UUID
				var username string
				if err := rows.Scan(&dbID, &username); err == nil {
					fmt.Printf("  - ID: %s, Username: %s\n", dbID.String(), username)
				}
			}
		}
	}

	return exists, nil
}

func (s *UserStore) CheckUserExistsByUsername(username string) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM users
            WHERE username = $1
        )
    `

	var exists bool
	err := s.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %v", err)
	}

	return exists, nil
}

func (s *UserStore) CheckUserExistsByLeetcodeUsername(leetcodeUsername string) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM users
            WHERE leetcode_username = $1
        )
    `

	var exists bool
	err := s.db.QueryRow(query, leetcodeUsername).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence by leetcode username: %v", err)
	}

	return exists, nil
}

func (s *UserStore) GetUserByLeetcodeUsername(leetcodeUsername string) (User, error) {
	var user User

	query := `
        SELECT id, username, leetcode_username, created_at
        FROM users WHERE leetcode_username = $1
    `

	err := s.db.QueryRow(query, leetcodeUsername).Scan(
		&user.ID,
		&user.Username,
		&user.LeetcodeUsername,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with leetcode username %s not found", leetcodeUsername)
		}
		return User{}, fmt.Errorf("error fetching user by leetcode username: %v", err)
	}

	return user, nil
}
