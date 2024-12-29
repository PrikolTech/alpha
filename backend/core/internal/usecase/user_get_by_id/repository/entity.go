package repository

import (
	"time"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
	"github.com/google/uuid"
)

type entity struct {
	ID         uuid.UUID `db:"id"`
	Email      string    `db:"email"`
	FirstName  string    `db:"first_name"`
	MiddleName *string   `db:"middle_name"`
	LastName   string    `db:"last_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (e *entity) toDomain() *domain.User {
	return &domain.User{
		ID:         e.ID,
		Email:      e.Email,
		FirstName:  e.FirstName,
		MiddleName: e.MiddleName,
		LastName:   e.LastName,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}
