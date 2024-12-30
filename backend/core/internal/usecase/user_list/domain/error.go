package domain

import (
	"errors"
	"fmt"
)

var (
	ErrPageValue    = errors.New("must be positive")
	ErrPerPageValue = fmt.Errorf("must be between 1 and %d", MaxPerPage)
)

type ValidationError struct {
	Field  string
	Reason error
}

func NewValidationError(field string, err error) *ValidationError {
	return &ValidationError{Field: field, Reason: err}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("value of field '%s' is invalid: %s", e.Field, e.Reason)
}
