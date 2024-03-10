package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/forderation/user-task/util/auth_jwt"
	"github.com/gin-gonic/gin"
)

func (h *Middleware) CheckUserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		authHeader := c.Request.Header["Authorization"]
		var header string
		if len(authHeader) > 0 {
			header = authHeader[0]
		}
		if header == "" {
			c.JSON(http.StatusUnauthorized, ErrorBodyResponse{
				Error: "required authorization",
			})
			c.Abort()
			return
		}
		mapClaims, err := auth_jwt.GetJWT().GetClaims(header)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
				Error: err.Error(),
			})
			c.Abort()
			return
		}
		if !mapClaims.AccessToken {
			c.JSON(http.StatusUnauthorized, ErrorBodyResponse{
				Error: "token is not access token",
			})
			c.Abort()
			return
		}
		timeNow := time.Now().UTC().Unix()
		if timeNow > mapClaims.ExpUnix {
			c.JSON(http.StatusUnauthorized, ErrorBodyResponse{
				Error: "access token expired",
			})
			c.Abort()
			return
		}
		userOutput, err := h.UserUsecase.GetUser(ctx, user.GetUserInput{
			Email: null.StringFrom(mapClaims.Sub),
		})
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(http.StatusUnauthorized, ErrorBodyResponse{
				Error: "user not found",
			})
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorBodyResponse{
				Error: err.Error(),
			})
			c.Abort()
			return
		}
		c.Set(KEY_CONTEXT_USER_ID, userOutput.ID)
		c.Next()
	}
}
