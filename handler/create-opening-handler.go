package handler

import (
	"gopportunities/domain"
	handler_utils "gopportunities/handler/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateOpeningHandler(ctx *gin.Context) {
	opening := domain.Opening{}

	ctx.BindJSON(&opening)

	if err := opening.Validate(); err != nil {
		handler_utils.SendBadRequestError(ctx, err)
		return
	}

	opening.Id = uuid.NewString()
	opening.CreatedAt = time.Now()
	opening.UpdatedAt = time.Now()

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		handler_utils.SendInternalServerError(ctx)
		return
	}

	handler_utils.SendCreatedMessage(ctx, "opening created with success")
}
