package handler

import (
	"gopportunities/domain"
	handler_utils "gopportunities/handler/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindOpeningByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uuid.Validate(id); err != nil {
		handler_utils.SendBadRequestError(ctx, err)
		return
	}

	opening := domain.Opening{}
	if err := db.First(&opening, "id = ?", id).Error; err != nil {
		handler_utils.SendNotFoundError(ctx, "opening")
		return
	}

	handler_utils.SendData(ctx, opening)
}
