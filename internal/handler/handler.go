package handler

import "github.com/forderation/user-task/internal/usecase/user"

type Handler struct {
	UserUsecase user.UserUsecaseInterface
}

type NewHandlerOptions struct {
	UserUsecase user.UserUsecaseInterface
}

func NewHandler(opts NewHandlerOptions) *Handler {
	return &Handler{
		UserUsecase: opts.UserUsecase,
	}
}

type ErrorBodyResponse struct {
	Error string `json:"error"`
}
