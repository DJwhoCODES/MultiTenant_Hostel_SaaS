package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tenantID, exists := c.Get(string(TenantIDKey))
		if !exists || tenantID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "tenant not found"})
			c.Abort()
			return
		}

		headerTenant := c.GetHeader("x-tenant-id")
		if headerTenant != "" && headerTenant != tenantID {
			c.JSON(http.StatusForbidden, gin.H{"error": "tenant mismatch"})
			c.Abort()
			return
		}

		c.Next()
	}
}
