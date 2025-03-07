package main

import (
	"context"
	"flag"
	"fmt"
	"go-leetcode/backend/api/routes"
	"go-leetcode/backend/internal/database"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	var port string 
	flag.StringVar(&port, "port", "8080", "Server port")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning, failed to load .env file: %v\n", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		os.Exit(1)
	}

	defer logger.Sync()

	log := logger.Sugar()
	log.Info("Starting SPACECODE")

	dbConfig := &database.Config{
		Host:	getEnv("DB_HOST", "localhost"),
		Port: 	getEnv("DB_PORT", "5432"),
		User: 	getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		DBName: getEnv("DB_NAME", "leetcode_practice"),
	}

	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// test db connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database, %v", err)
	}

	log.Info("Connected to database")
	router := routes.SetupRoutes(db, logger)

	srv := &http.Server{
		Addr: 	":" + port,
		Handler:	router,
	}

	go func() {
		log.Infof("Server starting on :%s", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Info("Server exited gracefully")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}