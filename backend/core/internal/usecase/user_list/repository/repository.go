package repository

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, in domain.UserListIn) ([]domain.User, error) {
	op := "user list"
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("*").
		From("_user").
		Limit(uint64(in.PerPage)).
		Offset(uint64((in.Page - 1) * in.PerPage))

	builder = r.addWhereQuery(builder, in.Filters)
	builder = r.addOrderQuery(builder, in.Sorting)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build %q query: %w", op, err)
	}

	var entities []entity

	query = fmt.Sprintf("-- %s\n%s", op, query)
	err = r.db.SelectContext(ctx, &entities, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec %q query: %w", op, err)
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
		res = append(res, r.buildDatetimeQuery("created_at", filters.CreatedAt))
	}

	if filters.UpdatedAt != nil {
		res = append(res, r.buildDatetimeQuery("updated_at", filters.UpdatedAt))
	}

	if len(res) == 0 {
		return builder
	}

	return builder.Where(res)
}

func (r *Repository) buildLikeQuery(field string, value *string) sq.ILike {
	return sq.ILike{field: fmt.Sprintf("%%%s%%", strings.ToLower(*value))}
}

func (r *Repository) buildDatetimeQuery(field string, value *domain.DateTimeFilter) sq.And {
	res := make(sq.And, 0, 2)
	if value.Start != nil {
		res = append(res, sq.GtOrEq{field: *value.Start})
	}
	if value.End != nil {
		res = append(res, sq.LtOrEq{field: *value.End})
	}
	return res
}

func (r *Repository) GetTotalCount(ctx context.Context, filters domain.UserListFilters) (int, error) {
	op := "user get total count"
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("COUNT(*)").
		From("_user")

	builder = r.addWhereQuery(builder, filters)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("build %q query: %w", op, err)
	}

	var totalCount int
	query = fmt.Sprintf("-- %s\n%s", op, query)
	err = r.db.GetContext(ctx, &totalCount, query, args...)
	if err != nil {
		return 0, fmt.Errorf("exec %q query: %w", op, err)
	}

	return totalCount, nil
}
