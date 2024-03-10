package users

import (
	"context"
	"time"
)

func (repo *Repository) GetUsers(ctx context.Context, input GetUsersInput) (output GetUsersOutput, err error) {
	user := repo.Query.User
	query := user.WithContext(ctx).Select(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	models, err := query.Find()
	if err != nil {
		return
	}
	output = GetUsersOutput{
		Items: make([]GetUsersOutputItem, 0),
	}
	for _, model := range models {
		output.Items = append(output.Items, GetUsersOutputItem{
			ID:        model.ID,
			Name:      model.Name,
			Email:     model.Email,
			Password:  model.Password,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}
	return
}

type GetUsersInput struct {
}

type GetUsersOutput struct {
	Items []GetUsersOutputItem
}

type GetUsersOutputItem struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
