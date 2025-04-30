package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleAuth(expectedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != expectedRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			return
		}
		c.Next()
	}
}
