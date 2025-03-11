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

	userStore := models.NewUserStore(db)
	reviewStore := models.NewReviewScheduleStore(db)
	problemStore := models.NewProblemStore(db)
	submissionStore := models.NewSubmissionStore(db)

	userHandler := handlers.NewUserHandler(userStore)
	reviewHandler := handlers.NewReviewHandler(reviewStore)
	problemHandler := handlers.NewProblemHandler(problemStore)
	submissionHandler := handlers.NewSubmissionHandler(submissionStore)

	router.Get("/health", handlers.HealthCheck)

	router.Route("/api/users", func(router chi.Router) {
		router.Post("/", userHandler.Register)
	})

	router.Route("/api/reviews", func(router chi.Router) {
		router.Get("/", reviewHandler.GetReviews)
		router.Put("/", reviewHandler.UpdateReviewSchedule)
		router.Post("/", reviewHandler.CreateReview)
	})

	router.Route("/api/problems", func(router chi.Router) {
		router.Get("/by-id", problemHandler.GetProblemByID)
		router.Get("/by-frontend-id", problemHandler.GetProblemByFrontendID)
		router.Get("/by-slug", problemHandler.GetProblemBySlug)
		router.Get("/list", problemHandler.GetProblemList)
	})

	router.Get("/api/submissions", submissionHandler.GetSubmissions) 
	router.Post("/api/submissions", submissionHandler.CreateSubmission)

	return router
}
