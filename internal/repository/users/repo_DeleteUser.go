package users

import (
	"context"
)

func (repo *Repository) DeleteUser(ctx context.Context, input DeleteUserInput) (err error) {
	user := repo.Query.User
	query := user.WithContext(ctx).Where(
		user.ID.Eq(input.ID),
	)
	_, err = query.Delete()
	return
}

type DeleteUserInput struct {
	ID int32
}
