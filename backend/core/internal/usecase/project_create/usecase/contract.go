//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package usecase

import (
	"context"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
)

type projectRepo interface {
	ExistsByCode(ctx context.Context, code string) (bool, error)
	Create(ctx context.Context, in domain.ProjectCreateIn) error
}
