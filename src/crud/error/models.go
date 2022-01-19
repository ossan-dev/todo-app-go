package error

import "errors"

var ErrNotFound error = errors.New("todo not found")
