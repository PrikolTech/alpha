package user_get_by_id

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		var expectedUser domain.User
		require.NoError(t, gofakeit.Struct(&expectedUser))

		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(&expectedUser, nil)

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, api.UserGetByIdParams{ID: expectedUser.ID})
		require.NoError(t, err)

		user, ok := res.(*api.User)
		require.True(t, ok)
		require.Equal(t, expectedUser.Email, user.Email)
	})

	t.Run("userUsecase_DomainError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(nil, &common.DomainError{})

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, api.UserGetByIdParams{})
		require.NoError(t, err)
		require.IsType(t, &api.DomainError{}, res)
	})

	t.Run("userUsecase_InternalError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(nil, gofakeit.Error())

		handler := New(userUsecase)

		_, err := handler.Handle(ctx, api.UserGetByIdParams{})
		require.Error(t, err)
	})
}
