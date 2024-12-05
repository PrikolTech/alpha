//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package usecase

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

type userRepo interface {
	Create(ctx context.Context, in domain.UserCreateIn) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
