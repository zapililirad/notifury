package app

import "errors"

var (
	ErrGeneral        = errors.New("general error")
	ErrNotImplemented = errors.New("method does not implemented")
)
