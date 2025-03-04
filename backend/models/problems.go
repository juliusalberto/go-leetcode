package models 

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type TopicTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	TranslatedName *string `json:"translatedName,omitempty"`
}

type SimilarQuestion struct {
	Title       string `json:"title"`
	TitleSlug   string `json:"titleSlug"`
	Difficulty  string `json:"difficulty"`
	TranslatedTitle *string `json:"translatedTitle,omitempty"`
}

type Problem struct {
	ID              string          `json:"id"`
	FrontendID      string          `json:"frontend_id"`
	Title           string          `json:"title"`
	TitleSlug       string          `json:"title_slug"`
	Difficulty      string          `json:"difficulty"`
	IsPaidOnly      bool            `json:"is_paid_only"`
	Content         string          `json:"content"`
	TopicTags       []TopicTag      `json:"topic_tags"`
	ExampleTestcases string         `json:"example_testcases"`
	SimilarQuestions []SimilarQuestion `json:"similar_questions"`
	CreatedAt       time.Time       `json:"created_at"`
}

type ProblemStore struct {
	db *sql.DB
}

func NewProblemStore(db *sql.DB) *ProblemStore {
	return &ProblemStore{db: db}
}

func (s *ProblemStore) GetProblemBySlug(titleSlug string)(Problem, error) {
	var problem Problem
	var similarQuestionsString, topicTagsString string

	query := `
		SELECT id, frontend_id, title, title_slug, difficulty, is_paid_only, content, topic_tags, 
				example_testcases, similar_questions, created_at FROM problems WHERE title_slug = $1
	`
	
	err := s.db.QueryRow(query, titleSlug).Scan(
        &problem.ID,
        &problem.FrontendID,
        &problem.Title,
        &problem.TitleSlug,
        &problem.Difficulty,
        &problem.IsPaidOnly,
        &problem.Content,
        &topicTagsString,
        &problem.ExampleTestcases,
        &similarQuestionsString,
        &problem.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Problem{}, fmt.Errorf("Problem with slug %s not found", titleSlug)
		}

		return Problem{}, fmt.Errorf("error fetching problem: %v", err)
	}

	if err = json.Unmarshal([]byte(topicTagsString), &problem.TopicTags); err != nil {
		return Problem{}, fmt.Errorf("error in parsing the topic tags")
	}

	if err = json.Unmarshal([]byte(similarQuestionsString), &problem.SimilarQuestions); err != nil {
		return Problem{}, fmt.Errorf("")
	}

	return problem, nil
}

func (s *ProblemStore) GetProblemByID (ID string)(Problem, error) {
	var problem Problem
	var similarQuestionsString, topicTagsString string

	query := `
		SELECT id, frontend_id, title, title_slug, difficulty, is_paid_only, content, topic_tags, 
				example_testcases, similar_questions, created_at FROM problems WHERE id = $1
	`
	
	err := s.db.QueryRow(query, ID).Scan(
        &problem.ID,
        &problem.FrontendID,
        &problem.Title,
        &problem.TitleSlug,
        &problem.Difficulty,
        &problem.IsPaidOnly,
        &problem.Content,
        &topicTagsString,
        &problem.ExampleTestcases,
        &similarQuestionsString,
        &problem.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Problem{}, fmt.Errorf("problem with ID %s not found", ID)
		}

		return Problem{}, fmt.Errorf("error fetching problem: %v", err)
	}

	if err = json.Unmarshal([]byte(topicTagsString), &problem.TopicTags); err != nil {
		return Problem{}, fmt.Errorf("error in parsing the topic tags")
	}

	if err = json.Unmarshal([]byte(similarQuestionsString), &problem.SimilarQuestions); err != nil {
		return Problem{}, fmt.Errorf("")
	}

	return problem, nil
}

func (s *ProblemStore) GetProblemByFrontendID (FrontendID string)(Problem, error) {
	var problem Problem
	var similarQuestionsString, topicTagsString string

	query := `
		SELECT id, frontend_id, title, title_slug, difficulty, is_paid_only, content, topic_tags, 
				example_testcases, similar_questions, created_at FROM problems WHERE frontend_id = $1
	`
	
	err := s.db.QueryRow(query, FrontendID).Scan(
        &problem.ID,
        &problem.FrontendID,
        &problem.Title,
        &problem.TitleSlug,
        &problem.Difficulty,
        &problem.IsPaidOnly,
        &problem.Content,
        &topicTagsString,
        &problem.ExampleTestcases,
        &similarQuestionsString,
        &problem.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Problem{}, fmt.Errorf("problem with ID %s not found", FrontendID)
		}

		return Problem{}, fmt.Errorf("error fetching problem: %v", err)
	}

	if err = json.Unmarshal([]byte(topicTagsString), &problem.TopicTags); err != nil {
		return Problem{}, fmt.Errorf("error in parsing the topic tags")
	}

	if err = json.Unmarshal([]byte(similarQuestionsString), &problem.SimilarQuestions); err != nil {
		return Problem{}, fmt.Errorf("")
	}

	return problem, nil
}