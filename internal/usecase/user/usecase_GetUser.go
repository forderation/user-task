package user

import (
	"context"
	"errors"
	"time"

	"github.com/forderation/null"
	"github.com/forderation/user-task/internal/repository/users"
	"gorm.io/gorm"
)

func (usecase *Usecase) GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error) {
	user, err := usecase.UserRepository.GetUser(ctx, users.GetUserInput{
		ID:      input.ID,
		Email:   input.Email,
		NotEqID: input.NotEqID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.Join(err, ErrUserNotFound)
		}
		return
	}
	output = GetUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return
}

type GetUserInput struct {
	ID      null.Int32
	Email   null.String
	NotEqID null.Int32
}

type GetUserOutput struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
