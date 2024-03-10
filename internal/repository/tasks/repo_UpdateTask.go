package tasks

import (
	"context"
	"time"

	"github.com/forderation/null"
)

func (repo *Repository) UpdateTask(ctx context.Context, input UpdateTaskInput) (err error) {
	task := repo.Query.Task
	query := task.WithContext(ctx).Where(
		task.ID.Eq(input.ID),
	)
	timeNow := time.Now().UTC()
	updateParam := map[string]interface{}{}
	if input.Payload.Title.Valid {
		updateParam[task.Title.ColumnName().String()] = input.Payload.Title.String
	}
	if input.Payload.Description.Valid {
		updateParam[task.Description.ColumnName().String()] = input.Payload.Description.String
	}
	if input.Payload.Status.Valid {
		updateParam[task.Status.ColumnName().String()] = input.Payload.Status.String
	}
	updateParam[task.UpdatedAt.ColumnName().String()] = timeNow
	_, err = query.Updates(updateParam)
	if err != nil {
		return
	}
	return
}

type UpdateTaskInput struct {
	ID      int32
	Payload UpdateTaskInputPayload
}

type UpdateTaskInputPayload struct {
	Title       null.String
	Description null.String
	Status      null.String
}
