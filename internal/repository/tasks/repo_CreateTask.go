package tasks

import (
	"context"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/db/model"
)

func (repo *Repository) CreateTask(ctx context.Context, input CreateTaskInput) (output CreateTaskOutput, err error) {
	timeNow := time.Now().UTC()
	model := model.Task{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description.Ptr(),
		Status:      input.Status,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}
	err = repo.Query.Task.WithContext(ctx).Create(&model)
	if err != nil {
		return
	}
	output.ID = model.ID
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
