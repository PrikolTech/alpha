package common

type DomainError struct {
	Msg string
}

func NewDomainError(msg string) *DomainError {
	return &DomainError{Msg: msg}
}

func (e *DomainError) Error() string {
	return e.Msg
}
