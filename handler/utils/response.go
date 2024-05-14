package handler_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"msg": err.Error(),
	})
}

func SendInternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"msg": "Internal server error",
	})
}

func SendMessage(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func SendCreatedMessage(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusCreated, gin.H{
		"msg": msg,
	})
}

func SendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
