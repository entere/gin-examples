package util

import (
	"github.com/gin-gonic/gin"
)

// GetReqID ...
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-ID")
	if !ok {
		return ""
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}

// GetUserID ...
func GetUserID(c *gin.Context) string {
	v, ok := c.Get("X-User-ID")
	if !ok {
		return ""
	}
	if userID, ok := v.(string); ok {
		return userID
	}
	return ""
}
