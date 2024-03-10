package user

import (
	"context"

	"github.com/forderation/user-task/internal/repository/users"
)

func (uc *Usecase) CreateUser(ctx context.Context, input CreateUserInput) (output CreateUserOutput, err error) {
	userOutput, err := uc.UserRepository.CreateUser(ctx, users.CreateUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return
	}
	output.ID = userOutput.ID
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
