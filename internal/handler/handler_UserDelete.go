package handler

import (
	"net/http"
	"strconv"

	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserDelete(c *gin.Context) {
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
	err = h.UserUsecase.DeleteUser(ctx, user.DeleteUserInput{
		ID: userIDAuth,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
	}
	c.JSON(http.StatusOK, nil)
}
