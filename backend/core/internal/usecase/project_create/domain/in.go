package domain

type ProjectCreateIn struct {
	Name        string
	Description *string
	Code        string
}

func (i *ProjectCreateIn) Validate() error {
	if i.Name == "" {
		return NewValidationError("name", ErrEmptyValue)
	}

	if i.Description != nil && *i.Description == "" {
		return NewValidationError("description", ErrEmptyValue)
	}

	if i.Code == "" {
		return NewValidationError("code", ErrEmptyValue)
	}

	return nil
}
