package usecase

import (
	"context"
	"fmt"

	"github.com/avito-tech/go-transaction-manager/trm/v2"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

type Usecase struct {
	userRepo userRepo
	trm      trm.Manager
}

func New(userRepo userRepo, trm trm.Manager) *Usecase {
	return &Usecase{userRepo: userRepo, trm: trm}
}

func (u *Usecase) Handle(ctx context.Context, in domain.UserCreateIn) error {
	if err := in.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return u.trm.Do(ctx, func(ctx context.Context) error {
		exists, err := u.userRepo.ExistsByEmail(ctx, in.Email)
		if err != nil {
			return fmt.Errorf("user repo exists by email: %w", err)
		}
		if exists {
			return trm.Skippable(common.NewDomainError("user with email exists"))
		}

		err = u.userRepo.Create(ctx, in)
		if err != nil {
			return fmt.Errorf("user repo create: %w", err)
		}

		return nil
	})
}
