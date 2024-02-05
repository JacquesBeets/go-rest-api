package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Code here
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":9090")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
