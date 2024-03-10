package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/usecase/task"
	"github.com/forderation/user-task/util/pagination"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TaskGets(c *gin.Context) {
	ctx := c.Request.Context()
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	page := null.IntFrom(1)
	pageSize := null.IntFrom(10)
	if pageStr != "" {
		pageInt, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorBodyResponse{
				Error: "invalid page",
			})
			return
		}
		page = null.IntFrom(pageInt)
	}
	if pageSizeStr != "" {
		pageSizeInt, err := strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorBodyResponse{
				Error: "invalid page size",
			})
			return
		}
		pageSize = null.IntFrom(pageSizeInt)
	}
	paginationInput := pagination.PaginationInput{
		Page:     page,
		PageSize: pageSize,
	}
	userIDAuth := int32(c.GetInt(middleware.KEY_CONTEXT_USER_ID))
	tasksOutput, err := h.TaskUsecase.GetTasks(ctx, task.GetTasksInput{
		UserID:     null.Int32From(userIDAuth),
		Pagination: paginationInput,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	response := TaskGetsResponseBody{
		Pagination: bindPaginationOutput(tasksOutput.Pagination),
		Items:      make([]TaskGetsResponseBodyItem, 0),
	}
	for _, v := range tasksOutput.Items {
		response.Items = append(response.Items, TaskGetsResponseBodyItem{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description.Ptr(),
			Status:      v.Status,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, tasksOutput)
}

type TaskGetsResponseBody struct {
	Items      []TaskGetsResponseBodyItem `json:"items"`
	Pagination PaginationResponse         `json:"pagination"`
}

type TaskGetsResponseBodyItem struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
