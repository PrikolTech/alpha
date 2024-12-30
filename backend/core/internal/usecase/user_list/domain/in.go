package domain

const MaxPerPage = 500

type UserListIn struct {
	Page    int
	PerPage int
	Filters UserListFilters
	Sorting *UserListSorting
}

func (i *UserListIn) Validate() error {
	if i.Page < 1 {
		return NewValidationError("page", ErrPage)
	}

	if i.PerPage < 1 || i.PerPage > MaxPerPage {
		return NewValidationError("perPage", ErrPerPage)
	}

	if i.Sorting != nil {
		return i.Sorting.Validate()
	}

	return nil
}
