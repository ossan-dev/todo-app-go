package error_handler

import "errors"

var (
	ErrNotFound     error = errors.New("todo not found")
	ErrTodoNotValid error = errors.New("todo not valid")
)
