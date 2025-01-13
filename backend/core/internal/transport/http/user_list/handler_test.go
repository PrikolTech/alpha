package user_list_handler

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"

	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		params := api.UserListParams{
			Page:       api.NewOptInt(3),
			PerPage:    api.NewOptInt(10),
			Email:      api.NewOptString("email"),
			FirstName:  api.NewOptString("firstName"),
			MiddleName: api.NewOptString("middleName"),
			LastName:   api.NewOptString("lastName"),
			CreatedAt:  api.NewOptDateTimeFilter(api.DateTimeFilter{Start: api.OptNilDateTime{Value: gofakeit.Date(), Set: true}, End: api.OptNilDateTime{Value: gofakeit.Date(), Set: true}}),
			UpdatedAt:  api.NewOptDateTimeFilter(api.DateTimeFilter{Start: api.OptNilDateTime{Value: gofakeit.Date(), Set: true}, End: api.OptNilDateTime{Value: gofakeit.Date(), Set: true}}),
			Sorting:    api.NewOptSorting(api.Sorting{Field: "field", Direction: api.SortingDirectionAsc}),
		}

		in := domain.UserListIn{
			Page:    3,
			PerPage: 10,
			Filters: domain.UserListFilters{
				Email:      ptr.To("email"),
				FirstName:  ptr.To("firstName"),
				MiddleName: ptr.To("middleName"),
				LastName:   ptr.To("lastName"),
				CreatedAt: &domain.DateTimeFilter{
					Start: ptr.To(params.CreatedAt.Value.Start.Value),
					End:   ptr.To(params.CreatedAt.Value.End.Value),
				},
				UpdatedAt: &domain.DateTimeFilter{
					Start: ptr.To(params.UpdatedAt.Value.Start.Value),
					End:   ptr.To(params.UpdatedAt.Value.End.Value),
				},
			},
			Sorting: &domain.UserListSorting{
				Field:     "field",
				Direction: "ASC",
			},
		}

		out := domain.UserListOut{
			Data: make([]domain.User, 3),
		}

		for i := range out.Data {
			require.NoError(t, gofakeit.Struct(&out.Data[i]))
		}

		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, in).Return(&out, nil)

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, params)
		require.NoError(t, err)

		response, ok := res.(*api.UserListResponse)
		require.True(t, ok)
		require.Len(t, response.Data, len(out.Data))
		for i := range response.Data {
			require.Equal(t, out.Data[i].ID, response.Data[i].ID)
		}
	})

	t.Run("userUsecase_ValidationError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(nil, &domain.ValidationError{
			Reason: gofakeit.Error(),
		})

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, api.UserListParams{})
		require.NoError(t, err)
		require.IsType(t, &api.UserValidationError{}, res)
	})

	t.Run("userUsecase_InternalError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(nil, gofakeit.Error())

		handler := New(userUsecase)

		_, err := handler.Handle(ctx, api.UserListParams{})
		require.Error(t, err)
	})
}
