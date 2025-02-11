package repository

import (
	"context"
	test_db "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/db"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_Create(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	in := domain.ProjectCreateIn{
		Name:        gofakeit.Name(),
		Description: ptr.To(gofakeit.ProductDescription()),
		Code:        gofakeit.AppName(),
	}

	repo := New(c.DB(), trmsqlx.DefaultCtxGetter)
	err = repo.Create(ctx, in)
	require.NoError(t, err)

	projects, err := test_db.SelectEntitiesByIdCol[test_db.Project](c, test_db.TableProject, "code", []string{in.Code})
	require.NoError(t, err)

	require.Len(t, projects, 1)

	project := projects[0]
	defer func() { require.NoError(t, test_db.DeleteEntityById(c, test_db.TableProject, project.ID)) }()

	require.Equal(t, in.Name, project.Name)
	require.Equal(t, in.Description, project.Description)
	require.Equal(t, in.Code, project.Code)
}

func TestRepository_ExistsByCode(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	project, err := test_db.GenerateEntity[test_db.Project]()
	require.NoError(t, err)
	project.Code = gofakeit.AppName()

	repo := New(c.DB(), trmsqlx.DefaultCtxGetter)

	t.Run("NotExists", func(t *testing.T) {
		exists, err := repo.ExistsByCode(ctx, project.Code)
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("Exists", func(t *testing.T) {
		id, err := test_db.InsertEntityById[test_db.Project, uuid.UUID](c, test_db.TableProject, project)
		require.NoError(t, err)
		defer func() { require.NoError(t, test_db.DeleteEntityById(c, test_db.TableProject, id)) }()

		exists, err := repo.ExistsByCode(ctx, project.Code)
		require.NoError(t, err)
		require.True(t, exists)
	})
}
