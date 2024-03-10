package tasks

import "context"

func (repo *Repository) DeleteTask(ctx context.Context, input DeleteTaskInput) (err error) {
	task := repo.Query.Task
	query := task.WithContext(ctx).Where(
		task.ID.Eq(input.ID),
	)
	_, err = query.Delete()
	return
}

type DeleteTaskInput struct {
	ID int32
}
