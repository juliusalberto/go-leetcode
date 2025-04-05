package routes

import (
	"database/sql"
	"go-leetcode/backend/api/handlers"
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func SetupRoutes(db *sql.DB, logger *zap.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.RealIP)
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.CorsMiddleware)

	userStore := models.NewUserStore(db)
	reviewStore := models.NewReviewScheduleStore(db)
	problemStore := models.NewProblemStore(db)
	submissionStore := models.NewSubmissionStore(db)
	flashcardStore := models.NewFlashcardReviewStore(db) // Initialize flashcardStore first
	deckStore := models.NewDeckStore(db, flashcardStore) // Pass flashcardStore to NewDeckStore

	userHandler := handlers.NewUserHandler(userStore)
	reviewHandler := handlers.NewReviewHandler(reviewStore, submissionStore)
	problemHandler := handlers.NewProblemHandler(problemStore)
	problemStatusHandler := handlers.NewProblemStatusHandler(problemStore, submissionStore)
	submissionHandler := handlers.NewSubmissionHandler(submissionStore)
	solutionStore := models.NewSolutionStore(db)
	solutionHandler := handlers.NewSolutionHandler(solutionStore)
	authStatusHandler := handlers.NewAuthStatusHandler(userStore)
	deckHandler := handlers.NewDeckHandler(deckStore, problemStore, flashcardStore)
	flashcardHandler := handlers.NewFlashcardHandler(flashcardStore, problemStore, deckStore)


	router.Get("/health", handlers.HealthCheck)

	router.Route("/api/users", func(router chi.Router) {
		router.Post("/", userHandler.Register)
		router.Get("/", userHandler.GetUser)
	})

	router.Route("/api/problems", func(router chi.Router) {
		router.Get("/by-id", problemHandler.GetProblemByID)
		router.Get("/by-frontend-id", problemHandler.GetProblemByFrontendID)
		router.Get("/by-slug", problemHandler.GetProblemBySlug)
		router.Get("/list", problemHandler.GetProblemList)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware())

		r.Route("/api/submissions", func(submissionRouter chi.Router) {
			submissionRouter.Get("/api/submissions", submissionHandler.GetSubmissions)
			submissionRouter.Post("/api/submissions", submissionHandler.CreateSubmission)
		})

		r.Get("/api/auth/status", authStatusHandler.GetUserAuthStatus)
		r.Post("/api/users/profile", userHandler.CompleteProfile)

		r.Route("/api/reviews", func(reviewsRouter chi.Router) {
			reviewsRouter.Get("/", reviewHandler.GetReviews)
			reviewsRouter.Put("/", reviewHandler.UpdateReviewSchedule)
			reviewsRouter.Post("/", reviewHandler.CreateReview)
			reviewsRouter.Post("/update-or-create", reviewHandler.UpdateOrCreateReview)
			reviewsRouter.Post("/process-submission", reviewHandler.ProcessSubmission)
		})

		r.Get("/api/problems/with-status", problemStatusHandler.GetProblemsWithStatus)

		r.Route("/api/decks", func(deckRouter chi.Router) {
			deckRouter.Get("/", deckHandler.GetAllDecks)
			deckRouter.Post("/", deckHandler.CreateDeck)
			deckRouter.Put("/{id}", deckHandler.UpdateDeck)
			deckRouter.Delete("/{id}", deckHandler.DeleteDeck)
			deckRouter.Get("/{id}/problems", deckHandler.GetDeckProblems)
			deckRouter.Post("/{id}/problems", deckHandler.AddProblemToDeckAndCreateFlashcard)
			deckRouter.Delete("/{id}/problems/{problem_id}", deckHandler.RemoveProblemFromDeck)
			deckRouter.Post("/{id}/start-practice", deckHandler.StartPracticePublicDeck) // Add route for starting practice
		})

		r.Route("/api/flashcards", func(flashcardRouter chi.Router) {
			flashcardRouter.Get("/reviews", flashcardHandler.GetFlashcardReviews)
			flashcardRouter.Post("/reviews", flashcardHandler.SubmitFlashcardReview)
			flashcardRouter.Post("/decks/{deck_id}", flashcardHandler.AddDeckToFlashcards)
		})
	})

	// LeetCode API proxy endpoint
	router.Post("/api/proxy/leetcode", handlers.LeetCodeProxyHandler)

	router.Route("/api/solutions", func(router chi.Router) {
		router.Get("/", solutionHandler.GetSolutions)
		
		// probably wanna put this on auth
		// router.Post("/", solutionHandler.CreateSolution)
		// router.Put("/", solutionHandler.UpdateSolution)
		// router.Delete("/", solutionHandler.DeleteSolution)
	})



	return router
}
