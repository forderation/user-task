package handler

import (
	"errors"
	"net/http"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/forderation/user-task/util/auth_jwt"
	"github.com/forderation/user-task/util/encryption"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserLogin(c *gin.Context) {
	var bodyRequest UserLoginBodyRequest
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "invalid body request",
		})
		return
	}
	ctx := c.Request.Context()
	userOutput, err := h.UserUsecase.GetUser(ctx, user.GetUserInput{
		Email: null.StringFrom(bodyRequest.Email),
	})
	if errors.Is(err, user.ErrUserNotFound) {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "email or password not valid",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	if !encryption.CheckPasswordHash(bodyRequest.Password, userOutput.Password) {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "email or password not valid",
		})
		return
	}
	tokenJwt, err := auth_jwt.GetJWT().GetToken(userOutput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginBodyResponse{
		UserID:       userOutput.ID,
		AccessToken:  tokenJwt.AccessToken,
		RefreshToken: tokenJwt.RefreshToken,
	})
}

type UserLoginBodyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginBodyResponse struct {
	UserID       int32  `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
