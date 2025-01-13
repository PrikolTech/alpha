package usecase

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
)

func TestUsecase_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		in := domain.UserListIn{
			Page:    1,
			PerPage: 2,
		}

		expectedOut := domain.UserListOut{
			Data: make([]domain.User, 3),
			Meta: domain.Meta{
				Page:         1,
				PerPage:      2,
				TotalPages:   2,
				TotalRecords: 3,
			},
		}

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), in).Return(expectedOut.Data, nil)
		userRepo.EXPECT().GetTotalCount(gomock.Any()).Return(expectedOut.Meta.TotalRecords, nil)

		usecase := New(userRepo)
		out, err := usecase.Handle(ctx, in)
		require.NoError(t, err)
		require.Equal(t, expectedOut, *out)
	})

	t.Run("ValidationError", func(t *testing.T) {
		in := domain.UserListIn{}

		userRepo := NewMockuserRepo(ctrl)

		usecase := New(userRepo)
		_, err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("userRepo_GetError", func(t *testing.T) {
		in := domain.UserListIn{
			Page:    1,
			PerPage: 2,
		}

		expectedError := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), in).Return(nil, expectedError)
		userRepo.EXPECT().GetTotalCount(gomock.Any()).Return(3, nil)

		usecase := New(userRepo)
		_, err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("userRepo_GetTotalCountError", func(t *testing.T) {
		in := domain.UserListIn{
			Page:    1,
			PerPage: 2,
		}

		expectedError := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), in).Return(make([]domain.User, 1), nil)
		userRepo.EXPECT().GetTotalCount(gomock.Any()).Return(0, expectedError)

		usecase := New(userRepo)
		_, err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedError)
	})
}
