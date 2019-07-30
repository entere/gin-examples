package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestID ...
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-ID")

		// Create request id with UUID4
		if requestID == "" {
			u4 := uuid.NewV4()
			requestID = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-ID", requestID)

		// Set X-Request-ID header
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
