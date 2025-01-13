package domain

type UserCreateIn struct {
	Email      string
	FirstName  string
	MiddleName *string
	LastName   string
}

func (i *UserCreateIn) Validate() error {
	if i.Email == "" {
		return NewValidationError("email", ErrEmptyValue)
	}

	if i.FirstName == "" {
		return NewValidationError("firstname", ErrEmptyValue)
	}

	if i.LastName == "" {
		return NewValidationError("lastname", ErrEmptyValue)
	}

	if i.MiddleName != nil && *i.MiddleName == "" {
		return NewValidationError("middlename", ErrEmptyValue)
	}

	return nil
}
