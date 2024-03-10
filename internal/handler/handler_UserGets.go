package handler

import (
	"net/http"
	"strconv"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/forderation/user-task/util/pagination"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserGets(c *gin.Context) {
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
	usersOutput, err := h.UserUsecase.GetUsers(ctx, user.GetUsersInput{
		Pagination: paginationInput,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	response := UserGetsBodyResponse{
		Items:      make([]UserGetsBodyResponseItem, 0),
		Pagination: bindPaginationOutput(usersOutput.Pagination),
	}
	for _, v := range usersOutput.Items {
		response.Items = append(response.Items, UserGetsBodyResponseItem{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}
	c.JSON(http.StatusOK, response)
}

type UserGetsBodyResponse struct {
	Items      []UserGetsBodyResponseItem `json:"items"`
	Pagination PaginationResponse         `json:"pagination"`
}

type UserGetsBodyResponseItem struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
