//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package usecase

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
)

type userRepo interface {
	Get(ctx context.Context, in domain.UserListIn) ([]domain.User, error)
	GetTotalCount(ctx context.Context) (int, error)
}
