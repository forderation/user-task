package task

import (
	"context"

	"github.com/forderation/user-task/internal/repository/tasks"
)

func (uc *Usecase) DeleteTask(ctx context.Context, input DeleteTaskInput) (err error) {
	err = uc.TaskRepository.DeleteTask(ctx, tasks.DeleteTaskInput{
		ID: input.ID,
	})
	return
}

type DeleteTaskInput struct {
	ID int32
}
