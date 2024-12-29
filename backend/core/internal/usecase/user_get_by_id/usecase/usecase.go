package usecase

import (
	"context"
	"fmt"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
	"github.com/google/uuid"
)

type Usecase struct {
	userRepo userRepo
}

func New(userRepo userRepo) *Usecase {
	return &Usecase{userRepo: userRepo}
}

func (u *Usecase) Handle(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user get by id: %w", err)
	}

	if user == nil {
		return nil, common.NewDomainError("user does not exist")
	}

	return user, nil
}
