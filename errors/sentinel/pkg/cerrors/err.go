package cerrors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrTooLarge = errors.New("too large")
)
