package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSpending(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List spending",
	})
}
