package user_get_by_id_handler

import (
	"context"
	"errors"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
)

type Handler struct {
	userUsecase userUsecase
}

func New(userUsecase userUsecase) *Handler {
	return &Handler{userUsecase: userUsecase}
}

func (h *Handler) Handle(ctx context.Context, params api.UserGetByIdParams) (api.UserGetByIdRes, error) {
	user, err := h.userUsecase.Handle(ctx, params.ID)
	if err != nil {
		var domainErr *common.DomainError
		if errors.As(err, &domainErr) {
			res := &api.DomainError{
				Message: err.Error(),
			}
			return res, nil
		}
		return nil, err
	}

	return h.convertDomainToDto(user), nil
}

func (h *Handler) convertDomainToDto(user *domain.User) *api.User {
	dto := api.User{
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
