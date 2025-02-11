package user_create

import (
	"context"
	"errors"

	"github.com/samber/lo"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

type Handler struct {
	userUsecase userUsecase
}

func New(userUsecase userUsecase) *Handler {
	return &Handler{userUsecase: userUsecase}
}

func (h *Handler) Handle(ctx context.Context, req *api.UserCreateRequest) (api.UserCreateRes, error) {
	err := h.userUsecase.Handle(ctx, h.convertDtoToDomain(req))
	if err != nil {
		var (
			validationErr *domain.ValidationError
			domainErr     *common.DomainError
		)
		if errors.As(err, &validationErr) {
			res := &api.UserValidationError{
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

	return &api.UserCreateCreated{}, nil
}

func (h *Handler) convertDtoToDomain(req *api.UserCreateRequest) domain.UserCreateIn {
	in := domain.UserCreateIn{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if req.MiddleName.IsSet() && !req.MiddleName.IsNull() {
		in.MiddleName = lo.ToPtr(req.MiddleName.Value)
	}

	return in
}
