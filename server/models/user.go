package models 

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
    ID        int
    Username  string
	LeetcodeUsername string
    CreatedAt time.Time
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
        (username, leetcode_username, created_at)
        VALUES ($1, $2, $3)
        RETURNING id
    `

    // Use QueryRow with RETURNING to get the generated ID
    err := s.db.QueryRow(query, user.Username, user.LeetcodeUsername, user.CreatedAt).Scan(&user.ID)
    if err != nil {
        return fmt.Errorf("error creating user: %v", err)
    }

    return nil
}

func (s *UserStore) GetUserByID(id int)(User, error) {
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
			return User{}, fmt.Errorf("user with ID %d not found", id)
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

func (s *UserStore) CheckUserExistsByID(id int)(bool, error) {
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
        return false, fmt.Errorf("error checking user existence: %v", err)
    }

    return exists, nil
}

func (s *UserStore) CheckUserExistsByUsername(username string)(bool, error) {
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


