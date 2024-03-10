package middleware

import "github.com/forderation/user-task/internal/usecase/user"

type Middleware struct {
	UserUsecase user.UserUsecaseInterface
}

type NewMiddlewareOptions struct {
	UserUsecase user.UserUsecaseInterface
}

func NewMiddleware(opts NewMiddlewareOptions) *Middleware {
	return &Middleware{
		UserUsecase: opts.UserUsecase,
	}
}

type ErrorBodyResponse struct {
	Error string `json:"error"`
}
