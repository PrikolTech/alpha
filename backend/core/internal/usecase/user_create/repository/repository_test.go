package repository

import (
	"context"
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	test_db "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/db"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
)

func TestRepository_Create(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	in := domain.UserCreateIn{
		Email:      gofakeit.UUID(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: ptr.To(gofakeit.MiddleName()),
		LastName:   gofakeit.LastName(),
	}

	repo := New(c.DB(), trmsqlx.DefaultCtxGetter)
	err = repo.Create(ctx, in)
	require.NoError(t, err)

	users, err := test_db.SelectEntitiesByIdCol[test_db.User](c, test_db.TableUser, "email", []string{in.Email})
	require.NoError(t, err)

	require.Len(t, users, 1)

	user := users[0]
	defer func() { require.NoError(t, test_db.DeleteEntityById(c, test_db.TableUser, user.ID)) }()

	require.Equal(t, in.Email, user.Email)
	require.Equal(t, in.FirstName, user.FirstName)
	require.Equal(t, *in.MiddleName, *user.MiddleName)
	require.Equal(t, in.LastName, user.LastName)
}

func TestRepository_ExistsByEmail(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	user, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
		entity.Email = gofakeit.UUID()
	})
	require.NoError(t, err)

	repo := New(c.DB(), trmsqlx.DefaultCtxGetter)

	t.Run("NotExists", func(t *testing.T) {
		exists, err := repo.ExistsByEmail(ctx, user.Email)
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("Exists", func(t *testing.T) {
		id, err := test_db.InsertEntityById[test_db.User, uuid.UUID](c, test_db.TableUser, user)
		require.NoError(t, err)
		defer func() { require.NoError(t, test_db.DeleteEntityById(c, test_db.TableUser, id)) }()

		exists, err := repo.ExistsByEmail(ctx, user.Email)
		require.NoError(t, err)
		require.True(t, exists)
	})
}
