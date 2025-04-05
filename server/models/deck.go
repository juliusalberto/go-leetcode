package models

import (
	"database/sql"
	// "fmt" // Removed unused import
	"time"

	"github.com/google/uuid"
)

func NewDeckStore(db *sql.DB) *DeckStore {
	return &DeckStore{db: db}
}

type Deck struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsPublic    bool      `json:"is_public"`
	UserID      string    `json:"user_id,omitempty"` // Nullable
}

type DeckProblem struct {
	DeckID    int       `json:"deck_id"`
	ProblemID int       `json:"problem_id"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

// DeckStore for database operations
type DeckStore struct {
	db *sql.DB
}

func (s *DeckStore) GetAllPublicDecks() ([]Deck, error) {
	query := `SELECT id, name, description, created_at, is_public FROM decks WHERE is_public = true`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err // Reverted error handling
	}
	defer rows.Close()

	var decks []Deck
	for rows.Next() {
		var deck Deck
		err := rows.Scan(&deck.ID, &deck.Name, &deck.Description, &deck.CreatedAt, &deck.IsPublic)
		if err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}
	return decks, nil
}

func (s *DeckStore) GetUserDecks(userID uuid.UUID) ([]Deck, error) {
	query := `SELECT id, name, description, created_at, is_public FROM decks WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err // Reverted error handling
	}
	defer rows.Close()

	var decks []Deck
	for rows.Next() {
		var deck Deck
		// Reverted scan
		err := rows.Scan(&deck.ID, &deck.Name, &deck.Description, &deck.CreatedAt, &deck.IsPublic)
		if err != nil {
			return nil, err // Reverted error handling
		}
		decks = append(decks, deck)
	}
	return decks, nil
}

func (s *DeckStore) GetDeckByID(deckID int) (Deck, error) {
	query := `SELECT id, name, description, created_at, is_public, user_id FROM decks WHERE id = $1`
	var deck Deck
	err := s.db.QueryRow(query, deckID).Scan(
		&deck.ID,
		&deck.Name,
		&deck.Description,
		&deck.CreatedAt,
		&deck.IsPublic,
		&deck.UserID,
	)
	if err != nil {
		return Deck{}, err
	}
	return deck, nil
}

func (s *DeckStore) CreateDeck(deck *Deck) error {
	query := `INSERT INTO decks (name, description, is_public, user_id) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	return s.db.QueryRow(query, deck.Name, deck.Description, deck.IsPublic, deck.UserID).Scan(&deck.ID, &deck.CreatedAt)
}

func (s *DeckStore) UpdateDeck(deck *Deck) error {
	query := `UPDATE decks SET name = $1, description = $2, is_public = $3 WHERE id = $4`
	_, err := s.db.Exec(query, deck.Name, deck.Description, deck.IsPublic, deck.ID)
	return err
}

func (s *DeckStore) DeleteDeck(deckID int, userID uuid.UUID) error {
	query := `DELETE FROM decks WHERE id = $1 AND user_id = $2`
	_, err := s.db.Exec(query, deckID, userID)
	return err
}

func (s *DeckStore) AddProblemToDeck(deckID int, problemID int) error {
	query := `INSERT INTO deck_problems (deck_id, problem_id) VALUES ($1, $2)`
	_, err := s.db.Exec(query, deckID, problemID)
	return err
}

func (s *DeckStore) RemoveProblemFromDeck(deckID, problemID int) error {
	query := `DELETE FROM deck_problems WHERE deck_id = $1 AND problem_id = $2`
	_, err := s.db.Exec(query, deckID, problemID)
	return err
}

// GetProblemsInDeck retrieves problems associated with a specific deck with pagination
func (s *DeckStore) GetProblemsInDeck(deckID int, limit int, offset int) ([]Problem, error) {
	query := `
		SELECT p.id, p.frontend_id, p.title, p.title_slug, p.difficulty, p.is_paid_only
		FROM problems p
		JOIN deck_problems dp ON p.id = dp.problem_id
		WHERE dp.deck_id = $1
		ORDER BY p.frontend_id -- Or another meaningful order
		LIMIT $2 OFFSET $3
	`
	rows, err := s.db.Query(query, deckID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []Problem
	for rows.Next() {
		var problem Problem
		err := rows.Scan(
			&problem.ID,
			&problem.FrontendID,
			&problem.Title,
			&problem.TitleSlug,
			&problem.Difficulty,
			&problem.IsPaidOnly,
		)
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	return problems, nil
}
