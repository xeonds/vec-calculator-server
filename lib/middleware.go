package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT中间件
func JWTMiddleware(authToken func(*gin.Context, UserClaim) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		parsed, err := ParseToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if parsed.Valid() != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if authToken != nil && authToken(c, *parsed) != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
