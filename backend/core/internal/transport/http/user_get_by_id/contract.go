//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package user_get_by_id_handler

import (
	"context"

	"github.com/google/uuid"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
)

type userUsecase interface {
	Handle(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
