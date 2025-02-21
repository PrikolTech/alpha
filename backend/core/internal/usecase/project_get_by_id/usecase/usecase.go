package usecase

import (
	"context"
	"fmt"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_get_by_id/domain"
	"github.com/google/uuid"
)

type Usecase struct {
	projectRepo projectRepo
}

func New(projectRepo projectRepo) *Usecase {
	return &Usecase{projectRepo: projectRepo}
}

func (u *Usecase) GetProjectByID(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
	project, err := u.projectRepo.GetProjectById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("project get by id %w", err)
	}

	if project == nil {
		return nil, nil
	}

	return project, nil
}
