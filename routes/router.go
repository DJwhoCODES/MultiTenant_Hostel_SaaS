package routes

import (
	"github.com/djwhocodes/hostel_saas/internal/modules/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	authRepo := auth.NewRepository()
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	auth.RegisterRoutes(api, authHandler)
}
