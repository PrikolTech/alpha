package repository

import (
	"context"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	test_db "github.com/PrikolTech/alpha/backend/core/internal/pkg/test/db"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
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
					Email: lo.ToPtr("Test"),
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
					FirstName: lo.ToPtr("Test"),
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
					MiddleName: lo.ToPtr("Test"),
				},
			},
			setup: func() ([]test_db.User, error) {
				return test_db.GenerateEntities[test_db.User](2, func(entity *test_db.User) {
					entity.MiddleName = lo.ToPtr("_test_" + gofakeit.MiddleName())
				})
			},
		},
		{
			name: "FiltersLastName",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 5,
				Filters: domain.UserListFilters{
					LastName: lo.ToPtr("Test"),
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
						Start: lo.ToPtr(lo.Must1(time.Parse(time.DateOnly, "2001-02-03"))),
						End:   lo.ToPtr(lo.Must1(time.Parse(time.DateOnly, "2001-02-05"))),
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
		{
			name: "FiltersUpdatedAt",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 6,
				Filters: domain.UserListFilters{
					UpdatedAt: &domain.DateTimeFilter{
						Start: lo.ToPtr(lo.Must1(time.Parse(time.DateOnly, "2001-02-03"))),
						End:   lo.ToPtr(lo.Must1(time.Parse(time.DateOnly, "2001-02-05"))),
					},
				},
			},
			setup: func() ([]test_db.User, error) {
				user1, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
					entity.UpdatedAt = lo.Must1(time.Parse(time.DateOnly, "2001-02-03"))
				})
				if err != nil {
					return nil, err
				}
				user2, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
					entity.UpdatedAt = lo.Must1(time.Parse(time.DateOnly, "2001-02-04"))
				})
				if err != nil {
					return nil, err
				}
				user3, err := test_db.GenerateEntity[test_db.User](func(entity *test_db.User) {
					entity.UpdatedAt = lo.Must1(time.Parse(time.DateOnly, "2001-02-05"))
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
				entity.UpdatedAt = time.Time{}
			})
			require.NoError(t, err)

			expectedUsers, err := test.setup()
			require.NoError(t, err)
			expectedIDs := lo.Map(expectedUsers, func(item test_db.User, _ int) uuid.UUID { return item.ID })

			totalIDs, err := test_db.InsertEntitiesWithId[test_db.User, uuid.UUID](c, test_db.TableUser, append(expectedUsers, notExpectedUsers...))
			require.NoError(t, err)
			defer func() {
				require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, totalIDs))
			}()

			users, err := repo.Get(ctx, test.in)
			require.NoError(t, err)

			ids := lo.Map(users, func(item domain.User, _ int) uuid.UUID { return item.ID })

			require.Equal(t, expectedIDs, ids)

			totalCount, err := repo.GetTotalCount(ctx, test.in.Filters)
			require.NoError(t, err)
			require.Equal(t, len(expectedIDs), totalCount)
		})
	}
}

func TestRepository_Get_Sorting(t *testing.T) {
	ctx := context.Background()

	c, err := test_db.NewPsql()
	require.NoError(t, err)
	defer func() { require.NoError(t, c.Close()) }()

	repo := New(c.DB())

	tests := []struct {
		name    string
		in      domain.UserListIn
		compare func(a test_db.User, b test_db.User) int
	}{
		{
			name: "FirstNameAsc",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 10,
				Sorting: &domain.UserListSorting{
					Field:     "firstName",
					Direction: domain.SortingDirectionAsc,
				},
			},
			compare: func(a, b test_db.User) int { return strings.Compare(a.FirstName, b.FirstName) },
		},
		{
			name: "CreatedAtDesc",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 10,
				Sorting: &domain.UserListSorting{
					Field:     "createdAt",
					Direction: domain.SortingDirectionDesc,
				},
			},
			compare: func(a, b test_db.User) int { return b.CreatedAt.Compare(a.CreatedAt) },
		},
		{
			name: "InvalidColumn",
			in: domain.UserListIn{
				Page:    1,
				PerPage: 10,
				Sorting: &domain.UserListSorting{},
			},
			compare: func(a, b test_db.User) int { return 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedUsers, err := test_db.GenerateEntities[test_db.User](10)
			require.NoError(t, err)

			totalIDs, err := test_db.InsertEntitiesWithId[test_db.User, uuid.UUID](c, test_db.TableUser, expectedUsers)
			require.NoError(t, err)
			defer func() {
				require.NoError(t, test_db.DeleteEntitiesById(c, test_db.TableUser, totalIDs))
			}()

			slices.SortFunc(expectedUsers, test.compare)
			expectedIDs := lo.Map(expectedUsers, func(item test_db.User, _ int) uuid.UUID { return item.ID })

			users, err := repo.Get(ctx, test.in)
			require.NoError(t, err)

			ids := lo.Map(users, func(item domain.User, _ int) uuid.UUID { return item.ID })
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
