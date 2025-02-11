package usecase

import (
	"context"
	test_trm "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/trm"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUsecase_Handler(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trm := test_trm.New()

	t.Run("success", func(t *testing.T) {
		in := domain.ProjectCreateIn{
			Name:        gofakeit.Name(),
			Description: ptr.To(gofakeit.ProductDescription()),
			Code:        gofakeit.UUID(),
		}

		projectRepo := NewMockprojectRepo(ctrl)
		projectRepo.EXPECT().ExistsByCode(ctx, in.Code).Return(false, nil)
		projectRepo.EXPECT().Create(ctx, in).Return(nil).Times(1)

		usecase := New(projectRepo, trm)
		err := usecase.Handle(ctx, in)
		require.NoError(t, err)
	})

	t.Run("project already exists", func(t *testing.T) {
		in := domain.ProjectCreateIn{
			Name:        gofakeit.Name(),
			Description: ptr.To(gofakeit.ProductDescription()),
			Code:        gofakeit.UUID(),
		}
		projectRepo := NewMockprojectRepo(ctrl)
		projectRepo.EXPECT().ExistsByCode(ctx, in.Code).Return(true, nil)

		usecase := New(projectRepo, trm)
		err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("validation error", func(t *testing.T) {
		in := domain.ProjectCreateIn{}
		projectRepo := NewMockprojectRepo(ctrl)
		usecase := New(projectRepo, trm)
		err := usecase.Handle(ctx, in)
		require.Error(t, err)
	})

	t.Run("error from ExistsByCode", func(t *testing.T) {
		in := domain.ProjectCreateIn{
			Name:        gofakeit.Name(),
			Description: ptr.To(gofakeit.ProductDescription()),
			Code:        gofakeit.UUID(),
		}
		expectedErr := gofakeit.Error()

		projectRepo := NewMockprojectRepo(ctrl)
		projectRepo.EXPECT().ExistsByCode(ctx, in.Code).Return(false, expectedErr)

		usecase := New(projectRepo, trm)
		err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("error from Create", func(t *testing.T) {
		in := domain.ProjectCreateIn{
			Name:        gofakeit.Name(),
			Description: ptr.To(gofakeit.ProductDescription()),
			Code:        gofakeit.UUID(),
		}
		expectedErr := gofakeit.Error()

		projectRepo := NewMockprojectRepo(ctrl)
		projectRepo.EXPECT().ExistsByCode(ctx, in.Code).Return(false, nil)
		projectRepo.EXPECT().Create(ctx, in).Return(expectedErr)

		usecase := New(projectRepo, trm)
		err := usecase.Handle(ctx, in)
		require.ErrorIs(t, err, expectedErr)
	})
}
