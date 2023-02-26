package exception

import "errors"

var (
	ErrNotFound   = errors.New("not found error")
	ErrBadRequest = errors.New("bad request error")
	ErrService    = errors.New("service error")
	ErrUnAuth     = errors.New("not authorized")
)
