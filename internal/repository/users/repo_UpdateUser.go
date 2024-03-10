package users

import (
	"context"
	"time"

	"github.com/forderation/null"
)

func (repo *Repository) UpdateUser(ctx context.Context, input UpdateUserInput) (err error) {
	user := repo.Query.User
	query := user.WithContext(ctx).Where(
		user.ID.Eq(input.ID),
	)
	updateParam := map[string]interface{}{}
	timeNow := time.Now().UTC()
	updateParam[user.UpdatedAt.ColumnName().String()] = timeNow
	if input.Payload.Name.Valid {
		updateParam[user.Name.ColumnName().String()] = input.Payload.Name.String
	}
	if input.Payload.Email.Valid {
		updateParam[user.Email.ColumnName().String()] = input.Payload.Email.String
	}
	if input.Payload.Password.Valid {
		updateParam[user.Password.ColumnName().String()] = input.Payload.Password.String
	}
	_, err = query.Updates(updateParam)
	return

}

type UpdateUserInput struct {
	ID      int32
	Payload UpdateUserInputPayload
}

type UpdateUserInputPayload struct {
	Name     null.String
	Email    null.String
	Password null.String
}
