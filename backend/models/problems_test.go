package models

import (
	"go-leetcode/backend/internal/database"
	"go-leetcode/backend/internal/testutils"
	"strings"
	"testing"
)

func setupTestProblem(t* testing.T)(*ProblemStore, *database.TestDB) {
	testDB := database.SetupTestDB(t)
	store := NewProblemStore(testDB.DB)

	return store, testDB
}

func TestGetProblemByID(t *testing.T) {
	store, testDB := setupTestProblem(t)
	defer testDB.Cleanup(t)

	// check if we have the correct problem by id
	problem, err := store.GetProblemByID(1)
	testutils.CheckErr(t, err, "Failed to get problem")

	if (problem.TitleSlug != "two-sum" || problem.FrontendID != 1) {
		t.Errorf("Expected to get two sum, got: %v", problem)
	}

	// check another one
	problem, err = store.GetProblemByID(101)
	testutils.CheckErr(t, err, "Failed to get problem")
	if (problem.TitleSlug != "symmetric-tree" || problem.FrontendID != 101) {
		t.Errorf("Expected to get symmetric tree, got: %v", problem)
	}
}

func TestGetProblemByFrontendID(t* testing.T) {
	store, testDB := setupTestProblem(t)
	defer testDB.Cleanup(t)

	// check if we have the correct problem by id
	problem, err := store.GetProblemByFrontendID(1)
	testutils.CheckErr(t, err, "Failed to get problem")

	if (problem.TitleSlug != "two-sum" || problem.FrontendID != 1) {
		t.Errorf("Expected to get two sum, got: %v", problem)
	}

	// check another one
	problem, err = store.GetProblemByFrontendID(3271)
	testutils.CheckErr(t, err, "Failed to get problem")
	if (problem.TitleSlug != "hash-divided-string" || problem.ID != 3540) {
		t.Errorf("Expected to get hash divided string, got: %v", problem)
	}
}

func TestListProblems(t *testing.T) {
	// Connect to your existing test database
	store, testDB := setupTestProblem(t)
	
	// First get the total number of problems in the database
	var totalProblems int
	err := testDB.DB.QueryRow("SELECT COUNT(*) FROM problems").Scan(&totalProblems)
	if err != nil {
		t.Fatalf("Failed to count total problems: %v", err)
	}
	
	tests := []struct {
		name          string
		options       ListProblemOptions
		wantMinCount  int  // Minimum number of problems expected
		checkFirstProblem func(problem Problem) bool
	}{
		{
			name: "basic pagination",
			options: ListProblemOptions{
				Limit:  10,
				Offset: 0,
			},
			wantMinCount: 1, // At least one problem should exist
		},
		{
			name: "filter by easy difficulty",
			options: ListProblemOptions{
				Filter: ProblemFilter{
					Difficulty: "Easy",
				},
				Limit:  50,
				Offset: 0,
			},
			wantMinCount: 1,
			checkFirstProblem: func(p Problem) bool {
				return p.Difficulty == "Easy"
			},
		},
		{
			name: "filter by array tag",
			options: ListProblemOptions{
				Filter: ProblemFilter{
					Tags: []string{"array"},
				},
				Limit:  50,
				Offset: 0,
			},
			wantMinCount: 1,
			checkFirstProblem: func(p Problem) bool {
				for _, tag := range p.TopicTags {
					if tag.Slug == "array" {
						return true
					}
				}
				return false
			},
		},
		{
			name: "search for 'sum'",
			options: ListProblemOptions{
				Filter: ProblemFilter{
					SearchKeyword: "sum",
				},
				Limit:  50,
				Offset: 0,
			},
			wantMinCount: 1,
			checkFirstProblem: func(p Problem) bool {
				return strings.Contains(strings.ToLower(p.Title), "sum")
			},
		},
		{
			name: "order by difficulty desc",
			options: ListProblemOptions{
				OrderBy:  "difficulty",
				OrderDir: "desc",
				Limit:    10,
				Offset:   0,
			},
			wantMinCount: 1,
			checkFirstProblem: func(p Problem) bool {
				// Hard should come first in descending order
				return p.Difficulty == "Hard"
			},
		},
		{
			name: "order by frontend_id asc",
			options: ListProblemOptions{
				OrderBy:  "frontend_id",
				OrderDir: "asc",
				Limit:    10,
				Offset:   0,
			},
			wantMinCount: 1,
			checkFirstProblem: func(p Problem) bool {
				// Check that the frontend_id is numeric and low
				id := p.FrontendID
				return id < 100 // Assuming problem IDs start low
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := store.ListProblems(tt.options)
			if err != nil {
				t.Errorf("ListProblems() error = %v", err)
				return
			}
			
			if len(result.Problems) < tt.wantMinCount {
				t.Errorf("ListProblems() got %d problems, want at least %d", 
					len(result.Problems), tt.wantMinCount)
			}
			
			if result.Total < tt.wantMinCount {
				t.Errorf("ListProblems() total count %d, want at least %d", 
					result.Total, tt.wantMinCount)
			}
			
			if tt.checkFirstProblem != nil && len(result.Problems) > 0 {
				if !tt.checkFirstProblem(result.Problems[0]) {
					t.Errorf("ListProblems() first problem doesn't match expected condition: %+v", 
						result.Problems[0])
				}
			}
		})
	}
}