package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			handleInternalError(c, err)
			c.Abort()
		}
	}
}

func handleInternalError(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	// You can log the error for debugging or analysis here
	log.Printf("Internal Error: %v", err)
}
