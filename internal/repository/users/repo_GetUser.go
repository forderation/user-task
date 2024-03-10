package users

import (
	"context"
	"time"

	"github.com/forderation/null"
)

func (repo *Repository) GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error) {
	user := repo.Query.User
	query := user.WithContext(ctx).Select(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if input.ID.Valid {
		query = query.Where(user.ID.Eq(input.ID.Int32))
	}
	if input.Email.Valid {
		query = query.Where(user.Email.Eq(input.Email.String))
	}
	if input.NotEqID.Valid {
		query = query.Where(user.ID.Neq(input.NotEqID.Int32))
	}
	model, err := query.First()
	if err != nil {
		return
	}
	output = GetUserOutput{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
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
