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

func (h *Handler) UserRegister(c *gin.Context) {
	var bodyRequest UserRegisterBodyRequest
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	if bodyRequest.Name == "" {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "name is required",
		})
		return
	}
	ctx := c.Request.Context()
	emailExist := true
	_, err = h.UserUsecase.GetUser(ctx, user.GetUserInput{
		Email: null.StringFrom(bodyRequest.Email),
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
	if !encryption.ValidatePasswordRequirement(bodyRequest.Password) {
		c.JSON(http.StatusBadRequest, ErrorBodyResponse{
			Error: "password not valid, min 8 character",
		})
		return
	}
	passwordHash, err := encryption.HashPassword(bodyRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	email := bodyRequest.Email
	userOutput, err := h.UserUsecase.CreateUser(ctx, user.CreateUserInput{
		Name:     bodyRequest.Name,
		Email:    email,
		Password: passwordHash,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	tokenJwt, err := auth_jwt.GetJWT().GetToken(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterBodyResponse{
		UserID:       userOutput.ID,
		AccessToken:  tokenJwt.AccessToken,
		RefreshToken: tokenJwt.RefreshToken,
	})
}

type UserRegisterBodyRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterBodyResponse struct {
	UserID       int32  `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
