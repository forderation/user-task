package task

import (
	"context"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/repository/tasks"
)

func (uc *Usecase) UpdateTask(ctx context.Context, input UpdateTaskInput) (err error) {
	err = uc.TaskRepository.UpdateTask(ctx, tasks.UpdateTaskInput{
		ID: input.ID,
		Payload: tasks.UpdateTaskInputPayload{
			Title:       input.Title,
			Description: input.Description,
			Status:      input.Status,
		},
	})
	return
}

type UpdateTaskInput struct {
	ID          int32
	Title       null.String
	Description null.String
	Status      null.String
}
