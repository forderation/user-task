package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/usecase/task"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TaskGet(c *gin.Context) {
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
	userOutput, err := h.UserUsecase.GetUser(ctx, user.GetUserInput{
		ID: null.Int32From(taskOutput.UserID),
	})
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, ErrorBodyResponse{
				Error: "user not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	response := TaskGetResponseBodyResponse{
		ID:          taskOutput.ID,
		Title:       taskOutput.Title,
		Description: taskOutput.Description.Ptr(),
		Status:      taskOutput.Status,
		CreatedAt:   taskOutput.CreatedAt,
		UpdatedAt:   taskOutput.UpdatedAt,
		User: TaskGetResponseBodyResponseUser{
			ID:    userOutput.ID,
			Name:  userOutput.Name,
			Email: userOutput.Email,
		},
	}
	c.JSON(http.StatusOK, response)
}

type TaskGetResponseBodyResponse struct {
	ID          int32                           `json:"id"`
	Title       string                          `json:"title"`
	Description *string                         `json:"description,omitempty"`
	Status      string                          `json:"status"`
	CreatedAt   time.Time                       `json:"created_at"`
	UpdatedAt   time.Time                       `json:"updated_at"`
	User        TaskGetResponseBodyResponseUser `json:"user"`
}

type TaskGetResponseBodyResponseUser struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
