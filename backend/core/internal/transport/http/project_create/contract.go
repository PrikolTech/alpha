//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go

package project_create

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
)

type projectUsecase interface {
	Handle(ctx context.Context, in domain.ProjectCreateIn) error
}
