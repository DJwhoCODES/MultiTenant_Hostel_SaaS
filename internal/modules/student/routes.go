package student

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	students := r.Group("/students")

	students.POST("/", h.Create)
	students.GET("/", h.GetAll)
	students.GET("/:id", h.GetOne)
	students.PUT("/:id", h.Update)
	students.DELETE("/:id", h.Delete)
}
