package domain

import (
	"errors"
	"fmt"
)

var (
	ErrPage             = errors.New("must be positive")
	ErrPerPage          = fmt.Errorf("must be between 1 and %d", MaxPerPage)
	ErrSortingField     = errors.New("there is no such field")
	ErrSortingDirection = errors.New("direction must be 'asc' or 'desc'")
)

type ValidationError struct {
	Field  string
	Reason error
}

func NewValidationError(field string, err error) *ValidationError {
	return &ValidationError{Field: field, Reason: err}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("value of '%s' is invalid: %s", e.Field, e.Reason)
}
