package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TopicTag struct {
	Name           string  `json:"name"`
	Slug           string  `json:"slug"`
	TranslatedName *string `json:"translatedName,omitempty"`
}

type SimilarQuestion struct {
	Title           string  `json:"title"`
	TitleSlug       string  `json:"titleSlug"`
	Difficulty      string  `json:"difficulty"`
	TranslatedTitle *string `json:"translatedTitle,omitempty"`
}

type Problem struct {
	ID               int               `json:"id"`
	FrontendID       int               `json:"frontend_id"`
	Title            string            `json:"title"`
	TitleSlug        string            `json:"title_slug"`
	Difficulty       string            `json:"difficulty"`
	IsPaidOnly       bool              `json:"is_paid_only"`
	Content          string            `json:"content"`
	TopicTags        []TopicTag        `json:"topic_tags"`
	ExampleTestcases string            `json:"example_testcases"`
	SimilarQuestions []SimilarQuestion `json:"similar_questions"`
	CreatedAt        time.Time         `json:"created_at"`
	SolutionApproach string            `json:"solution_approach"`
}

type ProblemStore struct {
	db *sql.DB
}

func NewProblemStore(db *sql.DB) *ProblemStore {
	return &ProblemStore{db: db}
}

func (s *ProblemStore) GetProblemBySlug(titleSlug string) (Problem, error) {
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

func (s *ProblemStore) GetProblemByID(ID int) (Problem, error) {
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
			return Problem{}, fmt.Errorf("problem with ID %d not found", ID)
		}

		return Problem{}, fmt.Errorf("error fetching problem: %v", err)
	}

	if err = json.Unmarshal([]byte(topicTagsString), &problem.TopicTags); err != nil {
		return Problem{}, fmt.Errorf("error in parsing the topic tags")
	}

	if err = json.Unmarshal([]byte(similarQuestionsString), &problem.SimilarQuestions); err != nil {
		return Problem{}, fmt.Errorf("error in parsing the similar questions tag")
	}

	return problem, nil
}

func (s *ProblemStore) GetProblemByFrontendID(FrontendID int) (Problem, error) {
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
			return Problem{}, fmt.Errorf("problem with ID %d not found", FrontendID)
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

type ProblemFilter struct {
	Difficulty    string
	Tags          []string
	SearchKeyword string
	PaidOnly      *bool
}

type ListProblemOptions struct {
	Filter   ProblemFilter
	Limit    int
	Offset   int
	OrderBy  string // field to order by (e.g. difficulty)
	OrderDir string // asc or desc
}

type ProblemList struct {
	Problems []Problem `json:"problems"`
	Total    int       `json:"total"`
}

func (s *ProblemStore) ListProblems(options ListProblemOptions) (ProblemList, error) {
	baseQuery := `
		SELECT p.id, p.frontend_id, p.title, p.title_slug, p.difficulty, p.is_paid_only, p.content, p.topic_tags,
		p.example_testcases, p.similar_questions, p.created_at
		FROM problems p
		LEFT JOIN problems_topic pt ON p.id = pt.problem_id
		WHERE 1 = 1
	`

	countQuery := `SELECT COUNT(DISTINCT p.id) FROM problems p LEFT JOIN problems_topic pt ON p.id = pt.problem_id WHERE 1 = 1`

	// and then we add the query based on the filter
	var whereClause string
	var filterParams []interface{} // Parameters for WHERE clause (used in count)
	var params []interface{}       // All parameters including pagination (used in main query)
	paramPos := 1

	if options.Filter.Difficulty != "" {
		whereClause += fmt.Sprintf(" AND difficulty = $%d", paramPos)
		filterParams = append(filterParams, options.Filter.Difficulty) // Add to filterParams
		paramPos++
	}

	if options.Filter.PaidOnly != nil {
		whereClause += fmt.Sprintf(" AND is_paid_only = $%d", paramPos)
		filterParams = append(filterParams, *options.Filter.PaidOnly) // Add to filterParams
		paramPos++
	}

	if options.Filter.SearchKeyword != "" {
		whereClause += fmt.Sprintf(" AND title ILIKE $%d", paramPos)
		filterParams = append(filterParams, "%"+options.Filter.SearchKeyword+"%") // Add to filterParams
		paramPos++
	}

	if len(options.Filter.Tags) > 0 {
		tagConditions := []string{}
		for _, tag := range options.Filter.Tags {
			tagClause := fmt.Sprintf("pt.topic_slug = $%d", paramPos)
			tagConditions = append(tagConditions, tagClause)
			filterParams = append(filterParams, tag) // Add to filterParams
			paramPos++
		}
		whereClause += " AND (" + strings.Join(tagConditions, " OR ") + ")"
	}

	var orderClause string

	if options.OrderBy != "" {
		direction := "ASC"

		if options.OrderDir == "desc" {
			direction = "DESC"
		}

		if options.OrderBy == "difficulty" {
			// this is for difficulty
			// we want when we order by difficulty, the easy / hard one comes first
			// (depends on the desc or asc)

			orderClause = fmt.Sprintf(`
				ORDER BY 
                CASE 
                    WHEN LOWER(difficulty) = 'hard' THEN 3
                    WHEN LOWER(difficulty) = 'medium' THEN 2
                    WHEN LOWER(difficulty) = 'easy' THEN 1
                    ELSE 0
                END %s
			`, direction)
		} else {
			validColumns := map[string]string{
				"difficulty":  "difficulty",
				"title":       "title",
				"frontend_id": "frontend_id",
				"created_at":  "created_at",
			}

			if column, exists := validColumns[options.OrderBy]; exists {
				orderClause = fmt.Sprintf(" ORDER BY %s %s", column, direction)
			} else {
				orderClause = " ORDER BY frontend_id ASC"
			}
		}
	} else {
		orderClause = " ORDER BY frontend_id ASC"
	}

	// apply pagination

	// Prepare the full parameter list *after* filter params are finalized
	params = append(filterParams, options.Limit, options.Offset)

	// Use the correct parameter positions for LIMIT and OFFSET in the clause string
	limitParamPos := len(filterParams) + 1
	offsetParamPos := len(filterParams) + 2
	limitOffsetClause := fmt.Sprintf(" LIMIT $%d OFFSET $%d", limitParamPos, offsetParamPos)
	groupByClause := ` GROUP BY p.id, p.frontend_id, p.title, p.title_slug, p.difficulty, p.is_paid_only, p.content, p.topic_tags, p.example_testcases, p.similar_questions, p.created_at `

	query := baseQuery + whereClause + groupByClause + orderClause + limitOffsetClause
	countQuery = countQuery + whereClause

	var total int
	err := s.db.QueryRow(countQuery, filterParams...).Scan(&total) // Use filterParams

	if err != nil {
		// Also log the parameters used when the error occurs
		fmt.Printf("ERROR executing countQuery. Query: %s, Params: %v\n", countQuery, filterParams)
		return ProblemList{}, fmt.Errorf("error counting problems: %v", err)
	}

	var problems []Problem
	rows, err := s.db.Query(query, params...)

	if err != nil {
		// Also log the parameters used when the error occurs
		fmt.Printf("ERROR executing main query. Query: %s, Params: %v\n", query, params)
		return ProblemList{}, fmt.Errorf("error querying problems: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var problem Problem
		var topicTagsString, similarQuestionsString string

		err := rows.Scan(
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
			// Log the scan error
			fmt.Printf("ERROR scanning problems row: %v\n", err)
			return ProblemList{}, fmt.Errorf("error scanning problems row: %v", err)
		}

		// unmarshal the topic tags string
		err = json.Unmarshal([]byte(topicTagsString), &problem.TopicTags)

		if err != nil {
			// Log the unmarshal error
			fmt.Printf("ERROR unmarshaling topic tags for problem ID %d: %v\n", problem.ID, err)
			return ProblemList{}, fmt.Errorf("error unmarshaling topic tags: %v", err)
		}

		err = json.Unmarshal([]byte(similarQuestionsString), &problem.SimilarQuestions)
		if err != nil {
			// Log the unmarshal error
			fmt.Printf("ERROR unmarshaling similar questions for problem ID %d: %v\n", problem.ID, err)
			return ProblemList{}, fmt.Errorf("error unmarshaling similar questions: %v", err)
		}

		problems = append(problems, problem)
	}

	// now we already have the problems
	return ProblemList{
		Problems: problems,
		Total:    total,
	}, nil
}

type ProblemWithStatus struct {
	Problem   Problem `json:"problem"`
	Completed bool    `json:"completed"`
}

type ProblemListWithStatus struct {
	Problems []ProblemWithStatus `json:"problems"`
	Total    int                 `json:"total"`
}
