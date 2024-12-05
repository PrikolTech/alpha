package usecase

import (
	"context"
	"testing"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_usecase_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: ptr.To(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, in.Email).Return(false, nil)
		userRepo.EXPECT().Create(ctx, in).Return(nil)

		usecase := New(userRepo)
		err := usecase.Handle(ctx, in)
		require.NoError(t, err)
	})

	t.Run("ValidationError", func(t *testing.T) {
		in := domain.UserCreateIn{}

		userRepo := NewMockuserRepo(ctrl)

		usecase := New(userRepo)
		err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("EmailExistsError", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: ptr.To(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, in.Email).Return(true, nil)

		usecase := New(userRepo)
		err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("RepoExistsByEmailError", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: ptr.To(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}
		expectedError := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, gomock.Any()).Return(false, expectedError)

		usecase := New(userRepo)
		err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("RepoCreateError", func(t *testing.T) {
		in := domain.UserCreateIn{
			Email:      gofakeit.Email(),
			FirstName:  gofakeit.FirstName(),
			MiddleName: ptr.To(gofakeit.MiddleName()),
			LastName:   gofakeit.LastName(),
		}
		expectedError := gofakeit.Error()

		userRepo := NewMockuserRepo(ctrl)
		userRepo.EXPECT().ExistsByEmail(ctx, in.Email).Return(false, nil)
		userRepo.EXPECT().Create(ctx, in).Return(expectedError)

		usecase := New(userRepo)
		err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedError)
	})
}
