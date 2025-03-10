package user_get_by_id

import (
	"context"
	"errors"
	"fmt"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
)

type Handler struct {
	usecase usecase
}

func New(usecase usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Handle(ctx context.Context, params api.UserGetByIdParams) (api.UserGetByIdRes, error) {
	user, err := h.usecase.Handle(ctx, params.ID)
	if err != nil {
		var domainErr *common.DomainError
		if errors.As(err, &domainErr) {
			res := &api.DomainError{
				Message: err.Error(),
			}
			return res, nil
		}
		return nil, fmt.Errorf("user get by id: %w", err)
	}

	return h.convertDomainToDto(user), nil
}

func (h *Handler) convertDomainToDto(user *domain.User) *api.User {
	dto := api.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.MiddleName != nil {
		dto.MiddleName = api.OptNilString{
			Value: *user.MiddleName,
			Set:   true,
		}
	}

	return &dto
}
