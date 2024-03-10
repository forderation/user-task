package task

import "github.com/forderation/user-task/internal/repository/tasks"

type Usecase struct {
	TaskRepository tasks.TaskRepositoryInterface
}

type NewUsecaseOptions struct {
	TaskRepository tasks.TaskRepositoryInterface
}

func NewUsecase(opts NewUsecaseOptions) *Usecase {
	return &Usecase{
		TaskRepository: opts.TaskRepository,
	}
}
