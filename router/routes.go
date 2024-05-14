package router

import (
	"gopportunities/handler"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	g := router.Group("/api")
	{
		g.POST("/opening", handler.CreateOpeningHandler)
		g.GET("/openings", handler.ListOpeningsHandler)
	}
}
