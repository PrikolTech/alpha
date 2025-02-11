package user_list

import (
	"context"
	"errors"
	"strings"

	"github.com/samber/lo"

	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
)

type Handler struct {
	userUsecase userUsecase
}

func New(userUsecase userUsecase) *Handler {
	return &Handler{userUsecase: userUsecase}
}

func (h *Handler) Handle(ctx context.Context, params api.UserListParams) (api.UserListRes, error) {
	out, err := h.userUsecase.Handle(ctx, h.convertParamsToDomain(params))
	if err != nil {
		var validationErr *domain.ValidationError
		if errors.As(err, &validationErr) {
			res := &api.UserValidationError{
				Field:  validationErr.Field,
				Reason: validationErr.Reason.Error(),
			}
			return res, nil
		}
		return nil, err
	}

	return h.convertDomainToDto(out), nil
}

func (h *Handler) convertParamsToDomain(params api.UserListParams) domain.UserListIn {
	in := domain.UserListIn{
		Page:    params.Page.Or(1),
		PerPage: params.PerPage.Or(20),
		Sorting: h.convertDtoToSorting(params.Sorting),
	}

	in.Filters = domain.UserListFilters{
		CreatedAt: h.convertDtoToDateTimeFilter(params.CreatedAt),
		UpdatedAt: h.convertDtoToDateTimeFilter(params.UpdatedAt),
	}

	if params.Email.IsSet() {
		in.Filters.Email = ptr.To(params.Email.Value)
	}

	if params.FirstName.IsSet() {
		in.Filters.FirstName = ptr.To(params.FirstName.Value)
	}

	if params.MiddleName.IsSet() {
		in.Filters.MiddleName = ptr.To(params.MiddleName.Value)
	}

	if params.LastName.IsSet() {
		in.Filters.LastName = ptr.To(params.LastName.Value)
	}

	return in
}

func (h *Handler) convertDtoToSorting(dto api.OptSorting) *domain.UserListSorting {
	in, found := dto.Get()
	if !found {
		return nil
	}

	return &domain.UserListSorting{
		Field:     domain.SortingField(in.GetField()),
		Direction: domain.SortingDirection(strings.ToUpper(string(in.GetDirection()))),
	}
}

func (h *Handler) convertDtoToDateTimeFilter(dto api.OptDateTimeFilter) *domain.DateTimeFilter {
	in, found := dto.Get()
	if !found {
		return nil
	}
	var out domain.DateTimeFilter
	if start, found := in.Start.Get(); found {
		out.Start = ptr.To(start)
	}
	if end, found := in.End.Get(); found {
		out.End = ptr.To(end)
	}
	return &out
}

func (h *Handler) convertDomainToDto(out *domain.UserListOut) *api.UserListResponse {
	dto := api.UserListResponse{
		Data: lo.Map(out.Data, h.convertDomainUserToDto),
		Meta: api.Meta{
			Page:         out.Meta.Page,
			TotalPages:   out.Meta.TotalPages,
			PerPage:      out.Meta.PerPage,
			TotalRecords: out.Meta.TotalRecords,
		},
	}
	return &dto
}

func (h *Handler) convertDomainUserToDto(user domain.User, _ int) api.User {
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

	return dto
}
