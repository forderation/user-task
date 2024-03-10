package handler

import (
	"errors"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserUpdate(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid user id",
		})
		return
	}
	userIDEdited := int32(userID)
	userIDAuth := int32(c.GetInt(middleware.KEY_CONTEXT_USER_ID))
	if userIDAuth != userIDEdited {
		c.JSON(http.StatusUnauthorized, ErrorBodyResponse{
			Error: "unauthorized to edit user",
		})
		return
	}
	var bodyRequest UserUpdateBodyRequest
	err = c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid body request",
		})
		return
	}
	ctx := c.Request.Context()
	name := null.StringFromPtr(bodyRequest.Name)
	email := null.StringFromPtr(bodyRequest.Email)
	if name.Valid && name.String == "" {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "name is required",
		})
		return
	}
	if email.Valid {
		_, err = mail.ParseAddress(email.String)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorBodyResponse{
				Error: "invalid email",
			})
			return
		}
		emailExist := true
		_, err = h.UserUsecase.GetUser(ctx, user.GetUserInput{
			Email:   email,
			NotEqID: null.Int32From(userIDAuth),
		})
		if err != nil {
			if errors.Is(err, user.ErrUserNotFound) {
				emailExist = false
				err = nil
			} else {
				c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
					Error: err.Error(),
				})
				return
			}
		}
		if emailExist {
			c.JSON(http.StatusBadRequest, ErrorBodyResponse{
				Error: "email already registered",
			})
			return
		}
	}
	err = h.UserUsecase.UpdateUser(ctx, user.UpdateUserInput{
		ID:    userIDAuth,
		Name:  name,
		Email: email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	response := UserUpdateBodyResponse{
		ID: userIDAuth,
	}
	if email.Valid {
		response.Message = "because email changed you must log in again, to get new access token"
	}
	c.JSON(http.StatusOK, response)
}

type UserUpdateBodyRequest struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type UserUpdateBodyResponse struct {
	ID      int32  `json:"id"`
	Message string `json:"message"`
}
