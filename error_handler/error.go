package error_handler

import "errors"

var (
	ErrNotFound     error = errors.New("todo not found")
	ErrTodoNotValid error = errors.New("todo not valid")
)

const (
	BAD_REQUEST_CODE = "BAD_REQUEST"

	INVALID_BODY_REQUEST_MSG = "Invalid body request"
)

type TodoAppError struct {
	Msg     string   `json:"msg"`
	Code    string   `json:"code"`
	Details []string `json:"details,omitempty"`
}
