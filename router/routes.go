package router

import (
	"gopportunities/handler"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	g := router.Group("/api")
	{
		g.GET("/opening", handler.GetOpeningHandler)
		g.POST("/opening", handler.CreateOpeningHandler)
	}
}
