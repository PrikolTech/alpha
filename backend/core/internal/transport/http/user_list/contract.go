//go:generate go tool mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package user_list

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
)

type usecase interface {
	Handle(ctx context.Context, in domain.UserListIn) (*domain.UserListOut, error)
}
