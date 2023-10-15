package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSpending(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete spending",
	})
}
