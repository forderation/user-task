package users

import (
	"context"
	"time"

	"github.com/forderation/user-task/internal/db/model"
)

func (repo *Repository) CreateUser(ctx context.Context, input CreateUserInput) (output CreateUserOutput, err error) {
	timeNow := time.Now().UTC()
	model := model.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	err = repo.Query.User.WithContext(ctx).Create(&model)
	if err != nil {
		return
	}
	output.ID = model.ID
	return
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutput struct {
	ID int32
}
