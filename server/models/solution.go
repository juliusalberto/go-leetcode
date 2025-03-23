package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Solution struct {
	ID           int       `json:"id"`
	ProblemID    int       `json:"problem_id"`
	Language     string    `json:"language"`
	SolutionCode string    `json:"solution_code"`
	CreatedAt    time.Time `json:"created_at"`
}

type SolutionStore struct {
	db *sql.DB
}

func NewSolutionStore(db *sql.DB) *SolutionStore {
	return &SolutionStore{db: db}
}

// Create a new solution
func (s *SolutionStore) CreateSolution(solution Solution) (Solution, error) {
	query := `
		INSERT INTO problem_solutions (problem_id, language, solution_code)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := s.db.QueryRow(
		query,
		solution.ProblemID,
		solution.Language,
		solution.SolutionCode,
	).Scan(&solution.ID, &solution.CreatedAt)

	if err != nil {
		return Solution{}, fmt.Errorf("error creating solution: %v", err)
	}

	return solution, nil
}

// Get solution by ID
func (s *SolutionStore) GetSolutionByID(id int) (Solution, error) {
	query := `
		SELECT id, problem_id, language, solution_code, created_at
		FROM problem_solutions
		WHERE id = $1
	`

	var solution Solution
	err := s.db.QueryRow(query, id).Scan(
		&solution.ID,
		&solution.ProblemID,
		&solution.Language,
		&solution.SolutionCode,
		&solution.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Solution{}, fmt.Errorf("solution with ID %d not found", id)
		}
		return Solution{}, fmt.Errorf("error fetching solution: %v", err)
	}

	return solution, nil
}

// Get solutions by problem ID
func (s *SolutionStore) GetSolutionsByProblemID(problemID int) ([]Solution, error) {
	query := `
		SELECT id, problem_id, language, solution_code, created_at
		FROM problem_solutions
		WHERE problem_id = $1
		ORDER BY language
	`

	rows, err := s.db.Query(query, problemID)
	if err != nil {
		return nil, fmt.Errorf("error querying solutions: %v", err)
	}
	defer rows.Close()

	var solutions []Solution
	for rows.Next() {
		var solution Solution
		err := rows.Scan(
			&solution.ID,
			&solution.ProblemID,
			&solution.Language,
			&solution.SolutionCode,
			&solution.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning solution row: %v", err)
		}
		solutions = append(solutions, solution)
	}

	return solutions, nil
}

// Get solution by problem ID and language
func (s *SolutionStore) GetSolutionByProblemAndLanguage(problemID int, language string) (Solution, error) {
	query := `
		SELECT id, problem_id, language, solution_code, created_at
		FROM problem_solutions
		WHERE problem_id = $1 AND lower(language) = $2
	`

	var solution Solution
	err := s.db.QueryRow(query, problemID, language).Scan(
		&solution.ID,
		&solution.ProblemID,
		&solution.Language,
		&solution.SolutionCode,
		&solution.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Solution{}, fmt.Errorf("solution for problem ID %d and language %s not found", problemID, language)
		}
		return Solution{}, fmt.Errorf("error fetching solution: %v", err)
	}

	return solution, nil
}

// Update a solution
func (s *SolutionStore) UpdateSolution(solution Solution) error {
	query := `
		UPDATE problem_solutions
		SET solution_code = $1
		WHERE id = $2
	`

	result, err := s.db.Exec(query, solution.SolutionCode, solution.ID)
	if err != nil {
		return fmt.Errorf("error updating solution: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("solution with ID %d not found", solution.ID)
	}

	return nil
}

// Delete a solution
func (s *SolutionStore) DeleteSolution(id int) error {
	query := "DELETE FROM problem_solutions WHERE id = $1"

	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting solution: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("solution with ID %d not found", id)
	}

	return nil
}
