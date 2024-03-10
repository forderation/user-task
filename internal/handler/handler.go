package handler

import (
	"github.com/forderation/user-task/internal/usecase/task"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/forderation/user-task/util/pagination"
)

type Handler struct {
	UserUsecase user.UserUsecaseInterface
	TaskUsecase task.TaskUsecaseInterface
}

type NewHandlerOptions struct {
	UserUsecase user.UserUsecaseInterface
	TaskUsecase task.TaskUsecaseInterface
}

func NewHandler(opts NewHandlerOptions) *Handler {
	return &Handler{
		UserUsecase: opts.UserUsecase,
		TaskUsecase: opts.TaskUsecase,
	}
}

type ErrorBodyResponse struct {
	Error string `json:"error"`
}

type PaginationResponse struct {
	Page      int64 `json:"page"`
	PageCount int64 `json:"page_count"`
	PageSize  int64 `json:"page_size"`
	TotalData int64 `json:"total_data"`
}

func bindPaginationOutput(
	paginationOutput pagination.PaginationOutput,
) (output PaginationResponse) {
	resp := PaginationResponse{
		Page:      paginationOutput.Page,
		PageCount: paginationOutput.PageCount,
		TotalData: paginationOutput.TotalData,
		PageSize:  paginationOutput.PageSize,
	}
	return resp
}
