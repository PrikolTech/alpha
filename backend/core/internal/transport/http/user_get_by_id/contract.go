//go:generate go tool mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package user_get_by_id

import (
	"context"

	"github.com/google/uuid"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
)

type usecase interface {
	Handle(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
