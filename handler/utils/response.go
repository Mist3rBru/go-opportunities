package handler_utils

import (
	"gopportunities/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendBadRequestError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"msg": err.Error(),
	})
}

func SendNotFoundError(ctx *gin.Context, item string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"msg": domain.NotFoundError(item).Error(),
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
