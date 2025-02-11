package project_create

import (
	"context"
	"errors"
	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
)

type Handler struct {
	projectUsecase projectUsecase
}

func New(projectUsecase projectUsecase) *Handler {
	return &Handler{projectUsecase: projectUsecase}
}

func (h *Handler) Handle(ctx context.Context, req *api.ProjectCreateRequest) (api.ProjectCreateRes, error) {
	if err := h.projectUsecase.Handle(ctx, convertDtoToDomain(req)); err != nil {
		var (
			validationErr *domain.ValidationError
			domainErr     *common.DomainError
		)
		if errors.As(err, &validationErr) {
			res := &api.ProjectValidationError{
				Field:  validationErr.Field,
				Reason: validationErr.Reason.Error(),
			}
			return res, nil
		} else if errors.As(err, &domainErr) {
			res := &api.DomainError{
				Message: err.Error(),
			}
			return res, nil
		}

		return nil, err
	}
	return &api.ProjectCreateCreated{}, nil
}

func convertDtoToDomain(req *api.ProjectCreateRequest) domain.ProjectCreateIn {
	in := domain.ProjectCreateIn{
		Name: req.Name,
		Code: req.Code,
	}

	if req.Description.IsNull() {
		in.Description = ptr.To(req.Description.Value)
	}

	return in
}
