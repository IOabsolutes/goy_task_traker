package handlers

import (
	"github.com/gin-gonic/gin"
)

// Handler contains all the handler dependencies
type Handler struct {
	// Add your services/repositories here when you create them
	// taskService    services.TaskService
	// userService    services.UserService
}

// NewHandler creates a new handler instance
func NewHandler() *Handler {
	return &Handler{
		// Initialize your services here
	}
}

// InitRoutes initializes all application routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/health", h.healthCheck)

	// API v1 routes
	api := router.Group("/api/v1")
	{
		h.initUserRoutes(api)
		h.initTaskRoutes(api)
	}

	return router
}

// healthCheck is a simple health check endpoint
func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Task Tracker API is running",
	})
}
