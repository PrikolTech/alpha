package usecase

import (
	"context"
	"fmt"
	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type Usecase struct {
	projectRepo projectRepo
	trm         trm.Manager
}

func New(projectRepo projectRepo, trm trm.Manager) *Usecase {
	return &Usecase{projectRepo: projectRepo, trm: trm}
}

func (u *Usecase) Handle(ctx context.Context, in domain.ProjectCreateIn) error {
	if err := in.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return u.trm.Do(ctx, func(ctx context.Context) error {
		exists, err := u.projectRepo.ExistsByCode(ctx, in.Code)
		if err != nil {
			return fmt.Errorf("project repo exists by code: %w", err)
		}
		if exists {
			return common.NewDomainError("project with code exists")
		}

		err = u.projectRepo.Create(ctx, in)
		if err != nil {
			return fmt.Errorf("project repo create: %w", err)
		}

		return nil
	})
}
