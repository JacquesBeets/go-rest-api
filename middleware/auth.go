package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacquesbeets/go-rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
