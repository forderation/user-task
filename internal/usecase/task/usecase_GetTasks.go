package task

import (
	"context"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/repository/tasks"
	"github.com/forderation/user-task/util/pagination"
)

func (uc *Usecase) GetTasks(ctx context.Context, input GetTasksInput) (output GetTasksOutput, err error) {
	tasksOutput, err := uc.TaskRepository.GetTasks(ctx, tasks.GetTasksInput{
		Pagination: input.Pagination,
		UserID:     input.UserID,
	})
	if err != nil {
		return
	}
	output = GetTasksOutput{
		Pagination: tasksOutput.Pagination,
		Items:      make([]GetTasksOutputItem, 0),
	}
	for _, task := range tasksOutput.Items {
		output.Items = append(output.Items, GetTasksOutputItem{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return
}

type GetTasksInput struct {
	UserID     null.Int32
	Pagination pagination.PaginationInput
}

type GetTasksOutput struct {
	Pagination pagination.PaginationOutput
	Items      []GetTasksOutputItem
}

type GetTasksOutputItem struct {
	ID          int32
	UserID      int32
	Title       string
	Description null.String
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
