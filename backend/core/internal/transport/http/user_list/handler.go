package user_list

import (
	"context"
	"errors"
	"strings"

	"github.com/samber/lo"

	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
)

type Handler struct {
	usecase usecase
}

func New(usecase usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Handle(ctx context.Context, params api.UserListParams) (api.UserListRes, error) {
	out, err := h.usecase.Handle(ctx, h.convertParamsToDomain(params))
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
		in.Filters.Email = lo.ToPtr(params.Email.Value)
	}

	if params.FirstName.IsSet() {
		in.Filters.FirstName = lo.ToPtr(params.FirstName.Value)
	}

	if params.MiddleName.IsSet() {
		in.Filters.MiddleName = lo.ToPtr(params.MiddleName.Value)
	}

	if params.LastName.IsSet() {
		in.Filters.LastName = lo.ToPtr(params.LastName.Value)
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
		out.Start = lo.ToPtr(start)
	}
	if end, found := in.End.Get(); found {
		out.End = lo.ToPtr(end)
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
