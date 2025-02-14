package project_create

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
)

type Handler struct {
	usecase usecase
}

func New(usecase usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Handle(ctx context.Context, req *api.ProjectCreateRequest) (api.ProjectCreateRes, error) {
	if err := h.usecase.Handle(ctx, convertDtoToDomain(req)); err != nil {
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
		}
		if errors.As(err, &domainErr) {
			res := &api.DomainError{
				Message: err.Error(),
			}
			return res, nil
		}

		return nil, fmt.Errorf("project create: %w", err)
	}
	return &api.ProjectCreateCreated{}, nil
}

func convertDtoToDomain(req *api.ProjectCreateRequest) domain.ProjectCreateIn {
	in := domain.ProjectCreateIn{
		Name: req.Name,
		Code: req.Code,
	}

	if req.Description.IsNull() {
		in.Description = lo.ToPtr(req.Description.Value)
	}

	return in
}
