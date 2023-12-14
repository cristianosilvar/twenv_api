package spending

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowSpending(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Show spending",
	})
}
