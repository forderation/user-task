package tasks

import (
	"context"
	"time"

	"github.com/forderation/null"
)

func (repo *Repository) GetTask(ctx context.Context, input GetTaskInput) (output GetTaskOutput, err error) {
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
	if input.ID.Valid {
		query = query.Where(task.ID.Eq(input.ID.Int32))
	}
	model, err := query.First()
	if err != nil {
		return
	}
	output = GetTaskOutput{
		ID:          model.ID,
		UserID:      model.UserID,
		Title:       model.Title,
		Description: null.StringFromPtr(model.Description),
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	return
}

type GetTaskInput struct {
	ID null.Int32
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
