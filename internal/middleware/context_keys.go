package middleware

import "github.com/gin-gonic/gin"

type ContextKey string

const (
	UserIDKey   ContextKey = "userId"
	TenantIDKey ContextKey = "tenantId"
	RoleKey     ContextKey = "role"
)

func GetUserID(c *gin.Context) string {
	val, _ := c.Get(string(UserIDKey))
	return val.(string)
}

func GetTenantID(c *gin.Context) string {
	val, _ := c.Get(string(TenantIDKey))
	return val.(string)
}

func GetRole(c *gin.Context) string {
	val, _ := c.Get(string(RoleKey))
	return val.(string)
}
