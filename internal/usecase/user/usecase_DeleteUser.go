package user

import (
	"context"

	"github.com/forderation/user-task/internal/repository/users"
)

func (uc *Usecase) DeleteUser(ctx context.Context, input DeleteUserInput) (err error) {
	err = uc.UserRepository.DeleteUser(ctx, users.DeleteUserInput{
		ID: input.ID,
	})
	return
}

type DeleteUserInput struct {
	ID int32
}
