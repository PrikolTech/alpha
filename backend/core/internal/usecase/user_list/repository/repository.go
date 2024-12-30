package repository

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, in domain.UserListIn) ([]domain.User, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("*").
		From("_user").
		Limit(uint64(in.PerPage)).
		Offset(uint64((in.Page - 1) * in.PerPage))

	builder = r.addWhereQuery(builder, in.Filters)
	builder = r.addOrderQuery(builder, in.Sorting)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var entities []entity

	err = r.db.SelectContext(ctx, &entities, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return lo.Map(entities, func(item entity, _ int) domain.User { return item.toDomain() }), nil
}

var fieldToColumn = map[domain.SortingField]string{
	domain.SortingFieldEmail:      "email",
	domain.SortingFieldFirstName:  "first_name",
	domain.SortingFieldMiddleName: "middle_name",
	domain.SortingFieldLastName:   "last_name",
	domain.SortingFieldCreatedAt:  "created_at",
	domain.SortingFieldUpdatedAt:  "updated_at",
}

func (r *Repository) addOrderQuery(builder sq.SelectBuilder, sorting *domain.UserListSorting) sq.SelectBuilder {
	if sorting == nil {
		return builder
	}

	column, found := fieldToColumn[sorting.Field]
	if !found {
		return builder
	}

	return builder.OrderBy(fmt.Sprintf("%s %s", column, sorting.Direction))
}

func (r *Repository) addWhereQuery(builder sq.SelectBuilder, filters domain.UserListFilters) sq.SelectBuilder {
	var res sq.And

	if filters.Email != nil {
		res = append(res, r.buildLikeQuery("email", filters.Email))
	}

	if filters.FirstName != nil {
		res = append(res, r.buildLikeQuery("first_name", filters.FirstName))
	}

	if filters.MiddleName != nil {
		res = append(res, r.buildLikeQuery("middle_name", filters.MiddleName))
	}

	if filters.LastName != nil {
		res = append(res, r.buildLikeQuery("last_name", filters.LastName))
	}

	if filters.CreatedAt != nil {
		if filters.CreatedAt.Start != nil {
			res = append(res, sq.GtOrEq{"created_at": *filters.CreatedAt.Start})
		}
		if filters.CreatedAt.End != nil {
			res = append(res, sq.LtOrEq{"created_at": *filters.CreatedAt.End})
		}
	}

	if len(res) == 0 {
		return builder
	}

	return builder.Where(res)
}

func (r *Repository) buildLikeQuery(field string, value *string) sq.ILike {
	return sq.ILike{field: fmt.Sprintf("%%%s%%", strings.ToLower(*value))}
}

func (r *Repository) GetTotalCount(ctx context.Context) (int, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("COUNT(*)").
		From("_user")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	var totalCount int
	err = r.db.GetContext(ctx, &totalCount, query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to exec query: %w", err)
	}

	return totalCount, nil
}