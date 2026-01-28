package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/cache"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/database"
	"github.com/shubham21155102/golang-full-ci-cd-template/internal/models"
)

type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Redis    string `json:"redis"`
	Time     string `json:"time"`
}

func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
	}

	// Check database connection
	db := database.GetDB()
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			response.Database = "disconnected"
			response.Status = "degraded"
		} else {
			response.Database = "connected"
		}
	} else {
		response.Database = "not initialized"
		response.Status = "degraded"
	}

	// Check redis connection
	redisClient := cache.GetClient()
	if redisClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := redisClient.Ping(ctx).Err(); err != nil {
			response.Redis = "disconnected"
			response.Status = "degraded"
		} else {
			response.Redis = "connected"
		}
	} else {
		response.Redis = "not initialized"
		response.Status = "degraded"
	}

	if response.Status == "ok" {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusServiceUnavailable, response)
	}
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: true,
	}

	if err := database.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Cache the user in Redis
	ctx := context.Background()
	userJSON, _ := json.Marshal(user)
	cacheKey := fmt.Sprintf("user:%d", user.ID)
	cache.GetClient().Set(ctx, cacheKey, userJSON, 10*time.Minute)

	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	cacheKey := fmt.Sprintf("user:%s", id)

	// Try to get from cache first
	ctx := context.Background()
	cachedUser, err := cache.GetClient().Get(ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		if json.Unmarshal([]byte(cachedUser), &user) == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"source": "cache",
			})
			return
		}
	}

	// Get from database
	var user models.User
	if err := database.GetDB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Cache the result
	userJSON, _ := json.Marshal(user)
	cache.GetClient().Set(ctx, cacheKey, userJSON, 10*time.Minute)

	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"source": "database",
	})
}

func ListUsers(c *gin.Context) {
	var users []models.User
	if err := database.GetDB().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
	})
}
