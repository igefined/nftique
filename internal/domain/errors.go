package domain

import "errors"

var (
	ErrEntityNotFound  = errors.New("entity not found")
	ErrEntityDuplicate = errors.New("duplicate entity value")
)
