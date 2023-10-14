package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sucess",
		})
	})

	r.Run(":4000")
}
