//go:generate go tool mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go
package usecase

import (
	"context"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_get_by_id/domain"
	"github.com/google/uuid"
)

type projectRepo interface {
	GetProjectById(ctx context.Context, id uuid.UUID) (*domain.Project, error)
}
