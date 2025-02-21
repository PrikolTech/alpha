package domain

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
