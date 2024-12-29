package domain

const MaxPerPage = 500

type UserGetAllIn struct {
	Page    int
	PerPage int
}

func (i *UserGetAllIn) Validate() error {
	if i.Page < 1 {
		return NewValidationError("page", ErrPageValue)
	}

	if i.PerPage < 1 || i.PerPage > MaxPerPage {
		return NewValidationError("perPage", ErrPerPageValue)
	}

	return nil
}
