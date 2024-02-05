package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacquesbeets/go-rest-api/db"
	"github.com/jacquesbeets/go-rest-api/models"
)

func main() {
	// Code here
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":9090")
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = len(models.GetEvents()) + 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
