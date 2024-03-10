package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/usecase/task"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TaskPut(c *gin.Context) {
	var bodyRequest TaskPutBodyRequest
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid body request",
		})
		return
	}
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid task id",
		})
		return
	}
	userIDAuth := int32(c.GetInt(middleware.KEY_CONTEXT_USER_ID))
	taskOutput, err := h.TaskUsecase.GetTask(c.Request.Context(), task.GetTaskInput{
		ID:     null.Int32From(int32(taskID)),
		UserID: null.Int32From(userIDAuth),
	})
	if err != nil {
		if errors.Is(err, task.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, ErrorBodyResponse{
				Error: "task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	title := null.StringFromPtr(bodyRequest.Title)
	if title.Valid && title.String == "" {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "title is required",
		})
		return
	}
	description := null.StringFromPtr(bodyRequest.Description)
	status := null.StringFromPtr(bodyRequest.Status)
	if status.Valid && status.String == "" {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "status is required",
		})
		return
	}
	err = h.TaskUsecase.UpdateTask(c.Request.Context(), task.UpdateTaskInput{
		ID:          taskOutput.ID,
		Title:       title,
		Description: description,
		Status:      status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, TaskPutResponseBodyResponse{
		ID: taskOutput.ID,
	})
}

type TaskPutBodyRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type TaskPutResponseBodyResponse struct {
	ID int32 `json:"id"`
}
