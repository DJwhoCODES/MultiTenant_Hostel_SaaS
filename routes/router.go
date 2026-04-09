package routes

import (
	"github.com/djwhocodes/hostel_saas/internal/middleware"
	"github.com/djwhocodes/hostel_saas/internal/modules/auth"
	"github.com/djwhocodes/hostel_saas/internal/modules/student"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	authRepo := auth.NewRepository()
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	auth.RegisterRoutes(api, authHandler)

	protected := api.Group("/")
	protected.Use(
		middleware.AuthMiddleware(),
		middleware.TenantMiddleware(),
	)

	protected.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"userId":   middleware.GetUserID(c),
			"tenantId": middleware.GetTenantID(c),
			"role":     middleware.GetRole(c),
		})
	})

	studentRepo := student.NewRepository()
	studentService := student.NewService(studentRepo)
	studentHandler := student.NewHandler(studentService)

	student.RegisterRoutes(protected, studentHandler)
}
