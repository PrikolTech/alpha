package domain

import (
	"fmt"
	"slices"
)

type SortingDirection string

const (
	SortingDirectionAsc  SortingDirection = "ASC"
	SortingDirectionDesc SortingDirection = "DESC"
)

type SortingField string

const (
	SortingFieldEmail      SortingField = "email"
	SortingFieldFirstName  SortingField = "firstName"
	SortingFieldMiddleName SortingField = "middleName"
	SortingFieldLastName   SortingField = "lastName"
	SortingFieldCreatedAt  SortingField = "createdAt"
	SortingFieldUpdatedAt  SortingField = "updatedAt"
)

var SortingFields = []SortingField{SortingFieldEmail, SortingFieldFirstName, SortingFieldMiddleName, SortingFieldLastName, SortingFieldCreatedAt, SortingFieldUpdatedAt}

type UserListSorting struct {
	Field     SortingField
	Direction SortingDirection
}

func (s UserListSorting) Validate() error {
	field := fmt.Sprintf("sorting[%s]", s.Field)

	if !slices.Contains(SortingFields, s.Field) {
		return NewValidationError(field, ErrSortingField)
	}

	if s.Direction != SortingDirectionAsc && s.Direction != SortingDirectionDesc {
		return NewValidationError(field, ErrSortingDirection)
	}

	return nil
}
