# Spaced Repetition Flashcard System Implementation Plan

## Overview

This plan outlines the implementation of a spaced repetition flashcard system for LeetCode problems, leveraging the existing FSRS (Free Spaced Repetition Scheduler) library. The system will help users improve long-term retention of problem-solving approaches by presenting problems with optimally spaced intervals. It will also organize problems into meaningful decks like Blind75, Neetcode150, etc.

## Database Changes

### 1. Add Solution Approach to Problems Table

```sql
-- Migration to add solution_approach column to problems table
ALTER TABLE problem_solutions ADD COLUMN solution_approach text DEFAULT NULL;
```

This will store a one-line, language-agnostic solution approach (e.g., "Dynamic Programming with state compression", "Binary Search with prefix sum").

### 2. Create Decks Tables

```sql
-- Create decks table
CREATE TABLE decks (
    id serial4 NOT NULL,
    name varchar NOT NULL,
    description text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_public boolean DEFAULT false,
    user_id uuid NULL, -- NULL for system decks like Blind75
    CONSTRAINT decks_pkey PRIMARY KEY (id),
    CONSTRAINT decks_name_user_unique UNIQUE (name, user_id)
);

-- Create deck_problems junction table
CREATE TABLE deck_problems (
    deck_id int4 NOT NULL,
    problem_id int4 NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT deck_problems_pkey PRIMARY KEY (deck_id, problem_id),
    CONSTRAINT fk_deck FOREIGN KEY (deck_id) REFERENCES decks(id) ON DELETE CASCADE,
    CONSTRAINT fk_problem FOREIGN KEY (problem_id) REFERENCES problems(id) ON DELETE CASCADE
);

-- Create index for faster lookups
CREATE INDEX idx_deck_problems_problem ON deck_problems(problem_id);
```

### 3. Create Flashcard Reviews Tables

```sql
-- Create flashcard_reviews table
CREATE TABLE flashcard_reviews (
    id serial4 NOT NULL,
    problem_id int4 NOT NULL,
    user_id uuid NOT NULL,
    deck_id int4 NOT NULL,
    next_review_at timestamp NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    stability float8 DEFAULT 0.0,
    difficulty float8 DEFAULT 0.0,
    elapsed_days int4 DEFAULT 0,
    scheduled_days int4 DEFAULT 0,
    reps int4 DEFAULT 0,
    lapses int4 DEFAULT 0,
    state int2 DEFAULT 0,
    last_review timestamp NULL,
    CONSTRAINT flashcard_reviews_pkey PRIMARY KEY (id),
    CONSTRAINT fk_flashcard_problem FOREIGN KEY (problem_id) REFERENCES problems(id),
    CONSTRAINT fk_flashcard_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_flashcard_deck FOREIGN KEY (deck_id) REFERENCES decks(id),
    CONSTRAINT unique_user_problem_deck UNIQUE (user_id, problem_id, deck_id)
);

-- Create flashcard_review_logs table
CREATE TABLE flashcard_review_logs (
    id serial4 NOT NULL,
    flashcard_review_id int4 NOT NULL,
    rating int2 NOT NULL,
    review_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    elapsed_days int4 NOT NULL,
    scheduled_days int4 NOT NULL,
    state int2 NOT NULL,
    CONSTRAINT flashcard_review_logs_pkey PRIMARY KEY (id),
    CONSTRAINT flashcard_review_logs_rating_check CHECK ((rating = ANY (ARRAY[1, 2, 3, 4]))),
    CONSTRAINT flashcard_review_logs_state_check CHECK ((state = ANY (ARRAY[0, 1, 2, 3]))),
    CONSTRAINT fk_flashcard_review FOREIGN KEY (flashcard_review_id) REFERENCES flashcard_reviews(id) ON DELETE CASCADE
);

-- Create index for faster lookups
CREATE INDEX idx_flashcard_reviews_user_time ON flashcard_reviews(user_id, next_review_at);
CREATE INDEX idx_flashcard_reviews_deck ON flashcard_reviews(deck_id);
```

## Backend Implementation

### 1. Create Models for Decks and Flashcards

```go
// models/deck.go
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

// models/flashcard_review.go
type FlashcardReview struct {
    ID            int       `json:"id"`
    ProblemID     int       `json:"problem_id"`
    UserID        string    `json:"user_id"`
    DeckID        int       `json:"deck_id"`
    NextReviewAt  time.Time `json:"next_review_at"`
    CreatedAt     time.Time `json:"created_at"`
    Stability     float64   `json:"stability"`
    Difficulty    float64   `json:"difficulty"`
    ElapsedDays   int       `json:"elapsed_days"`
    ScheduledDays int       `json:"scheduled_days"`
    Reps          int       `json:"reps"`
    Lapses        int       `json:"lapses"`
    State         int       `json:"state"`
    LastReview    time.Time `json:"last_review"`
}

type FlashcardReviewLog struct {
    ID                int       `json:"id"`
    FlashcardReviewID int       `json:"flashcard_review_id"`
    Rating            int       `json:"rating"`
    ReviewDate        time.Time `json:"review_date"`
    ElapsedDays       int       `json:"elapsed_days"`
    ScheduledDays     int       `json:"scheduled_days"`
    State             int       `json:"state"`
}

// Add conversion functions for FSRS similar to review_schedule.go
func ConvertFlashcardReviewToFSRS(review *FlashcardReview) fsrs.Card {
    return fsrs.Card{
        Due:           review.NextReviewAt,
        Stability:     review.Stability,
        Difficulty:    review.Difficulty,
        ElapsedDays:   review.ElapsedDays,
        ScheduledDays: review.ScheduledDays,
        Reps:          review.Reps,
        Lapses:        review.Lapses,
        State:         fsrs.State(review.State),
        LastReview:    review.LastReview,
    }
}

func ConvertFSRSToFlashcardReview(card fsrs.Card, review *FlashcardReview) {
    review.NextReviewAt = card.Due
    review.Stability = card.Stability
    review.Difficulty = card.Difficulty
    review.ElapsedDays = card.ElapsedDays
    review.ScheduledDays = card.ScheduledDays
    review.Reps = card.Reps
    review.Lapses = card.Lapses
    review.State = int(card.State)
}
```

### 2. Create Database Store Functions

```go
// models/deck_store.go

// GetAllPublicDecks returns all public decks
func (s *DeckStore) GetAllPublicDecks() ([]Deck, error) {
    // Query public decks
}

// GetUserDecks returns decks created by a specific user
func (s *DeckStore) GetUserDecks(userID uuid.UUID) ([]Deck, error) {
    // Query user's decks
}

// GetDeckByID returns a deck by its ID
func (s *DeckStore) GetDeckByID(deckID int) (Deck, error) {
    // Query specific deck
}

// CreateDeck creates a new deck
func (s *DeckStore) CreateDeck(deck *Deck) error {
    // Insert deck
}

// UpdateDeck updates an existing deck
func (s *DeckStore) UpdateDeck(deck *Deck) error {
    // Update deck
}

// DeleteDeck deletes a deck
func (s *DeckStore) DeleteDeck(deckID int, userID uuid.UUID) error {
    // Delete deck, ensuring it belongs to the user
}

// AddProblemToDeck adds a problem to a deck
func (s *DeckStore) AddProblemToDeck(deckID, problemID, position int) error {
    // Insert deck-problem relation
}

// RemoveProblemFromDeck removes a problem from a deck
func (s *DeckStore) RemoveProblemFromDeck(deckID, problemID int) error {
    // Delete deck-problem relation
}

// GetProblemsInDeck returns all problems in a deck
func (s *DeckStore) GetProblemsInDeck(deckID int) ([]models.Problem, error) {
    // Query problems in a deck with ordering
}

// models/flashcard_review_store.go

// GetDueFlashcardReviews returns flashcard reviews that are due for a user
func (s *FlashcardReviewStore) GetDueFlashcardReviews(userID uuid.UUID, deckID int, limit, offset int) ([]FlashcardReviewWithProblem, int, error) {
    // Query reviews that are due with problem details, filtered by deck if deckID > 0
}

// CreateFlashcardReview creates a new flashcard review record
func (s *FlashcardReviewStore) CreateFlashcardReview(review *FlashcardReview) error {
    // SQL INSERT logic
}

// UpdateFlashcardReview updates an existing flashcard review
func (s *FlashcardReviewStore) UpdateFlashcardReview(review *FlashcardReview) error {
    // SQL UPDATE logic
}

// CreateFlashcardReviewLog records a review log entry
func (s *FlashcardReviewStore) CreateFlashcardReviewLog(log *FlashcardReviewLog) error {
    // SQL INSERT logic for log
}

// AddDeckToUserFlashcards adds all problems from a deck to a user's flashcard queue
func (s *FlashcardReviewStore) AddDeckToUserFlashcards(userID uuid.UUID, deckID int) error {
    // Create flashcard reviews for all problems in the deck that aren't already in queue
}
```

### 3. Create API Handlers

```go
// api/handlers/deck.go
type DeckHandler struct {
    store        *models.DeckStore
    problemStore *models.ProblemStore
}

// GetAllDecks returns all accessible decks (public and user's own)
func (h *DeckHandler) GetAllDecks(w http.ResponseWriter, r *http.Request) {
    // Get user ID from context
    // Return public decks + user's own decks
}

// CreateDeck creates a new deck
func (h *DeckHandler) CreateDeck(w http.ResponseWriter, r *http.Request) {
    // Parse deck data and create
}

// UpdateDeck updates an existing deck
func (h *DeckHandler) UpdateDeck(w http.ResponseWriter, r *http.Request) {
    // Parse deck data and update
}

// DeleteDeck deletes a deck
func (h *DeckHandler) DeleteDeck(w http.ResponseWriter, r *http.Request) {
    // Delete the deck
}

// GetDeckProblems gets problems from a deck
func (h *DeckHandler) GetDeckProblems(w http.ResponseWriter, r *http.Request) {
    // Return problems in a deck
}

// AddProblemToDeck adds a problem to a deck
func (h *DeckHandler) AddProblemToDeck(w http.ResponseWriter, r *http.Request) {
    // Add problem to deck
}

// RemoveProblemFromDeck removes a problem from a deck
func (h *DeckHandler) RemoveProblemFromDeck(w http.ResponseWriter, r *http.Request) {
    // Remove problem from deck
}

// api/handlers/flashcard.go
type FlashcardHandler struct {
    store           *models.FlashcardReviewStore
    problemStore    *models.ProblemStore
    deckStore       *models.DeckStore
}

// GetFlashcardReviews returns due flashcard reviews for the authenticated user
func (h *FlashcardHandler) GetFlashcardReviews(w http.ResponseWriter, r *http.Request) {
    // Get userID from context
    // Parse deckID from query params (optional)
    // Query for due flashcard reviews, filtered by deck if provided
    // Return with problem details
}

// SubmitFlashcardReview processes a flashcard review submission with rating
func (h *FlashcardHandler) SubmitFlashcardReview(w http.ResponseWriter, r *http.Request) {
    // Parse review ID and rating
    // Use FSRS library to calculate next review date
    // Update review record and create log
    // Return updated schedule
}

// AddDeckToFlashcards adds all problems from a deck to user's flashcard queue
func (h *FlashcardHandler) AddDeckToFlashcards(w http.ResponseWriter, r *http.Request) {
    // Add all problems from the specified deck to the user's flashcard queue
}

// GetFlashcardStats returns statistics about flashcard usage
func (h *FlashcardHandler) GetFlashcardStats(w http.ResponseWriter, r *http.Request) {
    // Return stats like reviews completed, streak, etc.
    // Include deck-specific stats if deckID is provided
}
```

### 4. Add Routes

```go
// api/routes/routes.go
// Inside SetupRoutes function

// Deck routes
r.GET("/api/decks", handlers.GetAllDecks)
r.POST("/api/decks", middleware.Auth(handlers.CreateDeck))
r.PUT("/api/decks/:id", middleware.Auth(handlers.UpdateDeck))
r.DELETE("/api/decks/:id", middleware.Auth(handlers.DeleteDeck))
r.GET("/api/decks/:id/problems", handlers.GetDeckProblems)
r.POST("/api/decks/:id/problems", middleware.Auth(handlers.AddProblemToDeck))
r.DELETE("/api/decks/:id/problems/:problem_id", middleware.Auth(handlers.RemoveProblemFromDeck))

// Flashcard routes
r.GET("/api/flashcards", middleware.Auth(handlers.GetFlashcardReviews))
r.POST("/api/flashcards/:id/review", middleware.Auth(handlers.SubmitFlashcardReview))
r.POST("/api/flashcards/deck/:id", middleware.Auth(handlers.AddDeckToFlashcards))
r.GET("/api/flashcards/stats", middleware.Auth(handlers.GetFlashcardStats))
```

### 5. Seed Script for Default Decks

```go
// internal/seed/decks.go
func SeedDefaultDecks(db *sql.DB) error {
    // Create Blind75 deck
    blind75 := models.Deck{
        Name:        "Blind 75",
        Description: "A carefully curated list of 75 LeetCode questions covering key concepts",
        IsPublic:    true,
    }
    
    // Create Neetcode 150 deck
    neetcode150 := models.Deck{
        Name:        "Neetcode 150",
        Description: "An expanded collection of 150 carefully selected problems",
        IsPublic:    true,
    }
    
    // Insert decks and problems
    // ...
    
    return nil
}
```

## OpenRouter/DeepSeek v3 Integration

### 1. Script to Generate Solution Approaches

```python
# generate_solution_approaches.py
import openai
import psycopg2
import json
import os
from dotenv import load_dotenv

load_dotenv()

# Configure OpenRouter API
openai.api_base = "https://openrouter.ai/api/v1"
openai.api_key = os.getenv("OPENROUTER_API_KEY")

# Connect to database
conn = psycopg2.connect(os.getenv("DATABASE_URL"))
cur = conn.cursor()

def get_problems_without_approach():
    """Retrieve problems that don't have a solution approach yet"""
    cur.execute("""
        SELECT p.id, p.title, p.content, ps.solution_code, ps.language 
        FROM problems p
        LEFT JOIN problem_solutions ps ON p.id = ps.problem_id
        WHERE p.solution_approach IS NULL
        LIMIT 100;  # Process in batches
    """)
    return cur.fetchall()

def analyze_problem(problem_id, title, content, solution_code, language):
    """Use DeepSeek to analyze the problem and generate a solution approach"""
    prompt = f"""
    You are an expert LeetCode problem analyzer. Below is a problem and its solution.
    
    PROBLEM TITLE: {title}
    
    PROBLEM DESCRIPTION:
    {content}
    
    SOLUTION ({language}):
    {solution_code}
    
    Please analyze this solution and provide a one-line summary of the approach used.
    Focus on the key algorithmic technique or data structure (e.g., "Dynamic Programming with state compression", 
    "Binary Search with two pointers", "BFS traversal with visited set").
    
    Respond with just the one-line solution approach, nothing else.
    """
    
    try:
        response = openai.ChatCompletion.create(
            model="deepseek/deepseek-coder-v3", 
            messages=[{"role": "user", "content": prompt}],
            max_tokens=100
        )
        
        approach = response.choices[0].message.content.strip()
        print(f"Generated approach for problem {problem_id}: {approach}")
        return approach
    except Exception as e:
        print(f"Error analyzing problem {problem_id}: {e}")
        return None

def update_solution_approach(problem_id, approach):
    """Update the database with the generated solution approach"""
    cur.execute(
        "UPDATE problems SET solution_approach = %s WHERE id = %s",
        (approach, problem_id)
    )
    conn.commit()

def main():
    problems = get_problems_without_approach()
    print(f"Found {len(problems)} problems without solution approaches")
    
    for problem in problems:
        problem_id, title, content, solution_code, language = problem
        approach = analyze_problem(problem_id, title, content, solution_code, language)
        
        if approach:
            update_solution_approach(problem_id, approach)
    
    print("Process completed")

if __name__ == "__main__":
    main()
    cur.close()
    conn.close()
```

## Frontend Implementation

### 1. Create Deck and Flashcard Services

```typescript
// services/api/decks.ts
import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { getApiUrl } from '../../utils/apiUrl';

export interface Deck {
  id: number;
  name: string;
  description: string;
  is_public: boolean;
  created_at: string;
  user_id?: string;
  problem_count?: number;
}

export interface DeckProblem {
  id: number;
  title: string;
  difficulty: string;
  title_slug: string;
  position: number;
}

// Get all accessible decks (public + user's own)
export const useDecks = () => {
  const { get } = useFetchWithAuth();
  
  return useQuery({
    queryKey: ['decks'],
    queryFn: async () => {
      const url = getApiUrl(`http://localhost:8080/api/decks`);
      const response = await get(url);
      return response?.data || [];
    }
  });
};

// Get problems in a deck
export const useDeckProblems = (deckId: number) => {
  const { get } = useFetchWithAuth();
  
  return useQuery({
    queryKey: ['deckProblems', deckId],
    queryFn: async () => {
      const url = getApiUrl(`http://localhost:8080/api/decks/${deckId}/problems`);
      const response = await get(url);
      return response?.data || [];
    },
    enabled: !!deckId
  });
};

// Create a new deck
export const useCreateDeck = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async (deck: Omit<Deck, 'id' | 'created_at'>) => {
      const url = getApiUrl(`http://localhost:8080/api/decks`);
      return await post(url, deck);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['decks'] });
    }
  });
};

// Add deck to flashcards
export const useAddDeckToFlashcards = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async (deckId: number) => {
      const url = getApiUrl(`http://localhost:8080/api/flashcards/deck/${deckId}`);
      return await post(url, {});
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['flashcardReviews'] });
      queryClient.invalidateQueries({ queryKey: ['flashcardStats'] });
    }
  });
};

// services/api/flashcards.ts
import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { getApiUrl } from '../../utils/apiUrl';

export interface FlashcardReview {
  id: number;
  problem_id: number;
  deck_id: number;
  deck_name: string;
  title: string;
  title_slug: string;
  difficulty: string;
  content: string;
  solution_approach: string;
  next_review_at: string;
  solutions: Record<string, string>; // Language -> solution code
}

export interface FlashcardRating {
  review_id: number;
  rating: 1 | 2 | 3 | 4; // 1=Very Hard, 2=Hard, 3=Good, 4=Easy
}

export interface FlashcardStats {
  total: number;
  due: number;
  completed_today: number;
  decks: {
    deck_id: number;
    deck_name: string;
    total: number;
    due: number;
  }[];
}

// Get due flashcard reviews, optionally filtered by deck
export const useFlashcardReviews = (deckId?: number) => {
  const { get } = useFetchWithAuth();
  
  const fetchFlashcards = async ({ pageParam = 1 }) => {
    const url = getApiUrl(
      `http://localhost:8080/api/flashcards?page=${pageParam}${deckId ? `&deck_id=${deckId}` : ''}`
    );
    const response = await get(url);
    return response?.data || [];
  };
  
  return useInfiniteQuery({
    queryKey: ['flashcardReviews', deckId],
    queryFn: ({ pageParam }) => fetchFlashcards({ pageParam }),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => 
      lastPage.length ? allPages.length + 1 : undefined
  });
};

// Submit flashcard rating
export const useSubmitFlashcardRating = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async (rating: FlashcardRating) => {
      const url = getApiUrl(`http://localhost:8080/api/flashcards/${rating.review_id}/review`);
      return await post(url, { rating: rating.rating });
    },
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({ queryKey: ['flashcardReviews'] });
      // Update flashcard stats
      queryClient.invalidateQueries({ queryKey: ['flashcardStats'] });
    }
  });
};

// Get flashcard stats
export const useFlashcardStats = () => {
  const { get } = useFetchWithAuth();
  
  return useQuery({
    queryKey: ['flashcardStats'],
    queryFn: async () => {
      const url = getApiUrl(`http://localhost:8080/api/flashcards/stats`);
      const response = await get(url);
      return response?.data || { total: 0, due: 0, completed_today: 0, decks: [] };
    }
  });
};
```

### 2. Create Deck List Screen

```tsx
// app/decks.tsx
import React from 'react';
import { View, Text, FlatList, TouchableOpacity, ActivityIndicator, Alert } from 'react-native';
import { useDecks, useAddDeckToFlashcards, Deck } from '../services/api/decks';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';

export default function DecksScreen() {
  const { data: decks = [], isLoading } = useDecks();
  const addDeckToFlashcards = useAddDeckToFlashcards();
  
  const handleAddToFlashcards = (deck: Deck) => {
    Alert.alert(
      "Add to Flashcards",
      `Add all problems from "${deck.name}" to your flashcard queue?`,
      [
        { text: "Cancel", style: "cancel" },
        { 
          text: "Add", 
          onPress: () => {
            addDeckToFlashcards.mutate(deck.id, {
              onSuccess: () => {
                Alert.alert(
                  "Success",
                  `Added ${deck.name} to your flashcard queue!`,
                  [{ text: "OK" }]
                );
              },
              onError: (error) => {
                Alert.alert(
                  "Error",
                  `Failed to add deck: ${error}`,
                  [{ text: "OK" }]
                );
              }
            });
          }
        }
      ]
    );
  };
  
  const renderDeckItem = ({ item: deck }: { item: Deck }) => (
    <View className="bg-[#1E2A3A] rounded-lg p-4 mb-4">
      <View className="flex-row justify-between items-center mb-2">
        <Text className="text-[#F8F9FB] text-xl font-bold">{deck.name}</Text>
        <View className="flex-row">
          <TouchableOpacity 
            className="p-2"
            onPress={() => router.push(`/deck/${deck.id}`)}
          >
            <Ionicons name="eye-outline" size={24} color="#6366F1" />
          </TouchableOpacity>
          <TouchableOpacity 
            className="p-2"
            onPress={() => handleAddToFlashcards(deck)}
          >
            <Ionicons name="add-circle-outline" size={24} color="#4CD137" />
          </TouchableOpacity>
        </View>
      </View>
      <Text className="text-[#8A9DC0] mb-2">{deck.description}</Text>
      <View className="flex-row justify-between items-center">
        <Text className="text-[#6366F1]">
          {deck.problem_count || 0} problems
        </Text>
        <Text className="text-[#8A9DC0] text-sm">
          {deck.is_public ? 'Public Deck' : 'Your Deck'}
        </Text>
      </View>
    </View>
  );
  
  if (isLoading) {
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24]">
        <ActivityIndicator size="large" color="#6366F1" />
      </View>
    );
  }
  
  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center justify-between p-4">
        <Text className="text-[#F8F9FB] text-xl font-bold">Problem Decks</Text>
        <TouchableOpacity 
          className="bg-[#6366F1] p-2 rounded-full"
          onPress={() => router.push('/create-deck')}
        >
          <Ionicons name="add" size={24} color="#F8F9FB" />
        </TouchableOpacity>
      </View>
      
      <FlatList
        data={decks}
        renderItem={renderDeckItem}
        keyExtractor={(item) => item.id.toString()}
        contentContainerStyle={{ padding: 16 }}
        ListEmptyComponent={
          <View className="flex-1 justify-center items-center p-4">
            <Text className="text-[#8A9DC0] text-center">
              No decks found. Create your first deck!
            </Text>
          </View>
        }
      />
    </View>
  );
}
```

### 3. Create Deck Detail Screen

```tsx
// app/deck/[id].tsx
import React from 'react';
import { View, Text, FlatList, TouchableOpacity, ActivityIndicator } from 'react-native';
import { useDeckProblems, useAddDeckToFlashcards } from '../../services/api/decks';
import { Ionicons } from '@expo/vector-icons';
import { router, useLocalSearchParams } from 'expo-router';

export default function DeckDetailScreen() {
  const { id } = useLocalSearchParams();
  const deckId = parseInt(id as string);
  
  const { data: problems = [], isLoading } = useDeckProblems(deckId);
  const addDeckToFlashcards = useAddDeckToFlashcards();
  
  const handleAddToFlashcards = () => {
    addDeckToFlashcards.mutate(deckId);
  };
  
  const renderProblemItem = ({ item }) => (
    <TouchableOpacity 
      className="bg-[#1E2A3A] p-4 rounded-lg mb-2"
      onPress={() => router.push(`/problem/${item.title_slug}`)}
    >
      <View className="flex-row justify-between items-center">
        <Text className="text-[#F8F9FB] font-medium flex-1">
          {item.title}
        </Text>
        <View className="ml-2">
          <Text className={`px-2 py-1 rounded-full text-xs ${
            item.difficulty === 'Easy' ? 'bg-[#4CD137]/20 text-[#4CD137]' :
            item.difficulty === 'Medium' ? 'bg-[#F39C12]/20 text-[#F39C12]' :
            'bg-[#E74C3C]/20 text-[#E74C3C]'
          }`}>
            {item.difficulty}
          </Text>
        </View>
      </View>
    </TouchableOpacity>
  );
  
  if (isLoading) {
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24]">
        <ActivityIndicator size="large" color="#6366F1" />
      </View>
    );
  }
  
  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center p-4">
        <TouchableOpacity onPress={() => router.back()}>
          <Ionicons name="chevron-back" size={24} color="#F8F9FB" />
        </TouchableOpacity>
        <Text className="text-[#F8F9FB] text-xl font-bold ml-4 flex-1">
          Deck Details
        </Text>
        <TouchableOpacity 
          className="bg-[#6366F1] px-4 py-2 rounded-full"
          onPress={handleAddToFlashcards}
        >
          <Text className="text-white font-medium">Add to Flashcards</Text>
        </TouchableOpacity>
      </View>
      
      <FlatList
        data={problems}
        renderItem={renderProblemItem}
        keyExtractor={(item) => item.id.toString()}
        contentContainerStyle={{ padding: 16 }}
        ListHeaderComponent={
          <View className="mb-4">
            <Text className="text-[#F8F9FB] text-2xl font-bold mb-1">
              {problems[0]?.deck_name || 'Problems'}
            </Text>
            <Text className="text-[#8A9DC0] mb-4">
              {problems.length} problems in this deck
            </Text>
          </View>
        }
        ListEmptyComponent={
          <View className="flex-1 justify-center items-center p-4">
            <Text className="text-[#8A9DC0] text-center">
              No problems in this deck yet.
            </Text>
          </View>
        }
      />
    </View>
  );
}
```

### 4. Update Flashcard Review Screen

```tsx
// app/flashcards.tsx
import React, { useState } from 'react';
import { View, Text, TouchableOpacity, ScrollView, ActivityIndicator } from 'react-native';
import { useFlashcardReviews, useSubmitFlashcardRating, useFlashcardStats } from '../services/api/flashcards';
import { useDecks } from '../services/api/decks';
import { WebView } from 'react-native-webview';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';
import SyntaxHighlighter from 'react-native-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';

export default function FlashcardsScreen() {
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [showSolutionApproach, setShowSolutionApproach] = useState(false);
  const [showFullSolution, setShowFullSolution] = useState(false);
  const [reviewCompleted, setReviewCompleted] = useState(false);
  const [selectedLanguage, setSelectedLanguage] = useState('python');
  const [selectedDeckId, setSelectedDeckId] = useState<number | undefined>(undefined);
  
  // Fetch decks for the selector
  const { data: decks = [] } = useDecks();
  
  // Fetch flashcard reviews, filtered by deck if selected
  const { 
    data, 
    isLoading, 
    refetch 
  } = useFlashcardReviews(selectedDeckId);
  
  const { data: stats } = useFlashcardStats();
  const submitRating = useSubmitFlashcardRating();
  
  // Combine all pages of data
  const flashcards = data?.pages.flatMap(page => page) || [];
  const currentCard = flashcards[currentCardIndex];
  
  // Handler for rating submission
  const handleRating = async (rating: 1 | 2 | 3 | 4) => {
    if (!currentCard) return;
    
    try {
      await submitRating.mutateAsync({
        review_id: currentCard.id,
        rating
      });
      
      // Move to next card or show completion screen
      if (currentCardIndex < flashcards.length - 1) {
        setCurrentCardIndex(prev => prev + 1);
        setShowSolutionApproach(false);
        setShowFullSolution(false);
      } else {
        setReviewCompleted(true);
      }
    } catch (error) {
      console.error('Error submitting rating:', error);
      // Show error toast
    }
  };
  
  // Switch to a different deck
  const handleDeckChange = (deckId: number | undefined) => {
    setSelectedDeckId(deckId);
    setCurrentCardIndex(0);
    setShowSolutionApproach(false);
    setShowFullSolution(false);
    setReviewCompleted(false);
  };
  
  // Loading state
  if (isLoading) {
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24]">
        <ActivityIndicator size="large" color="#6366F1" />
      </View>
    );
  }
  
  // All reviews completed state
  if (reviewCompleted || flashcards.length === 0) {
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24] p-6">
        <Ionicons name="checkmark-circle" size={64} color="#4CD137" />
        <Text className="text-[#F8F9FB] text-2xl font-bold mt-4 text-center">
          All Done!
        </Text>
        <Text className="text-[#8A9DC0] text-base mt-2 text-center">
          You've completed all your flashcard reviews for now.
        </Text>
        <View className="flex-row mt-8">
          <TouchableOpacity 
            className="bg-[#6366F1] px-6 py-3 rounded-full mr-2"
            onPress={() => router.push('/decks')}
          >
            <Text className="text-white font-medium">Browse Decks</Text>
          </TouchableOpacity>
          <TouchableOpacity 
            className="bg-[#29374C] px-6 py-3 rounded-full"
            onPress={() => {
              setReviewCompleted(false);
              setCurrentCardIndex(0);
              refetch();
            }}
          >
            <Text className="text-[#F8F9FB] font-medium">Try Again</Text>
          </TouchableOpacity>
        </View>
      </View>
    );
  }
  
  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center p-4 justify-between">
        <TouchableOpacity onPress={() => router.back()}>
          <Ionicons name="chevron-back" size={24} color="#F8F9FB" />
        </TouchableOpacity>
        <Text className="text-[#F8F9FB] text-lg font-bold">
          Flashcard Review
        </Text>
        <Text className="text-[#8A9DC0]">
          {currentCardIndex + 1}/{flashcards.length}
        </Text>
      </View>
      
      {/* Deck selector */}
      <View className="px-4 mb-4">
        <ScrollView 
          horizontal 
          showsHorizontalScrollIndicator={false}
          contentContainerStyle={{ paddingRight: 20 }}
        >
          <TouchableOpacity 
            className={`px-4 py-2 rounded-full mr-2 ${
              selectedDeckId === undefined ? 'bg-[#6366F1]' : 'bg-[#29374C]'
            }`}
            onPress={() => handleDeckChange(undefined)}
          >
            <Text className={`${
              selectedDeckId === undefined ? 'text-white' : 'text-[#8A9DC0]'
            }`}>
              All Decks
            </Text>
          </TouchableOpacity>
          
          {stats?.decks.map(deck => (
            <TouchableOpacity 
              key={deck.deck_id}
              className={`px-4 py-2 rounded-full mr-2 ${
                selectedDeckId === deck.deck_id ? 'bg-[#6366F1]' : 'bg-[#29374C]'
              }`}
              onPress={() => handleDeckChange(deck.deck_id)}
            >
              <Text className={`${
                selectedDeckId === deck.deck_id ? 'text-white' : 'text-[#8A9DC0]'
              }`}>
                {deck.deck_name} ({deck.due})
              </Text>
            </TouchableOpacity>
          ))}
        </ScrollView>
      </View>
      
      {/* Flashcard Content */}
      <ScrollView className="flex-1 px-4">
        {/* Deck indicator */}
        <View className="mb-2">
          <Text className="text-[#6366F1] text-sm">
            {currentCard.deck_name}
          </Text>
        </View>
        
        {/* Problem Title and Difficulty */}
        <View className="mb-4">
          <Text className="text-[#F8F9FB] text-2xl font-bold">
            {currentCard.title}
          </Text>
          <View className="flex-row mt-2">
            <Text className={`text-sm font-medium px-2 py-1 rounded-full ${
              currentCard.difficulty === 'Easy' ? 'bg-[#4CD137]/20 text-[#4CD137]' :
              currentCard.difficulty === 'Medium' ? 'bg-[#F39C12]/20 text-[#F39C12]' :
              'bg-[#E74C3C]/20 text-[#E74C3C]'
            }`}>
              {currentCard.difficulty}
            </Text>
          </View>
        </View>
        
        {/* Problem Description */}
        <View className="bg-[#1E2A3A] rounded-lg p-4 mb-6">
          <WebView
            originWhitelist={['*']}
            source={{ 
              html: `
              <html>
              <head>
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <style>
                body {
                  font-family: -apple-system, sans-serif;
                  padding: 0;
                  margin: 0;
                  color: #F8F9FB;
                  background-color: #1E2A3A;
                  font-size: 16px;
                  line-height: 1.5;
                }
                p { margin-bottom: 16px; }
                code {
                  font-family: monospace;
                  background-color: #29374C;
                  padding: 2px 4px;
                  border-radius: 4px;
                }
                pre {
                  background-color: #29374C;
                  padding: 16px;
                  border-radius: 8px;
                  overflow-x: auto;
                  font-family: monospace;
                }
                </style>
              </head>
              <body>
                ${currentCard.content || ''}
              </body>
              </html>
              `
            }}
            style={{ 
              height: 300,
              backgroundColor: '#1E2A3A' 
            }}
            onMessage={(event) => {
              // Adjust height dynamically if needed
            }}
          />
        </View>
        
        {/* Solution Approach Button */}
        {!showSolutionApproach ? (
          <TouchableOpacity 
            className="bg-[#29374C] p-4 rounded-lg mb-4"
            onPress={() => setShowSolutionApproach(true)}
          >
            <Text className="text-[#F8F9FB] text-center font-medium">
              Reveal Solution Approach
            </Text>
          </TouchableOpacity>
        ) : (
          <View className="bg-[#29374C] p-4 rounded-lg mb-4">
            <Text className="text-[#F8F9FB] font-medium mb-2">
              Solution Approach:
            </Text>
            <Text className="text-[#6366F1] font-bold">
              {currentCard.solution_approach || "Dynamic Programming with state compression"}
            </Text>
          </View>
        )}
        
        {/* Full Solution Section */}
        {showSolutionApproach && (
          <>
            {/* Language Selection Tabs */}
            <View className="flex-row flex-wrap mb-4">
              {Object.keys(currentCard.solutions || {}).length > 0 ? (
                Object.keys(currentCard.solutions).map(lang => (
                  <TouchableOpacity
                    key={lang}
                    onPress={() => setSelectedLanguage(lang)}
                    className={`mr-2 mb-2 px-4 py-2 rounded-full ${
                      selectedLanguage === lang ? 'bg-[#29374C]' : 'bg-[#1E2A3A]'
                    }`}
                  >
                    <Text 
                      className="text-[#F8F9FB] text-base capitalize"
                    >
                      {lang === 'cpp' ? 'C++' : lang.charAt(0).toUpperCase() + lang.slice(1)}
                    </Text>
                  </TouchableOpacity>
                ))
              ) : (
                ['python', 'cpp'].map(lang => (
                  <TouchableOpacity
                    key={lang}
                    onPress={() => setSelectedLanguage(lang)}
                    className={`mr-2 mb-2 px-4 py-2 rounded-full ${
                      selectedLanguage === lang ? 'bg-[#29374C]' : 'bg-[#1E2A3A]'
                    }`}
                  >
                    <Text className="text-[#F8F9FB] text-base capitalize">
                      {lang === 'cpp' ? 'C++' : lang.charAt(0).toUpperCase() + lang.slice(1)}
                    </Text>
                  </TouchableOpacity>
                ))
              )}
            </View>
            
            {/* Solution Code Button/Content */}
            {!showFullSolution ? (
              <TouchableOpacity 
                className="bg-[#29374C] p-4 rounded-lg mb-4"
                onPress={() => setShowFullSolution(true)}
              >
                <Text className="text-[#F8F9FB] text-center font-medium">
                  Reveal Full Solution
                </Text>
              </TouchableOpacity>
            ) : (
              <View className="bg-[#1E2A3A] rounded-lg p-4 mb-6">
                {currentCard.solutions && currentCard.solutions[selectedLanguage] ? (
                  <SyntaxHighlighter
                    language={selectedLanguage}
                    style={atomOneDark}
                    customStyle={{ backgroundColor: '#1E2A3A' }}
                  >
                    {currentCard.solutions[selectedLanguage]}
                  </SyntaxHighlighter>
                ) : (
                  <Text className="text-[#8A9DC0]">
                    Solution not available for this language
                  </Text>
                )}
              </View>
            )}
          </>
        )}
        
        {/* Rating Buttons - only show after solution approach is revealed */}
        {showSolutionApproach && (
          <View className="mb-8">
            <Text className="text-[#F8F9FB] text-lg font-medium mb-4 text-center">
              How well did you know this solution?
            </Text>
            <View className="flex-row justify-between">
              <TouchableOpacity 
                className="bg-[#E74C3C] p-3 rounded-lg flex-1 mx-1"
                onPress={() => handleRating(1)}
              >
                <Text className="text-white text-center">Very Hard</Text>
              </TouchableOpacity>
              <TouchableOpacity 
                className="bg-[#F39C12] p-3 rounded-lg flex-1 mx-1"
                onPress={() => handleRating(2)}
              >
                <Text className="text-white text-center">Hard</Text>
              </TouchableOpacity>
              <TouchableOpacity 
                className="bg-[#3498DB] p-3 rounded-lg flex-1 mx-1"
                onPress={() => handleRating(3)}
              >
                <Text className="text-white text-center">Good</Text>
              </TouchableOpacity>
              <TouchableOpacity 
                className="bg-[#4CD137] p-3 rounded-lg flex-1 mx-1"
                onPress={() => handleRating(4)}
              >
                <Text className="text-white text-center">Easy</Text>
              </TouchableOpacity>
            </View>
          </View>
        )}
      </ScrollView>
    </View>
  );
}
```

### 5. Update Tab Navigation

```tsx
// app/_layout.tsx (update to include flashcards and decks tabs)
<Tabs
  screenOptions={{
    tabBarActiveTintColor: "#6366F1",
    tabBarInactiveTintColor: "#8A9DC0",
    tabBarStyle: {
      backgroundColor: "#131C24",
      borderTopColor: "#32415D",
    },
  }}
>
  <Tab.Screen
    name="index"
    options={{
      title: "Home",
      tabBarIcon: ({ color }) => <Ionicons name="home-outline" size={24} color={color} />,
    }}
  />
  <Tab.Screen
    name="dashboard"
    options={{
      title: "Dashboard",
      tabBarIcon: ({ color }) => <Ionicons name="stats-chart-outline" size={24} color={color} />,
    }}
  />
  <Tab.Screen
    name="reviews"
    options={{
      title: "Reviews",
      tabBarIcon: ({ color }) => <Ionicons name="calendar-outline" size={24} color={color} />,
    }}
  />
  <Tab.Screen
    name="flashcards"
    options={{
      title: "Flashcards",
      tabBarIcon: ({ color }) => <Ionicons name="layers-outline" size={24} color={color} />,
    }}
  />
  <Tab.Screen
    name="decks"
    options={{
      title: "Decks",
      tabBarIcon: ({ color }) => <Ionicons name="albums-outline" size={24} color={color} />,
    }}
  />
  {/* Other tabs */}
</Tabs>
```

### 6. Update Dashboard to Show Deck-Specific Stats

```tsx
// Update dashboard.tsx
// In the dashboard component
const {
  data: flashcardStats = { 
    total: 0, 
    due: 0, 
    completed_today: 0,
    decks: []
  },
  isLoading: flashcardsLoading
} = useFlashcardStats();

// Then in the JSX
<TouchableOpacity onPress={() => router.push('/flashcards')}>
  <InfoCard
    icon="layers-outline"
    title="Flashcard Reviews"
    subtitle={`${flashcardStats.due} cards due`}
  />
</TouchableOpacity>

{/* Add deck stats section if there are decks with due cards */}
{flashcardStats.decks.length > 0 && (
  <View className="mt-4">
    <Text 
      className="text-[#F8F9FB] text-[22px] font-bold leading-tight px-4 pb-3"
    >
      Deck Status
    </Text>
    {flashcardStats.decks
      .filter(deck => deck.due > 0)
      .slice(0, 3)
      .map(deck => (
        <TouchableOpacity 
          key={deck.deck_id}
          onPress={() => router.push(`/flashcards?deck=${deck.deck_id}`)}
          className="flex-row items-center justify-between px-4 py-3"
        >
          <Text className="text-[#F8F9FB]">{deck.deck_name}</Text>
          <Text className="text-[#6366F1] font-medium">
            {deck.due} due
          </Text>
        </TouchableOpacity>
      ))
    }
    
    {flashcardStats.decks.length > 3 && (
      <TouchableOpacity 
        className="items-center mt-2"
        onPress={() => router.push('/decks')}
      >
        <Text className="text-[#6366F1]">View all decks</Text>
      </TouchableOpacity>
    )}
  </View>
)}
```

## Data Flow and Implementation Strategy

### Database Polling and Frontend State Management

To avoid excessive polling after ratings are submitted:

1. **Optimistic UI Updates**: After a user submits a rating, optimistically update the UI to show the next card before the API request completes.

2. **Batch Fetching**: When loading the flashcards screen, fetch all due flashcards at once rather than one at a time.

3. **Efficient Invalidation**: Only invalidate the required queries after mutations, using TanStack Query's fine-grained invalidation.

4. **Client-Side Filtering**: Sort and filter flashcards on the client-side after initial fetch.

### Implementation Sequence

1. **Database Setup**
   - Add solution_approach column to problems table
   - Create decks and deck_problems tables
   - Create flashcard_reviews and flashcard_review_logs tables

2. **Backend API Development**
   - Create models and database functions for decks and flashcards
   - Implement API handlers for deck and flashcard operations
   - Integrate with existing FSRS implementation

3. **Seed Default Decks**
   - Create script to seed Blind75, Neetcode150 and other popular problem sets

4. **OpenRouter Integration**
   - Develop and run the Python script to generate solution approaches for existing problems
   - Validate and refine the approaches as needed

5. **Frontend Implementation**
   - Create the deck browsing and management screens
   - Implement the flashcard services and screens
   - Implement the review flow with reveal stages
   - Add tab navigation and dashboard integration

6. **Testing and Refinement**
   - Test the full flow from adding decks to reviewing and rating
   - Optimize performance and address any issues

## Final Notes

- The deck system allows users to organize problems into meaningful groups like Blind75 or Neetcode150
- Users can browse public decks or create their own custom decks
- The flashcard system will focus on helping users memorize solution approaches rather than recall problem details
- The FSRS algorithm will optimally space reviews to maximize long-term retention
- The OpenRouter/DeepSeek integration will provide AI-generated solution approaches that are concise and language-agnostic

## Data Flow Diagram

```mermaid
graph TD
    A[User browses problem decks] -->|Selects deck| B[View deck problems]
    B -->|Add to flashcards| C[Add all problems to review queue]
    D[User visits flashcard page] -->|Shows due cards| E[Filter by deck if desired]
    E -->|User reviews card| F[Show problem + description]
    F -->|Reveal| G[Show solution approach]
    G -->|Reveal| H[Show full solution]
    H -->|User rates knowledge| I[Submit rating]
    I -->|FSRS Algorithm| J[Calculate next review date]
    J -->|Update review queue| C
    K[AI analyzes problems] -->|Generates approach| L[Stores in database]
    L -->|Used in flashcards| G