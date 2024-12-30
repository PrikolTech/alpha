package repository

import (
	"context"
	"testing"
	"time"

	test_db "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/db"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	"github.com/PrikolTech/alpha/backend/core/pkg/ptr"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestRepository_Get_Filters(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	repo := New(c.DB())

	tests := []struct {
		name  string
		in    domain.UserListIn
		setup func() ([]test_db.User, error)
	}{
		{
			name: "FiltersEmail",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 5,
				Filters: domain.UserListFilters{
					Email: ptr.To("Test"),
				},
			},
			setup: func() ([]test_db.User, error) {
				return test_db.GenerateEntities[test_db.User](2, func(entity *test_db.User) {
					entity.Email = "_test_" + gofakeit.UUID()
				})
			},
		},
		{
			name: "FiltersFirstName",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 5,
				Filters: domain.UserListFilters{
					FirstName: ptr.To("Test"),
				},
			},
			setup: func() ([]test_db.User, error) {
				return test_db.GenerateEntities[test_db.User](2, func(entity *test_db.User) {
					entity.FirstName = "_test_" + gofakeit.FirstName()
				})
			},
		},
		{
			name: "FiltersMiddleName",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 5,
				Filters: domain.UserListFilters{
					MiddleName: ptr.To("Test"),
				},
			},
			setup: func() ([]test_db.User, error) {
				return test_db.GenerateEntities[test_db.User](2, func(entity *test_db.User) {
					entity.MiddleName = ptr.To("_test_" + gofakeit.MiddleName())
				})
			},
		},
		{
			name: "FiltersLastName",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 5,
				Filters: domain.UserListFilters{
					LastName: ptr.To("Test"),
				},
			},
			setup: func() ([]test_db.User, error) {
				return test_db.GenerateEntities[test_db.User](2, func(entity *test_db.User) {
					entity.LastName = "_test_" + gofakeit.LastName()
				})
			},
		},
		{
			name: "FiltersCreatedAt",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 6,
				Filters: domain.UserListFilters{
					CreatedAt: &domain.DateTimeFilter{
						Start: ptr.To(lo.Must1(time.Parse(time.DateOnly, "2001-02-03"))),
						End:   ptr.To(lo.Must1(time.Parse(time.DateOnly, "2001-02-05"))),
					},
				},
			},
			setup: func() ([]test_db.User, error) {
				user1, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
					entity.CreatedAt = lo.Must1(time.Parse(time.DateOnly, "2001-02-03"))
				})
				if err != nil {
					return nil, err
				}
				user2, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
					entity.CreatedAt = lo.Must1(time.Parse(time.DateOnly, "2001-02-04"))
				})
				if err != nil {
					return nil, err
				}
				user3, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
					entity.CreatedAt = lo.Must1(time.Parse(time.DateOnly, "2001-02-05"))
				})
				if err != nil {
					return nil, err
				}
				return []test_db.User{user1, user2, user3}, nil
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			notExpectedUsers, err := test_db.GenerateEntities[test_db.User](3, func(entity *test_db.User) {
				entity.CreatedAt = time.Time{}
			})
			require.NoError(t, err)

			expectedUsers, err := test.setup()
			require.NoError(t, err)
			expectedIDs := lo.Map(expectedUsers, func(item test_db.User, _ int) string { return item.ID.String() })

			totalIDs, err := test_db.InsertEntitiesWithId[test_db.User, uuid.UUID](c, test_db.TableUser, append(expectedUsers, notExpectedUsers...))
			require.NoError(t, err)
			defer func() {
				require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, totalIDs))
			}()

			users, err := repo.Get(ctx, test.in)
			require.NoError(t, err)

			ids := lo.Map(users, func(item domain.User, _ int) string { return item.ID.String() })

			require.Equal(t, expectedIDs, ids)
		})
	}
}

func TestRepository_Get_Pagination(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	repo := New(c.DB())

	tests := []struct {
		name        string
		in          domain.UserListIn
		expectedLen int
	}{
		{
			name: "Page=1,PerPage=5",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 5,
			},
			expectedLen: 5,
		},
		{
			name: "Page=1,PerPage=3",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 3,
			},
			expectedLen: 3,
		},
		{
			name: "Page=2,PerPage=3",
			in: domain.UserListIn{
				Page:    2,
				PerPage: 3,
			},
			expectedLen: 2,
		},
		{
			name: "Page=3,PerPage=3",
			in: domain.UserListIn{
				Page:    3,
				PerPage: 3,
			},
			expectedLen: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			totalUsers, err := test_db.GenerateEntities[test_db.User](5)
			require.NoError(t, err)

			totalIDs, err := test_db.InsertEntitiesWithId[test_db.User, uuid.UUID](c, test_db.TableUser, totalUsers)
			require.NoError(t, err)
			defer func() {
				require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, totalIDs))
			}()

			users, err := repo.Get(ctx, test.in)
			require.NoError(t, err)
			require.Len(t, users, test.expectedLen)
		})
	}
}

func TestRepository_GetTotalCount(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	users, err := test_db.GenerateEntities[test_db.User](5)
	require.NoError(t, err)

	ids, err := test_db.InsertEntitiesById[test_db.User, uuid.UUID](c, test_db.TableUser, users)
	require.NoError(t, err)
	defer func() { require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, ids)) }()

	repo := New(c.DB())
	totalCount, err := repo.GetTotalCount(ctx)
	require.NoError(t, err)
	require.Equal(t, len(users), totalCount)
}
