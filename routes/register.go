package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacquesbeets/go-rest-api/models"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event"})
}

func unregisterFromEvent(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventID

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	err = event.UnregisterFromEvent(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Unregistered from event"})
}
