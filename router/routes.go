package router

import (
	"twenv/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/spending", handlers.ShowSpending)
		v1.POST("/spending", handlers.CreateSpending)
		v1.PUT("/spending", handlers.UpdateSpending)
		v1.DELETE("/spending", handlers.DeleteSpending)
		v1.GET("/spendings", handlers.ListSpending)
	}
}
