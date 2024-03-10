package users

import (
	"context"
	"time"

	"github.com/forderation/user-task/internal/db"
	"github.com/forderation/user-task/util/pagination"
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
	totalData, err := query.Count()
	if err != nil {
		return
	}
	db.ApplyPaginationQuery(query.DO, input.Pagination)
	models, err := query.Find()
	if err != nil {
		return
	}
	output = GetUsersOutput{
		Items:      make([]GetUsersOutputItem, 0),
		Pagination: pagination.CreatePaginationOutput(input.Pagination, totalData),
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
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
