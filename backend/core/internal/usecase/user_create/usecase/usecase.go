package usecase

import (
	"context"
	"fmt"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

type usecase struct {
	userRepo userRepo
}

func New(userRepo userRepo) *usecase {
	return &usecase{userRepo: userRepo}
}

func (u *usecase) Handle(ctx context.Context, in domain.UserCreateIn) error {
	if err := in.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	exists, err := u.userRepo.ExistsByEmail(ctx, in.Email)
	if err != nil {
		return fmt.Errorf("user exists by email: %w", err)
	}

	if exists {
		return domain.NewDomainError("user with email exists")
	}

	err = u.userRepo.Create(ctx, in)
	if err != nil {
		return fmt.Errorf("user repo create: %w", err)
	}

	return nil
}
