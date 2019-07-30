package middleware

import (
	"micro-gin/handler"
	"micro-gin/pkg/errno"
	"micro-gin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware ...
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		claims, err := jwtauth.ParseRequest(c)
		if err != nil {

			handler.SendResponse(c, errno.ErrTokenInvalid, err)
			c.Abort()
			return
		}
		c.Set("X-User-ID", claims.UserID)

		// Set X-Request-ID header
		c.Writer.Header().Set("X-User-ID", claims.UserID)

		c.Next()

	}
}
