package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacquesbeets/go-rest-api/db"
	"github.com/jacquesbeets/go-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":9090")
}
