package task

import (
	"context"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/repository/tasks"
)

func (uc *Usecase) CreateTask(ctx context.Context, input CreateTaskInput) (output CreateTaskOutput, err error) {
	taskOutput, err := uc.TaskRepository.CreateTask(ctx, tasks.CreateTaskInput{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	})
	if err != nil {
		return
	}
	output.ID = taskOutput.ID
	return
}

type CreateTaskInput struct {
	UserID      int32
	Title       string
	Description null.String
	Status      string
}

type CreateTaskOutput struct {
	ID int32
}
