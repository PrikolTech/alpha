//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package usecase

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_all/domain"
)

type userRepo interface {
	GetAll(ctx context.Context, in domain.UserGetAllIn) ([]domain.User, error)
	GetTotalCount(ctx context.Context) (int, error)
}
