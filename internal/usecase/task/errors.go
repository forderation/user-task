package task

import "errors"

var (
	ErrTaskNotFound error = errors.New("task not found")
)
