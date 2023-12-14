package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"sucess":  true,
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}
