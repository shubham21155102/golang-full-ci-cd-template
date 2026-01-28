package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/cache"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/config"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/database"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/handlers"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/models"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	if err := database.Connect(&cfg.Database); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := database.GetDB().AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Connect to Redis
	if err := cache.Connect(&cfg.Redis); err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	defer cache.Close()

	// Setup router
	r := gin.Default()

	// Routes
	r.GET("/health", handlers.HealthCheck)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin Go Application",
			"version": "1.0.0",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", handlers.CreateUser)
		v1.GET("/users/:id", handlers.GetUser)
		v1.GET("/users", handlers.ListUsers)
	}

	// Graceful shutdown
	go func() {
		if err := r.Run(":" + cfg.Server.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on port %s in %s mode", cfg.Server.Port, cfg.Server.Env)

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
