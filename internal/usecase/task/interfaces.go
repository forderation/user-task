package task

import "context"

type TaskUsecaseInterface interface {
	CreateTask(ctx context.Context, input CreateTaskInput) (CreateTaskOutput, error)
	GetTask(ctx context.Context, input GetTaskInput) (GetTaskOutput, error)
	GetTasks(ctx context.Context, input GetTasksInput) (GetTasksOutput, error)
	UpdateTask(ctx context.Context, input UpdateTaskInput) error
	DeleteTask(ctx context.Context, input DeleteTaskInput) error
}
