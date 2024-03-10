package user

import (
	"context"
	"time"

	"github.com/forderation/user-task/internal/repository/users"
	"github.com/forderation/user-task/util/pagination"
)

func (uc *Usecase) GetUsers(ctx context.Context, input GetUsersInput) (output GetUsersOutput, err error) {
	usersOutput, err := uc.UserRepository.GetUsers(ctx, users.GetUsersInput{
		Pagination: input.Pagination,
	})
	if err != nil {
		return
	}
	output = GetUsersOutput{
		Pagination: usersOutput.Pagination,
		Items:      make([]GetUsersOutputItem, 0),
	}
	for _, user := range usersOutput.Items {
		output.Items = append(output.Items, GetUsersOutputItem{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return
}

type GetUsersInput struct {
	Pagination pagination.PaginationInput
}

type GetUsersOutput struct {
	Pagination pagination.PaginationOutput
	Items      []GetUsersOutputItem
}

type GetUsersOutputItem struct {
	ID        int32
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
