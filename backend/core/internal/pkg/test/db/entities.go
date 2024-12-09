package test_db

import "github.com/google/uuid"

const (
	TableUser = "_user"
)

type User struct {
	ID         uuid.UUID `db:"id"`
	Email      string    `db:"email"`
	FirstName  string    `db:"first_name"`
	MiddleName *string   `db:"middle_name"`
	LastName   string    `db:"last_name"`
}
