package handler

import (
	"gopportunities/domain"
	handler_utils "gopportunities/handler/utils"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []domain.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("failed to list openings: %v", err)
		handler_utils.SendInternalServerError(ctx)
		return
	}

	handler_utils.SendData(ctx, openings)
}
