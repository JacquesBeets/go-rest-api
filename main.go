package main

import "github.com/gin-gonic/gin"

func main() {
	// Code here
	server := gin.Default()

	server.Run(":9090")
}
