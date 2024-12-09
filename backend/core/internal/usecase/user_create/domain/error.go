package domain

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyValue = errors.New("value is empty")
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

type DomainError struct {
	Msg string
}

func NewDomainError(msg string) *DomainError {
	return &DomainError{Msg: msg}
}

func (e *DomainError) Error() string {
	return e.Msg
}
