package usecase

import (
	"context"
	"testing"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUsecase_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		id := uuid.New()
		var expectedUser domain.User
		require.NoError(t, gofakeit.Struct(&expectedUser))
		expectedUser.ID = id

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().GetByID(ctx, id).Return(&expectedUser, nil)

		usecase := New(userRepo)
		user, err := usecase.Handle(ctx, id)
		require.NoError(t, err)
		require.Equal(t, expectedUser, *user)
	})

	t.Run("DoesNotExistError", func(t *testing.T) {
		id := uuid.New()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().GetByID(ctx, id).Return(nil, nil)

		usecase := New(userRepo)
		_, err := usecase.Handle(ctx, id)
		require.Error(t, err)
	})

	t.Run("userRepo_GetByIdError", func(t *testing.T) {
		id := uuid.New()
		expectedErr := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().GetByID(ctx, id).Return(nil, expectedErr)

		usecase := New(userRepo)
		_, err := usecase.Handle(ctx, id)
		require.ErrorIs(t, err, expectedErr)
	})
}
