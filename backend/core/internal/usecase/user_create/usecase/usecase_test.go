package usecase

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	test_trm "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/trm"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

func TestUsecase_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trm := test_trm.New()

	t.Run("Success", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: lo.ToPtr(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, in.Email).Return(false, nil)
		userRepo.EXPECT().Create(ctx, in).Return(nil)

		usecase := New(userRepo, trm)
		err := usecase.Handle(ctx, in)
		require.NoError(t, err)
	})

	t.Run("ValidationError", func(t *testing.T) {
		in := domain.UserCreateIn{}

		userRepo := NewMockuserRepo(ctrl)

		usecase := New(userRepo, trm)
		err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("EmailExistsError", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: lo.ToPtr(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, in.Email).Return(true, nil)

		usecase := New(userRepo, trm)
		err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("userRepo_ExistsByEmailError", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: lo.ToPtr(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}
		expectedErr := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, gomock.Any()).Return(false, expectedErr)

		usecase := New(userRepo, trm)
		err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("userRepo_CreateError", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: lo.ToPtr(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}
		expectedErr := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, in.Email).Return(false, nil)
		userRepo.EXPECT().Create(ctx, in).Return(expectedErr)

		usecase := New(userRepo, trm)
		err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedErr)
	})
}
