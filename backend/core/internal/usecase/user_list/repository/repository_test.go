package repository

import (
	"context"
	"testing"

	test_db "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/db"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetAll(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	users, err := test_db.GenerateEntities[test_db.User](
		5,
		func(entity *test_db.User) {
			entity.Email = gofakeit.UUID()
		},
	)
	require.NoError(t, err)

	ids, err := test_db.InsertEntitiesById[test_db.User, uuid.UUID](c, test_db.TableUser, users)
	require.NoError(t, err)
	defer func() { require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, ids)) }()

	// TDT: TODO

	in := domain.UserListIn{
		Page:    1,
		PerPage: 3,
	}

	repo := New(c.DB())
	gotUsers, err := repo.GetAll(ctx, in)
	require.NoError(t, err)
	require.Len(t, gotUsers, in.PerPage)
}

func TestRepository_GetTotalCount(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	users, err := test_db.GenerateEntities[test_db.User](
		5,
		func(entity *test_db.User) {
			entity.Email = gofakeit.UUID()
		},
	)
	require.NoError(t, err)

	ids, err := test_db.InsertEntitiesById[test_db.User, uuid.UUID](c, test_db.TableUser, users)
	require.NoError(t, err)
	defer func() { require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, ids)) }()

	repo := New(c.DB())
	totalCount, err := repo.GetTotalCount(ctx)
	require.NoError(t, err)
	require.Equal(t, len(users), totalCount)
}
