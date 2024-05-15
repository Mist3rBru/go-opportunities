package router

import (
	"gopportunities/internal/handler"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	g := router.Group("/api")
	{
		g.POST("/opening", handler.CreateOpeningHandler)
		g.PUT("/openings", handler.UpdateOpeningHandler)
		g.GET("/openings", handler.ListOpeningsHandler)
		g.GET("/openings/:id", handler.FindOpeningByIdHandler)
		g.DELETE("/openings/:id", handler.DeleteOpeningHandler)
	}
}
