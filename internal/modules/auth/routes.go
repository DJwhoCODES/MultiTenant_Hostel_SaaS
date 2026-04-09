package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	auth := r.Group("/auth")

	auth.POST("/signup", h.Signup)
	auth.POST("/login", h.Login)
}
