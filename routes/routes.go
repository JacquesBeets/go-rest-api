package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jacquesbeets/go-rest-api/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// Users
	server.POST("/signup", createUser)
	server.POST("/login", loginUser)

	// Events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", unregisterFromEvent)

}
