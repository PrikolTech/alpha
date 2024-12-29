package user_create_handler

import (
	"context"
	"testing"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(nil)

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, &api.UserCreateRequest{
			MiddleName: api.OptNilString{Set: true, Null: false},
		})
		require.NoError(t, err)
		require.IsType(t, &api.UserCreateCreated{}, res)
	})

	t.Run("userUsecase_ValidationError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(&domain.ValidationError{
			Reason: gofakeit.Error(),
		})

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, &api.UserCreateRequest{})
		require.NoError(t, err)
		require.IsType(t, &api.UserCreateValidationError{}, res)
	})

	t.Run("userUsecase_DomainError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(&common.DomainError{})

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, &api.UserCreateRequest{})
		require.NoError(t, err)
		require.IsType(t, &api.DomainError{}, res)
	})

	t.Run("userUsecase_InternalError", func(t *testing.T) {
		userUsecase := NewMockuserUsecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(gofakeit.Error())

		handler := New(userUsecase)

		_, err := handler.Handle(ctx, &api.UserCreateRequest{})
		require.Error(t, err)
	})
}
