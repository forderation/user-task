package task

import (
	"context"
	"errors"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/repository/tasks"
	"gorm.io/gorm"
)

func (uc *Usecase) GetTask(ctx context.Context, input GetTaskInput) (output GetTaskOutput, err error) {
	taskOutput, err := uc.TaskRepository.GetTask(ctx, tasks.GetTaskInput{
		ID:     input.ID,
		UserID: input.UserID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.Join(err, ErrTaskNotFound)
		}
		return
	}
	output = GetTaskOutput{
		ID:          taskOutput.ID,
		UserID:      taskOutput.UserID,
		Title:       taskOutput.Title,
		Description: taskOutput.Description,
		Status:      taskOutput.Status,
		CreatedAt:   taskOutput.CreatedAt,
		UpdatedAt:   taskOutput.UpdatedAt,
	}
	return
}

type GetTaskInput struct {
	ID     null.Int32
	UserID null.Int32
}

type GetTaskOutput struct {
	ID          int32
	UserID      int32
	Title       string
	Description null.String
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
