package user

import (
	"context"
	"errors"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/repository/users"
	"gorm.io/gorm"
)

func (uc *Usecase) UpdateUser(ctx context.Context, input UpdateUserInput) (err error) {
	existUser, err := uc.UserRepository.GetUser(ctx, users.GetUserInput{
		ID: null.Int32From(input.ID),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.Join(err, ErrUserNotFound)
		}
		return
	}
	err = uc.UserRepository.UpdateUser(ctx, users.UpdateUserInput{
		ID: existUser.ID,
		Payload: users.UpdateUserInputPayload{
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
		},
	})
	return
}

type UpdateUserInput struct {
	ID       int32
	Name     null.String
	Email    null.String
	Password null.String
}
