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

func (h *Handler) TaskDelete(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid task id",
		})
		return
	}
	userIDAuth := int32(c.GetInt(middleware.KEY_CONTEXT_USER_ID))
	ctx := c.Request.Context()
	taskOutput, err := h.TaskUsecase.GetTask(ctx, task.GetTaskInput{
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
	err = h.TaskUsecase.DeleteTask(ctx, task.DeleteTaskInput{
		ID: taskOutput.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
