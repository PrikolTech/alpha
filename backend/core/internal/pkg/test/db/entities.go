package test_db

import (
	"time"

	"github.com/google/uuid"
)

const (
	TableUser    = "_user"
	TableProject = "project"
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

type Project struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Code        string    `db:"code"`
	IsArchived  bool      `db:"is_archived"`
}
