package users

import "context"

//go:generate mockgen -source=interfaces.go -destination=interfaces.mock.gen.go -package=users
type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, input CreateUserInput) (output CreateUserOutput, err error)
	DeleteUser(ctx context.Context, input DeleteUserInput) (err error)
	GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error)
	GetUsers(ctx context.Context, input GetUsersInput) (output GetUsersOutput, err error)
}
