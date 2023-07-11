package group

import "errors"

var (
	ErrAppendSelf       = errors.New("try to append self")
	ErrNotContained     = errors.New("group does not contain given security principal")
	ErrEmpty            = errors.New("group is empty")
	ErrRecursiveNesting = errors.New("attemption to create recursive nesting")
)
