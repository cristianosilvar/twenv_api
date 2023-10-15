package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/spending", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Show spending",
			})
		})
		v1.POST("/spending", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Create spending",
			})
		})
		v1.PUT("/spending", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Put spending",
			})
		})
		v1.DELETE("/spending", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Delete spending",
			})
		})
		v1.GET("/spendings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "All spending",
			})
		})
	}
}
