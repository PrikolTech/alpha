package domain

import "time"

const MaxPerPage = 500

type UserListIn struct {
	Page    int
	PerPage int
	Filters UserListFilters
}

type UserListFilters struct {
	Email      *string
	FirstName  *string
	MiddleName *string
	LastName   *string
	CreatedAt  *DateTimeFilter
}

type DateTimeFilter struct {
	Start *time.Time
	End   *time.Time
}

func (i *UserListIn) Validate() error {
	if i.Page < 1 {
		return NewValidationError("page", ErrPageValue)
	}

	if i.PerPage < 1 || i.PerPage > MaxPerPage {
		return NewValidationError("perPage", ErrPerPageValue)
	}

	return nil
}
