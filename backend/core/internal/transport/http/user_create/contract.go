//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package user_create

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

type userUsecase interface {
	Handle(ctx context.Context, in domain.UserCreateIn) error
}
