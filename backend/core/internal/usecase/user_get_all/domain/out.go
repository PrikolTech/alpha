package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Email      string
	FirstName  string
	MiddleName *string
	LastName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Meta struct {
	Page         int
	PerPage      int
	TotalPages   int
	TotalRecords int
}

type UserGetAllOut struct {
	Data []User
	Meta Meta
}
