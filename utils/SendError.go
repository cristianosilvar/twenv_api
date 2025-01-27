package utils

import "github.com/gin-gonic/gin"

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"success":   false,
		"message":   msg,
		"errorCode": code,
	})
}
