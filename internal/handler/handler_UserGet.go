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

func (h *Handler) UserGet(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid user id",
		})
		return
	}
	userIDAuth := int32(c.GetInt(middleware.KEY_CONTEXT_USER_ID))
	if userIDAuth != int32(userID) {
		c.JSON(http.StatusUnauthorized, ErrorBodyResponse{
			Error: "unauthorized to delete user",
		})
		return
	}
	ctx := c.Request.Context()
	userOutput, err := h.UserUsecase.GetUser(ctx, user.GetUserInput{
		ID: null.Int32From(userIDAuth),
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
	response := UserGetBodyResponse{
		ID:    userOutput.ID,
		Name:  userOutput.Name,
		Email: userOutput.Email,
		Tasks: make([]UserGetBodyResponseTask, 0),
	}
	tasksOutput, err := h.TaskUsecase.GetTasks(ctx, task.GetTasksInput{
		UserID: null.Int32From(userOutput.ID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	for _, taskOutput := range tasksOutput.Items {
		response.Tasks = append(response.Tasks, UserGetBodyResponseTask{
			ID:          taskOutput.ID,
			Title:       taskOutput.Title,
			Description: taskOutput.Description.Ptr(),
			Status:      taskOutput.Status,
			CreatedAt:   taskOutput.CreatedAt,
			UpdatedAt:   taskOutput.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, response)
}

type UserGetBodyResponse struct {
	ID    int32                     `json:"id"`
	Name  string                    `json:"name"`
	Email string                    `json:"email"`
	Tasks []UserGetBodyResponseTask `json:"tasks"`
}

type UserGetBodyResponseTask struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
