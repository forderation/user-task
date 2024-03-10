package handler

import (
	"net/http"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/usecase/task"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TaskPost(c *gin.Context) {
	ctx := c.Request.Context()
	var bodyRequest TaskPostBodyRequest
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid body request",
		})
		return
	}
	if bodyRequest.Title == "" {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "title is required",
		})
		return
	}
	if bodyRequest.Status == "" {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "status is required",
		})
		return
	}
	userIDAuth := int32(c.GetInt(middleware.KEY_CONTEXT_USER_ID))
	taskOutput, err := h.TaskUsecase.CreateTask(ctx, task.CreateTaskInput{
		Title:       bodyRequest.Title,
		Description: null.StringFromPtr(bodyRequest.Description),
		Status:      bodyRequest.Status,
		UserID:      userIDAuth,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, TaskPostBodyResponse{
		ID: taskOutput.ID,
	})
}

type TaskPostBodyRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
}

type TaskPostBodyResponse struct {
	ID int32 `json:"id"`
}
