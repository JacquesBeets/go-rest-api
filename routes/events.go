package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacquesbeets/go-rest-api/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to update this event"})
		return
	}

	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.ID = eventID
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
