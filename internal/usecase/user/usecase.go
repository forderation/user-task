package user

import (
	"github.com/forderation/user-task/internal/repository/tasks"
	"github.com/forderation/user-task/internal/repository/users"
)

type Usecase struct {
	TaskRepository tasks.TaskRepositoryInterface
	UserRepository users.UserRepositoryInterface
}

type NewUsecaseOptions struct {
	TaskRepository tasks.TaskRepositoryInterface
	UserRepository users.UserRepositoryInterface
}

func NewUsecase(opts NewUsecaseOptions) *Usecase {
	return &Usecase{
		TaskRepository: opts.TaskRepository,
		UserRepository: opts.UserRepository,
	}
}
