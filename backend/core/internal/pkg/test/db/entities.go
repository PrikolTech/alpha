package test_db

import (
	"time"

	"github.com/google/uuid"
)

const (
	TableUser = "_user"
)

type User struct {
	ID         uuid.UUID `db:"id"`
	Email      string    `db:"email" fake:"{uuid}"`
	FirstName  string    `db:"first_name"`
	MiddleName *string   `db:"middle_name"`
	LastName   string    `db:"last_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
