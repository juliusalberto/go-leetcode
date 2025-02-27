package routes

import (
	"go-leetcode/backend/api/handlers"
	"go-leetcode/backend/api/middleware"
	"go-leetcode/backend/models"
	"database/sql"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func SetupRoutes(db *sql.DB, logger *zap.Logger,) chi.Router {
	router := chi.NewRouter()

	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.RealIP)
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.LoggingMiddleware(logger))

	userStore := models.NewUserStore(db)
	reviewStore := models.NewReviewScheduleStore(db)

	userHandler := handlers.NewUserHandler(userStore)
	reviewHandler := handlers.NewReviewHandler(reviewStore)

	router.Get("/health", handlers.HealthCheck)

	router.Route("/api/users", func(router chi.Router) {
		router.Post("/register", userHandler.Register)
	})

	router.Route("/api/reviews", func(router chi.Router) {
		router.Get("/upcoming", reviewHandler.GetUpcomingReviews)
		router.Put("/update", reviewHandler.UpdateReviewSchedule)
		router.Post("/create", reviewHandler.CreateReview)
	})

	return router
}