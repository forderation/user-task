package tasks

import (
	"context"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/db"
	"github.com/forderation/user-task/util/pagination"
)

func (repo *Repository) GetTasks(ctx context.Context, input GetTasksInput) (output GetTasksOutput, err error) {
	task := repo.Query.Task
	query := task.WithContext(ctx).Select(
		task.ID,
		task.UserID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	)
	if input.UserID.Valid {
		query = query.Where(task.UserID.Eq(input.UserID.Int32))
	}
	totalData, err := query.Count()
	if err != nil {
		return
	}
	db.ApplyPaginationQuery(query.DO, input.Pagination)
	models, err := query.Find()
	if err != nil {
		return
	}
	output = GetTasksOutput{
		Items:      make([]GetTasksOutputItem, 0),
		Pagination: pagination.CreatePaginationOutput(input.Pagination, totalData),
	}
	for _, model := range models {
		output.Items = append(output.Items, GetTasksOutputItem{
			ID:          model.ID,
			UserID:      model.UserID,
			Title:       model.Title,
			Description: null.StringFromPtr(model.Description),
			Status:      model.Status,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		})
	}
	return
}

type GetTasksInput struct {
	Pagination pagination.PaginationInput
	UserID     null.Int32
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
