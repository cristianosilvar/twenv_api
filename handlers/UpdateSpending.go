package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSpending(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update spending",
	})
}
