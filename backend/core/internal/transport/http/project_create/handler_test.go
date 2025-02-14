package project_create

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/PrikolTech/alpha/backend/core/internal/common"
	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
)

func TestHandle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		userUsecase := NewMockusecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(nil)

		handler := New(userUsecase)

		req := &api.ProjectCreateRequest{
			Name:        gofakeit.Name(),
			Code:        gofakeit.UUID(),
			Description: api.NewNilString(gofakeit.ProductDescription()),
		}

		res, err := handler.Handle(ctx, req)
		require.NoError(t, err)
		require.IsType(t, &api.ProjectCreateCreated{}, res)
	})

	t.Run("internal_error", func(t *testing.T) {
		userUsecase := NewMockusecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(gofakeit.Error())

		handler := New(userUsecase)

		_, err := handler.Handle(ctx, &api.ProjectCreateRequest{})
		require.Error(t, err)
	})

	t.Run("validation_error", func(t *testing.T) {
		userUsecase := NewMockusecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(&domain.ValidationError{
			Reason: gofakeit.Error(),
		})

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, &api.ProjectCreateRequest{})
		require.NoError(t, err)
		require.IsType(t, &api.ProjectValidationError{}, res)
	})

	t.Run("domain_error", func(t *testing.T) {
		userUsecase := NewMockusecase(ctrl)
		userUsecase.EXPECT().Handle(ctx, gomock.Any()).Return(&common.DomainError{
			Msg: gofakeit.Error().Error(),
		})

		handler := New(userUsecase)

		res, err := handler.Handle(ctx, &api.ProjectCreateRequest{})
		require.NoError(t, err)
		require.IsType(t, &api.DomainError{}, res)
	})
}
