package handler

import (
	"gopportunities/domain"
	handler_utils "gopportunities/handler/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	req := domain.Opening{}
	ctx.BindJSON(&req)

	if err := uuid.Validate(req.Id); err != nil {
		handler_utils.SendBadRequestError(ctx, err)
		return
	}

	opening := domain.Opening{}
	if err := db.First(&opening, "id = ?", req.Id).Error; err != nil {
		handler_utils.SendNotFoundError(ctx, "opening")
		return
	}

	if req.Company != "" {
		opening.Company = req.Company
	}
	if req.Role != "" {
		opening.Role = req.Role
	}
	if req.Link != "" {
		opening.Link = req.Link
	}
	if req.Location != "" {
		opening.Location = req.Location
	}
	if req.Salary != "" {
		opening.Salary = req.Salary
	}
	if req.Remote != nil {
		opening.Remote = req.Remote
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("failed to update opening: %v", err)
		handler_utils.SendInternalServerError(ctx)
		return
	}

	handler_utils.SendMessage(ctx, "opening update with success")
}
