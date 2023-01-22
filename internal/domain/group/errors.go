package group

import "errors"

var (
	ErrAppendSelf   = errors.New("try to append self")
	ErrNotContained = errors.New("Group does not contain given security principal")
	ErrEmpty        = errors.New("Group is empty")
)
