package domain

import "time"

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
