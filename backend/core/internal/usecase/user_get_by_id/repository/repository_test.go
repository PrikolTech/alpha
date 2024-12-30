package repository

import (
	"context"
	"testing"

	test_db "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetById(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	expectedUser, err := test_db.GenerateEntity[test_db.User]()
	require.NoError(t, err)

	repo := New(c.DB())

	t.Run("NotExists", func(t *testing.T) {
		user, err := repo.GetByID(ctx, expectedUser.ID)
		require.NoError(t, err)
		require.Nil(t, user)
	})

	t.Run("Exists", func(t *testing.T) {
		id, err := test_db.InsertEntityWithId[test_db.User, uuid.UUID](c, test_db.TableUser, expectedUser)
		require.NoError(t, err)
		defer func() { require.NoError(t, test_db.DeleteEntityById(c, test_db.TableUser, id)) }()

		user, err := repo.GetByID(ctx, expectedUser.ID)
		require.NoError(t, err)
		require.Equal(t, expectedUser.ID, user.ID)
		require.Equal(t, expectedUser.Email, user.Email)
	})
}
