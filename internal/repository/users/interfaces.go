package users

import "context"

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, input CreateUserInput) (output CreateUserOutput, err error)
	DeleteUser(ctx context.Context, input DeleteUserInput) (err error)
	GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error)
	GetUsers(ctx context.Context, input GetUsersInput) (output GetUsersOutput, err error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (err error)
}
